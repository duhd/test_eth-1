
package utils

import (
  "strings"
  "fmt"
  "time"
  "context"
  "test_eth/contracts"
  "math/big"
    "github.com/go-redis/redis"
    "encoding/json"
      "strconv"
  // "github.com/ethereum/go-ethereum"
  // "github.com/ethereum/go-ethereum/accounts/keystore"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/crypto"
  "sync"
  "errors"
  // "io/ioutil"

)

// var sha hash.Hash
type EthClient struct {
	Client   *ethclient.Client
  Redis    *redis.Client
	mux sync.Mutex
}

type Transaction struct {
        Id                string  `json:"Id"`
        RequestTime       int64   `json:"RequestTime"`
        TxReceiveTime     int64   `json:"TxReceiveTime"`
        TxConfirmedTime    []int64 `json:"TxConfiredTime"`
   }


func NewEthClient(url string) (*EthClient, error) {
    fmt.Println("Connect to host: ",url)
    cl, err  := ethclient.Dial("ws://" + url)
    if err != nil {
       fmt.Println("Unable to connect to network:%v\n", err)
       return nil, err
    }

    rd := redis.NewClient(&redis.Options{
      Addr:     cfg.Redis.Host,
      Password: cfg.Redis.Password, // no password set
      DB:       cfg.Redis.Db,  // use default DB
    })

    return &EthClient{Client: cl, Redis: rd}, nil
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
           c.LogEnd(key)
      }
}

func (c *EthClient) TransferToken1(from string,to string,amount string,append string) (string,error) {


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


      c.mux.Lock()
      defer 	c.mux.Unlock()
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
      c.LogStart(key,requestTime)

      return key, nil
}


