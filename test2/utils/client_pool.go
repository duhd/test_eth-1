package utils

import (
  "time"
  "test_eth/contracts"
  "math/big"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
  "errors"
  "strings"
  "fmt"
)

type ClientPool struct {
  Clients []*EthClient
  Current int
}

var clientPool *ClientPool


func NewClientPool()  *ClientPool{
    var clients []*EthClient
    max_connection := cfg.Webserver.MaxRpcConnection
    for i:=0 ; i<max_connection; i++ {
         for _,host := range cfg.Networks {
              ethclient, err := NewEthClient(host.Http)
              if err != nil {
                continue
              }
              clients = append(clients,ethclient)
            }
     }
     clientPool =  &ClientPool{
        Clients: clients,
        Current:0,
     }
     return clientPool
}

func (cp *ClientPool) GetClient() (*EthClient) {
    client := cp.Clients[cp.Current]
    cp.Current = cp.Current + 1
    cp.Current = cp.Current % len(cp.Clients)
    return client
}


func BalaneOf(account string) (*big.Float,error) {
  client := clientPool.GetClient()
  return client.BalaneOf(account)
}


func PrepareTransferToken(from string,to string,amount string,append string)  (*types.Transaction, error,uint64)  {
      wallet := GetWallet(from)

      if wallet == nil {
        return nil,errors.New("Cannot load wallet"),0
      }
      privateKey := wallet.PrivateKey

      to_address := common.HexToAddress(to)
      value_transfer := new(big.Int)
      value_transfer, ok := value_transfer.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return nil, errors.New("convert amount error"),0
      }
      note :=  fmt.Sprintf("Transaction:  %s", append)

      contract_address := common.HexToAddress(cfg.Contract.Address)

      //Get contract
      parsed, err := abi.JSON(strings.NewReader(contracts.VNDWalletABI))
      if err != nil {
          fmt.Println("Error in parse contract ABI: ", contracts.VNDWalletABI)
          return nil, err,0
      }

      input, err := parsed.Pack("transfer", to_address, value_transfer, []byte(note))
      if err != nil {
        fmt.Println("Error in pack function in ABI: ", contracts.VNDWalletABI)
        return nil, err,0
      }

      // Ensure a valid value field and resolve the account nonce
      value := new(big.Int)

      nonce := wallet.GetNonce()
      gasPrice := new(big.Int)
      gasPrice, ok = gasPrice.SetString(cfg.Contract.GasPrice, 10)
      var gasLimit uint64 = cfg.Contract.GasLimit


      // Create the transaction, sign it and schedule it for execution
      var rawTx *types.Transaction
      rawTx = types.NewTransaction(nonce, contract_address, value, gasLimit, gasPrice, input)

      //signedTx, err := auth.Signer(types.HomesteadSigner{}, keyAddr, rawTx)

      signer := types.HomesteadSigner{}

      signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), privateKey)
      if err != nil {
        fmt.Println(" Cannot sign contract: ", err)
        return nil,err,0
      }

      signedTx, err := rawTx.WithSignature(signer, signature)
      return  signedTx, err , nonce
}
func TransferToken(from string,to string,amount string,append string) (string,error) {
  requestTime := time.Now().UnixNano()
  signedTx, err, nonce := PrepareTransferToken(from,to,amount,append)
  if err != nil {
    fmt.Println("Create Transaction error: ", err)
    return "", err
  }

  txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")
  prepareTime := time.Now().UnixNano()
  stored := LogStart(txhash, nonce, requestTime)

  redisTime := time.Now().UnixNano()
  client := clientPool.GetClient()
  if stored {
      _, err := client.TransferToken(signedTx,nonce)
      if err != nil {
        return txhash, err
      }
  }
  submitTime := time.Now().UnixNano()
  diff0 := (prepareTime - requestTime)/1000
  diff1 := (redisTime - prepareTime)/1000
  diff2 := (submitTime - redisTime)/1000
  fmt.Println("Transfer: ", nonce," from ",from," to ",to, " amount: ",amount, " note:",append)
  fmt.Println("prepareTime, redisTime, submitTime : ",diff0,diff1, diff2, " Transaction =",txhash)

  return txhash, err
}
