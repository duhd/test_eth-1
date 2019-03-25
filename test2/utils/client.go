
package utils

import (
  "strings"
  "fmt"
  "time"
  "context"
  "test_eth/contracts"
  "math/big"
  // "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "sync"
  "errors"
)

// var sha hash.Hash
type EthClient struct {
	Client   *ethclient.Client
	mux sync.Mutex
}

func NewEthClient(url string) (*EthClient, error) {
    fmt.Println("Connect to host: ",url)
    cl, err  := ethclient.Dial("http://" + url)
    if err != nil {
       fmt.Println("Unable to connect to network:%v\n", err)
       return nil, err
    }
    return &EthClient{Client: cl}, nil
}
func (c *EthClient) BalaneOf(account string) (*big.Float,error) {
    	c.mux.Lock()
      defer   c.mux.Unlock()

      address := common.HexToAddress("0x" + account)
      fmt.Println("Add contract: ", cfg.Contract.Address)
      wallet, err1 := contracts.NewVNDWallet(common.HexToAddress(cfg.Contract.Address), c.Client)
      if err1 != nil {
         fmt.Println("Unable to bind to deployed instance of contract:%v\n")
         return nil,err1
     }

      bal, err := wallet.BalanceOf(&bind.CallOpts{}, address)

      if err != nil {
        fmt.Println("Get balanceof: ", err)
        return nil,err
      }

      fbal := new(big.Float)

      fbal.SetString(bal.String())
      fmt.Printf("balance: %f", bal) // "balance: 74605500.647409"


      return fbal, nil
}

func (c *EthClient) UpdateReceipt(header *types.Header ){
      c.mux.Lock()
      defer 	c.mux.Unlock()

      block, err := c.Client.BlockByHash(context.Background(), header.Hash())
      if err != nil {
        fmt.Println("Errror blockbyhash: ",err)
        return
        //log.Fatal(err)
      }
      for _, transaction := range block.Transactions(){
           fmt.Println("Transaction: ",transaction.Hash().Hex())
           key := strings.TrimPrefix(transaction.Hash().Hex(),"0x")
           LogEnd(key)
      }
}
func (c *EthClient) TransferTokenRaw(from string,to string,amount string,append string) (string,error) {
    	c.mux.Lock()
      defer 	c.mux.Unlock()

      requestTime := time.Now().UnixNano()

      keyjson, err := Redis_client.Get("account:"+from).Result()
      if err != nil {
          return "", err
      }

      auth, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
      if err != nil {
            fmt.Println("Failed to create authorized transactor: %v", err)
            return "", err
      }

      address := common.HexToAddress(to)
      value := new(big.Int)
      value, ok := value.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return "", errors.New("convert amount error")
      }

      note :=  fmt.Sprintf("Transaction:  %s", append)

      fmt.Println("Add contract: ", cfg.Contract.Address)
      wallet, err1 := contracts.NewVNDWallet(common.HexToAddress(cfg.Contract.Address), c.Client)
      if err1 != nil {
         fmt.Println("Unable to bind to deployed instance of contract:%v\n")
         return "",err1
     }

      tx, err := wallet.Transfer(auth, address, value, []byte(note))
      if err != nil {
          fmt.Println(" Transaction create error: ", err)
          return "",err
      }
      fmt.Println(" Transaction =",tx.Hash().Hex())
      // seed := rand.Intn(100)
      // sha.Write([]byte(strconv.Itoa(seed)))
      // key := "Transfer:" + base64.URLEncoding.EncodeToString(sha.Sum(nil))
      key := strings.TrimPrefix(tx.Hash().Hex(),"0x")
      LogStart(key,requestTime)

      return key, nil
}



