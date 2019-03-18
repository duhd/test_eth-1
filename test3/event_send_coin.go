package main

import (
	  "context"
		"test_eth/contracts/metacoin"
		"github.com/ethereum/go-ethereum/ethclient"
		// "github.com/ethereum/go-ethereum/accounts/abi"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	  // "github.com/ethereum/go-ethereum/rpc"
		// "math/big"
		// "strings"
		"fmt"
		"time"
		// "os"
		// "github.com/ethereum/go-ethereum"
		"github.com/ethereum/go-ethereum/common"
		// "github.com/ethereum/go-ethereum/core/types"
		// "github.com/ethereum/go-ethereum/log"
		// "github.com/ethereum/go-ethereum/crypto"

)

//const key  = `paste the contents of your JSON key file here`
// const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
const key  = `{"address":"eb80964e1567064ba810b45300fd2ce3193d1684","crypto":{"cipher":"aes-128-ctr","ciphertext":"b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff","cipherparams":{"iv":"3a5118ff590d1a1b435389e754c007e6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"82fe2f19e0715fbdca86cf864ea0261e6593ca34bb9b9f9cc34ac2b1f5f056ec"},"mac":"2e5746c10e257ef3c9c434333cf40f78c73587fc3832bbdd8ce08808309c3865"},"id":"71fee6ee-4c06-4eee-8739-2c00701e0726","version":3}`
func main(){
    	// connect to an ethereum node  hosted by infura
    	client, err  := ethclient.Dial("ws://localhost:8546")
			// client, err  := ethclient.Dial("http://localhost:8502")

    	if err != nil {
    		fmt.Println("Unable to connect to network:%v\n", err)
    	}

			coinCallerContractAddr := "0xa341d306c6c90d19f7be11d4d347bd4206854b5c"
			contractAddress := common.HexToAddress(coinCallerContractAddr)

			instance, err := metacoin.NewCoinCaller(contractAddress, client)
			if err != nil {
				  fmt.Println("Unable to bind to deployed instance of contract")
			}
			fmt.Println("Start listening")

			eventCh := make(chan *metacoin.CoinCallerSendCoinEvt,10)

			sub,err := instance.WatchSendCoinEvt(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh )
			defer sub.Unsubscribe()
			//
			// var event *contracts.CounterCounterIncreasedEvt

	     for {
	         select {
								 case event := <-eventCh:
									  fmt.Println("time:",time.Now(),", From: ", event.From.Hex(),", To: ", event.To.Hex(), " Amount: ", event.Amount,", Status: ",event.Txstatus,", Balance: ",event.Balance )
			        }
	     }
}
