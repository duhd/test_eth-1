package utils

import (
  // "time"
  // "context"
  "test_eth/contracts"
  "math/big"
  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "errors"
  "strings"
  "fmt"
  // "encoding/json"
  "crypto/ecdsa"
  "encoding/hex"
  "time"
  "math"
  "github.com/jinzhu/gorm"
  "sync"
)

type TokenAccount struct {
    gorm.Model
    Address string
    WalletId string
    PrivateKey string `sql:"type:text"`
    Active bool `gorm:"default:false"`
}

type WalletAccount struct {
    Routing *RpcRouting
    Address string
    Nonce uint64
    Active bool
    PrivateKey *ecdsa.PrivateKey
    Account *TokenAccount
    Mutex sync.Mutex
}

func  (w *WalletAccount) GetPrivateKey() string {
     return hex.EncodeToString(crypto.FromECDSA(w.PrivateKey))
}

func  (w *WalletAccount) NewTransactor() *bind.TransactOpts {
      key := w.PrivateKey
    	keyAddr := crypto.PubkeyToAddress(key.PublicKey)
    	auth := &bind.TransactOpts{
    		From: keyAddr,
    		Signer: func(signer types.Signer, address common.Address, tx *types.Transaction) (*types.Transaction, error) {
    			if address != keyAddr {
    				return nil, errors.New("not authorized to sign this account")
    			}
    			signature, err := crypto.Sign(signer.Hash(tx).Bytes(), key)
    			if err != nil {
    				return nil, err
    			}
    			return tx.WithSignature(signer, signature)
    		},
    	}
     gasPrice := new(big.Int)
     gasPrice.SetString(cfg.F5Contract.GasPrice,10)
     auth.GasPrice = gasPrice
     auth.GasLimit = cfg.F5Contract.GasLimit
     return auth
}
func (w *WalletAccount) PrepareTransferTransaction(nonce uint64, gLimit uint64, gPrice string, cAddress string,  from string,to string,amount string,append string)  (*types.Transaction, error)  {
      //1. Get nonce and privateKey
      privateKey := w.PrivateKey

      //2. Prepare transaction
      to_address := common.HexToAddress(to)
      value_transfer := new(big.Int)
      value_transfer, ok := value_transfer.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return nil, errors.New("convert amount error")
      }
      note :=  fmt.Sprintf("Transaction:  %s", append)

      contract_address := common.HexToAddress(cAddress)

      //Get contract
      parsed, err := abi.JSON(strings.NewReader(contracts.VNDWalletABI))
      if err != nil {
          fmt.Println("Error in parse contract ABI: ", contracts.VNDWalletABI)
          return nil, err
      }

      input, err := parsed.Pack("transfer", to_address, value_transfer, []byte(note))
      if err != nil {
        fmt.Println("Error in pack function in ABI: ", contracts.VNDWalletABI)
        return nil, err
      }

      // Ensure a valid value field and resolve the account nonce
      value := new(big.Int)

      gasPrice := new(big.Int)
      gasPrice, ok = gasPrice.SetString(gPrice, 10)
      var gasLimit uint64 = gLimit


      // Create the transaction, sign it and schedule it for execution
      var rawTx *types.Transaction
      rawTx = types.NewTransaction(nonce, contract_address, value, gasLimit, gasPrice, input)

      //signedTx, err := auth.Signer(types.HomesteadSigner{}, keyAddr, rawTx)

      signer := types.HomesteadSigner{}

      signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), privateKey)
      if err != nil {
        fmt.Println(" Cannot sign contract: ", err)
        return nil,err
      }

      signedTx, err := rawTx.WithSignature(signer, signature)

      return  signedTx, err
}