func (c *EthClient) TransferToken(from string,to string,amount string,append string) (string,error) {
    	c.mux.Lock()
      defer 	c.mux.Unlock()

      requestTime := time.Now().UnixNano()

      keyjson, err := Redis_client.Get("account:"+from).Result()
      if err != nil {
          return "", err
      }

      auth, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
      if err != nil {
            fmt.Println("Failed to create authorized transactor: %v", err)
            return "", err
      }

      to_address := common.HexToAddress(to)
      value_transfer := new(big.Int)
      value_transfer, ok := value_transfer.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return "", errors.New("convert amount error")
      }

      note :=  fmt.Sprintf("Transaction:  %s", append)

      fmt.Println("Add contract: ", cfg.Contract.Address)
      contract_address := common.HexToAddress(cfg.Contract.Address)
      backend := c.Client

      //Get contract
      parsed, err := abi.JSON(strings.NewReader(contracts.VNDWalletABI))
      if err != nil {
          fmt.Println("Error in parse contract ABI: ", contracts.VNDWalletABI)
          return "", err
      }

      //contract := bind.NewBoundContract(contract_address, parsed, backend, backend, backend)
      //&VNDWallet{VNDWalletCaller: VNDWalletCaller{contract: contract}, VNDWalletTransactor: VNDWalletTransactor{contract: contract},
      //tx, err := contract.Transact(auth, "transfer", to_address, value, []byte(note))
      input, err := parsed.Pack("transfer", to_address, value_transfer, []byte(note))
    	if err != nil {
        fmt.Println("Error in pack function in ABI: ", contracts.VNDWalletABI)
    		return "", err
    	}
      //tx, err := contract.transact(opts, &contract_address, input)

      opts := auth
      // Ensure a valid value field and resolve the account nonce
    	value := opts.Value
      opts.Context = context.Background()
    	if value == nil {
    		value = new(big.Int)
    	}
    	var nonce uint64
    	if opts.Nonce == nil {
    		nonce, err = backend.PendingNonceAt(context.Background(), opts.From)
    		if err != nil {
    			return "", fmt.Errorf("failed to retrieve account nonce: %v", err)
    		}
    	} else {
    		nonce = opts.Nonce.Uint64()
    	}
    	// Figure out the gas allowance and gas price values
    	// gasPrice := opts.GasPrice
    	// if gasPrice == nil {
    	// 	gasPrice, err = backend.SuggestGasPrice(context.Background())
    	// 	if err != nil {
    	// 		return "", fmt.Errorf("failed to suggest gas price: %v", err)
    	// 	}
    	// }
      // fmt.Println("gasPrice:= ",gasPrice)

      gasPrice := new(big.Int)
      gasPrice, ok = gasPrice.SetString("1000", 10)

    	// gasLimit := opts.GasLimit
    	// if gasLimit == 0 {
    	// 	// Gas estimation cannot succeed without code for method invocations
      //   if code, err := backend.PendingCodeAt(context.Background(), contract_address); err != nil {
      //     return "", err
      //   } else if len(code) == 0 {
      //     return "",  errors.New("code = 0")
      //   }
    	// 	// If the contract surely has code (or code is not needed), estimate the transaction
    	// 	msg := ethereum.CallMsg{From: opts.From, To: &contract_address, Value: value, Data: input}
    	// 	gasLimit, err = backend.EstimateGas(context.Background(), msg)
    	// 	if err != nil {
    	// 		return "", fmt.Errorf("failed to estimate gas needed: %v", err)
    	// 	}
    	// }
      //
      // fmt.Println("gasLimit:= ",gasLimit)

      var gasLimit uint64 = 40818

    	// Create the transaction, sign it and schedule it for execution
    	var rawTx *types.Transaction
      rawTx = types.NewTransaction(nonce, contract_address, value, gasLimit, gasPrice, input)

      if opts.Signer == nil {
    		return "", errors.New("no signer to authorize the transaction with")
    	}

    	signedTx, err := opts.Signer(types.HomesteadSigner{}, opts.From, rawTx)
    	if err != nil {
    		return "", err
    	}
    	if err := backend.SendTransaction(opts.Context, signedTx); err != nil {
    		return "", err
    	}
      tx := signedTx


      if err != nil {
          fmt.Println(" Transaction create error: ", err)
          return "",err
      }
      fmt.Println(" Transaction =",tx.Hash().Hex())
      // seed := rand.Intn(100)
      // sha.Write([]byte(strconv.Itoa(seed)))
      // key := "Transfer:" + base64.URLEncoding.EncodeToString(sha.Sum(nil))
      key := strings.TrimPrefix(tx.Hash().Hex(),"0x")
      LogStart(key,requestTime)

      return key, nil
}