func (c *EthClient) TransferToken2(from string,to string,amount string,append string) (string,error) {
    	c.mux.Lock()
      defer 	c.mux.Unlock()

      requestTime := time.Now().UnixNano()

      keyjson, err := c.Redis.Get("account:"+from).Result()
      if err != nil {
          return "", err
      }
      redisTime := time.Now().UnixNano()

      auth, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
      if err != nil {
            fmt.Println("Failed to create authorized transactor: %v", err)
            return "", err
      }
      transactorTime := time.Now().UnixNano()

      to_address := common.HexToAddress(to)
      value_transfer := new(big.Int)
      value_transfer, ok := value_transfer.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return "", errors.New("convert amount error")
      }

      note :=  fmt.Sprintf("Transaction:  %s", append)

      prepareAccountTime := time.Now().UnixNano()


    //  fmt.Println("Add contract: ", cfg.Contract.Address)
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

      prepareContractTime := time.Now().UnixNano()

    	var nonce uint64
    	if opts.Nonce == nil {
    		nonce, err = backend.PendingNonceAt(context.Background(), opts.From)
    		if err != nil {
    			return "", fmt.Errorf("failed to retrieve account nonce: %v", err)
    		}
    	} else {
    		nonce = opts.Nonce.Uint64()
    	}

      //nonce := c.getNonce(from)
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

      nonceTime := time.Now().UnixNano()

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


      signTime := time.Now().UnixNano()

    	if err := backend.SendTransaction(opts.Context, signedTx); err != nil {
    		return "", err
    	}
      tx := signedTx

      if err != nil {
          fmt.Println(" Transaction create error: ", err)
          return "",err
      }
      diff0 := (redisTime - requestTime)/1000000
      diff01 := (transactorTime - redisTime)/1000000
      diff1 := (prepareAccountTime - transactorTime)/1000000
      diff2 := (prepareContractTime - prepareAccountTime)/1000000
      diff3 := (nonceTime - prepareContractTime)/1000000
      diff4 := (signTime - nonceTime)/1000000
      diff5 := (time.Now().UnixNano() - signTime)/1000000
      fmt.Println("Transfer: ", nonce," from ",from," to ",to, " amount: ",amount, " note:",append)
      fmt.Println("redisTime, transactorTime, prepareAccountTime,prepareContractTime, nonceTime,signTime, trasactionTime : ",diff0,diff01, diff1,diff2,diff3,diff4,diff5, " Transaction =",tx.Hash().Hex())

      // seed := rand.Intn(100)
      // sha.Write([]byte(strconv.Itoa(seed)))
      // key := "Transfer:" + base64.URLEncoding.EncodeToString(sha.Sum(nil))
      key := strings.TrimPrefix(tx.Hash().Hex(),"0x")
      c.LogStart(key,requestTime)

      return key, nil
}
func (c *EthClient) TransferToken(from string,to string,amount string,append string) (string,error) {
    	c.mux.Lock()
      defer 	c.mux.Unlock()

      requestTime := time.Now().UnixNano()

      // keyjson, err := c.Redis.Get("account:"+from).Result()
      // if err != nil {
      //     return "", err
      // }

      //
      // accountKey, err := keystore.DecryptKey( []byte(keyjson), cfg.Keys.Password)
    	// if err != nil {
      //     fmt.Println("Cannot decrypt key file: ", err)
    	// 	  return "", err
    	// }
      // privateKey := accountKey.PrivateKey
      privatejson, err := c.Redis.Get("private:"+from).Result()
      if err != nil {
          return "", err
      }
      redisTime := time.Now().UnixNano()

      privateKey := crypto.ToECDSAUnsafe([]byte(privatejson))

      //keyAddr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
      keyAddr := common.HexToAddress(from)

      decryptTime := time.Now().UnixNano()

      to_address := common.HexToAddress(to)
      value_transfer := new(big.Int)
      value_transfer, ok := value_transfer.SetString(amount, 10)
      if !ok {
           fmt.Println("SetString: error")
           return "", errors.New("convert amount error")
      }

      note :=  fmt.Sprintf("Transaction:  %s", append)

      prepareAccountTime := time.Now().UnixNano()

      contract_address := common.HexToAddress(cfg.Contract.Address)
      backend := c.Client

      //Get contract
      parsed, err := abi.JSON(strings.NewReader(contracts.VNDWalletABI))
      if err != nil {
          fmt.Println("Error in parse contract ABI: ", contracts.VNDWalletABI)
          return "", err
      }

      input, err := parsed.Pack("transfer", to_address, value_transfer, []byte(note))
    	if err != nil {
        fmt.Println("Error in pack function in ABI: ", contracts.VNDWalletABI)
    		return "", err
    	}

      // Ensure a valid value field and resolve the account nonce
      value := new(big.Int)

      prepareContractTime := time.Now().UnixNano()

      nonce, err := backend.PendingNonceAt(context.Background(), keyAddr)
      if err != nil {
        return "", fmt.Errorf("failed to retrieve account nonce: %v", err)
      }

      gasPrice := new(big.Int)
      gasPrice, ok = gasPrice.SetString("1000", 10)
      var gasLimit uint64 = 40818

      nonceTime := time.Now().UnixNano()

    	// Create the transaction, sign it and schedule it for execution
    	var rawTx *types.Transaction
      rawTx = types.NewTransaction(nonce, contract_address, value, gasLimit, gasPrice, input)

    	//signedTx, err := auth.Signer(types.HomesteadSigner{}, keyAddr, rawTx)

      signer := types.HomesteadSigner{}

      signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), privateKey)
      if err != nil {
        fmt.Println(" Cannot sign contract: ", err)
        return "", err
      }
      signedTx, err := rawTx.WithSignature(signer, signature)


    	if err != nil {
    		return "", err
    	}


      signTime := time.Now().UnixNano()

    	if err := backend.SendTransaction(context.Background(), signedTx); err != nil {
    		return "", err
    	}
      tx := signedTx

      if err != nil {
          fmt.Println(" Transaction create error: ", err)
          return "",err
      }
      diff0 := (redisTime - requestTime)/1000000
      diff01 := (decryptTime - redisTime)/1000000
      diff1 := (prepareAccountTime - decryptTime)/1000000
      diff2 := (prepareContractTime - prepareAccountTime)/1000000
      diff3 := (nonceTime - prepareContractTime)/1000000
      diff4 := (signTime - nonceTime)/1000000
      diff5 := (time.Now().UnixNano() - signTime)/1000000
      fmt.Println("Transfer: ", nonce," from ",from," to ",to, " amount: ",amount, " note:",append)
      fmt.Println("redisTime, decryptTime, prepareAccountTime,prepareContractTime, nonceTime,signTime, trasactionTime : ",diff0,diff01, diff1,diff2,diff3,diff4,diff5, " Transaction =",tx.Hash().Hex())

      // seed := rand.Intn(100)
      // sha.Write([]byte(strconv.Itoa(seed)))
      // key := "Transfer:" + base64.URLEncoding.EncodeToString(sha.Sum(nil))
      key := strings.TrimPrefix(tx.Hash().Hex(),"0x")
      c.LogStart(key,requestTime)

      return key, nil
}
func (c *EthClient) getNonce(account string) uint64 {
     nonce := c.GetNonce(account)
     if nonce == 0 {
        nonce, _ = c.UpdateNoneFromEth(account)
        c.CommitNonce(account,nonce)
     }
     c.NoneIncr(account)
     return nonce
}
func (c *EthClient) UpdateNoneFromEth(account string) (uint64,error) {
      keyjson, err := Redis_client.Get("account:"+account).Result()
      if err != nil {
          return 0, err
      }

      opts, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
      if err != nil {
            fmt.Println("Failed to create authorized transactor: %v", err)
            return 0, err
      }
      var nonce uint64
      if opts.Nonce == nil {
        nonce, err = c.Client.PendingNonceAt(context.Background(), opts.From)
        if err != nil {
          return 0, fmt.Errorf("failed to retrieve account nonce: %v", err)
        }
      } else {
        nonce = opts.Nonce.Uint64()
      }
      if c.CommitNonce(account,nonce) {
        fmt.Println("Failed to create authorized transactor: %v", err)
      }
      return nonce,nil
}