func  (w *WalletAccount) BalaneOf() (*big.Float,error) {

      //1. Get wallet address
      address := common.HexToAddress("0x" + w.Address)
      contract_address := common.HexToAddress(cfg.Contract.Address)

      //3. Prepare message to send to contract address
      abi, err := abi.JSON(strings.NewReader(contracts.VNDWalletABI))
    	if err != nil {
    		return nil, err
    	}


      opts := &bind.CallOpts{}
      opts.From = address

    	// Pack the input, call and unpack the results
      method := "balanceOf"
      params := []interface{}{address,}
      input, err := abi.Pack(method, params...)
    	if err != nil {
        fmt.Println("AbI Pack error: ", err)
    		return nil, err
    	}

    	var (
    		msg    = ethereum.CallMsg{From: opts.From, To: &contract_address, Data: input}
    		code   []byte
    		output []byte
    	)

      // 4. Start send message to contract address
    	output, err = w.Routing.CallContract(msg, opts.BlockNumber)
  		if err == nil && len(output) == 0 {
  			// Make sure we have a contract to operate on, and bail out otherwise.
  			if code, err = w.Routing.CodeAt(contract_address, opts.BlockNumber); err != nil {
  				return nil,  err
  			} else if len(code) == 0 {
  				return nil,  errors.New("No code. Please deploy contract")
  			}
  		}
    	if err != nil {
        fmt.Println("S Pack error: ", err)
    		return nil,  err
    	}
      // 5. Unpacke data
      var result = new(big.Int)
    	err = abi.Unpack(&result, method, output)

      fbal := new(big.Float)
      s := result.String()
      fbal.SetString(s)
      fmt.Printf("balance: %f", fbal) // "balance: 74605500.647409"
      return fbal, nil
}

func (w *WalletAccount) TransferToken(from string,to string,amount string,append string) (string,uint64,error) {
      requestTime := time.Now().UnixNano()

      //1. Get nonce
      nonce := w.GetNonce()

      nonceTime := time.Now().UnixNano()

      //2. Prepare transaction
      cAddress := cfg.Contract.Address
      gLimit := cfg.Contract.GasLimit
      gPrice := cfg.Contract.GasPrice

      signedTx, err := w.PrepareTransferTransaction(nonce, gLimit, gPrice, cAddress,  from ,to ,amount ,append )
      if err != nil {
        fmt.Println("Create Transaction error: ", err)
        return "",nonce,err
      }
      txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")

      prepareTime := time.Now().UnixNano()
      //3, Submit transaciton
      err = w.Routing.SubmitTransaction(signedTx,nonce)

      submitTime := time.Now().UnixNano()

      diff0 := (nonceTime - requestTime)/1000
      diff1 := (prepareTime - nonceTime)/1000
      diff2 := (submitTime - prepareTime)/1000

      fmt.Println("NonceTime, PrepareTime,SubmitTime : ",diff0,diff1,diff2, " Transaction =",txhash)
      return txhash, nonce, err
}

func (w *WalletAccount) EthTransfer(to string, amount string) (string,uint64,error)  {
      //1. Get nonce
      nonce := w.GetNonce()

      //2. Prepare transaction
      gLimit := cfg.Contract.GasLimit
      gPrice := cfg.Contract.GasPrice

      fromAddress := common.HexToAddress("0x" + w.Address)
      fmt.Println("Address: ", fromAddress.Hex(), ",Hash: ",fromAddress.Hash())

      gasLimit := uint64(gLimit)
      gasPrice := new(big.Int)
      gasPrice, _ = gasPrice.SetString(gPrice, 10)

      toAddress := common.HexToAddress("0x" + to)

      eth_unit := big.NewFloat(math.Pow10(18))
      amount_value := new(big.Float)
      value, ok := amount_value.SetString(amount)

      if !ok {
           fmt.Println("SetString: error")
           return "", 0, errors.New("convert amount error")
      }
      value = value.Mul(value,eth_unit)

      value_transfer := new(big.Int)
      value.Int(value_transfer)

      var data []byte
      rawTx := types.NewTransaction(nonce, toAddress, value_transfer, gasLimit, gasPrice, data)

      //signer := types.NewEIP155Signer(chainID)
      //signer := types.HomesteadSigner{}
      signer := types.FrontierSigner{}
      signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), w.PrivateKey)
      if err != nil {
        fmt.Println(" Cannot sign contract: ", err)
        return "",0,err
      }

      signedTx, err := rawTx.WithSignature(signer, signature)

      txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")

      err = w.Routing.SubmitTransaction(signedTx,nonce)

      return txhash, nonce, err
}


func (w *WalletAccount) EthBalaneOf() (*big.Float, error) {
    w.Mutex.Lock()
    defer w.Mutex.Unlock()
    account := common.HexToAddress("0x" + w.Address)
    return w.Routing.BalanceAt(account)
}
