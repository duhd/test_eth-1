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
    "encoding/json"
)

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

  err = clientPool.TransferToken(signedTx, nonce)

  // fmt.Println("Send message to log server (redis pool)")
  Rclients.LogStart(txhash, nonce, requestTime)

  submitTime := time.Now().UnixNano()
  diff0 := (prepareTime - requestTime)/1000
  diff1 := (submitTime - prepareTime)/1000
  fmt.Println("Transfer: ", nonce," from ",from," to ",to, " amount: ",amount, " note:",append)
  fmt.Println("prepareTime, submitTime : ",diff0,diff1, " Transaction =",txhash)

  return txhash, err
}


func Report() string {
      client := Rclients.getClient()
      keys, err  := client.Keys("transaction:*").Result()
      if err != nil {
        // handle error
        fmt.Println(time.Now()," Cannot get keys ")
      }
      vals, err1 := client.MGet(keys...).Result()
      if err1 != nil {
        // handle error
        fmt.Println(time.Now()," Cannot get values of  keys: ", keys)
      }

      fmt.Println("Elements: ", len(keys))
      diff_arr1 := []int64{}
      diff_arr := []int64{}

      for _, element := range vals {
          data := &Transaction{}
          err2 := json.Unmarshal([]byte(element.(string)), data)
          if err2 != nil {
              fmt.Println(time.Now()," Element:", element, ", Error:", err2)
              continue
          }
          fmt.Println("ID:",data.Id,"RequestTime:",data.RequestTime,
            "TxReceiveTime:",data.TxReceiveTime,"TxConfirmedTime:",data.TxConfirmedTime)

          var max int64 = 0
          if data.TxConfirmedTime != nil {
              for _,value := range data.TxConfirmedTime {
                  if value > max {
                     max = value
                  }
              }
              diff1 := data.TxReceiveTime - data.RequestTime
              diff_arr1 = append(diff_arr1,diff1)
          }
          // else {
          //     max = time.Now().UnixNano()
          // }
          if max >0 {
              diff := max  - data.TxReceiveTime
              diff_arr = append(diff_arr,diff)
          }
      }
      var total1 int64 = 0
    	for _, value1:= range diff_arr1 {
    		total1 += value1
    	}
      len1 := int64(len(diff_arr1))
      var avg1 int64 = 0
      if len1 >0 {
        	avg1 = total1/(len1 *1000)
      }

      var total int64 = 0
    	for _, value:= range diff_arr {
    		total += value
    	}
      len2 := int64(len(keys))
      len := int64(len(diff_arr))
      var avg int64 = 0
      if len >0 {
        	avg = total/(len *1000)
      }
      return fmt.Sprintf("Total Tx: %v , Total Complete TX: %v ,Avg RequestTime: %v , Avg Onchain: %v ", len2, len,avg1, avg)
}