func  (c *EthClient) LogStart(key string,requesttime int64){
  trans :=  &Transaction{
              Id: key,
              RequestTime: requesttime,
              TxReceiveTime: time.Now().UnixNano()}
  value, err := json.Marshal(trans)
  if err != nil {
      fmt.Println(err)
      return
  }
  err = c.Redis.Set("transaction:" + key,string(value), 0).Err()
	if err != nil {
		panic(err)
	}
}

func  (c *EthClient)  LogEnd(key string){
      val, err2 := c.Redis.Get("transaction:" + key).Result()
      if err2 != nil {
          fmt.Println("Cannot find transaction: ", key)
          return
      }
      data := &Transaction{}
      err := json.Unmarshal([]byte(val), data)
      if err != nil {
          fmt.Println("Cannot parse data ", err)
          return
      }
      data.TxConfirmedTime = append(data.TxConfirmedTime, time.Now().UnixNano())
      value, err := json.Marshal(data)

      err = c.Redis.Set("transaction:" + key,string(value), 0).Err()
    	if err != nil {
    	     fmt.Println("Cannot set data ", err)
    	}
      fmt.Println("Finish write transaction: ", key)
}
func  (c *EthClient)  GetNonce(account string) uint64 {
  val, err := c.Redis.Get("nonce:" + account).Result()
  if err != nil {
      fmt.Println("Cannot find nonce of account: ", account)
      return uint64(0)
  }
  value , err := strconv.ParseUint(val, 10, 64)
  if err != nil {
      fmt.Println("Cannot parce nonce of ", val)
      return uint64(0)
  }
  return value
}
func   (c *EthClient)  CommitNonce(account string, nonce uint64) bool {
  err := c.Redis.Set("nonce:" + account,uint64(nonce), 0).Err()
  if err != nil {
       fmt.Println("Cannot set nonce  ", err)
       return false
  }
  return true
}
func  (c *EthClient)  NoneIncr(account string) bool {
  _, err := c.Redis.Incr("nonce:" + account).Result()
	if err != nil {
    fmt.Println("Cannot increase nonce  ", err)
    return false
	}
  return true
}
