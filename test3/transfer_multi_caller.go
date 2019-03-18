package main

import (
			"log"
			"test_eth/contracts/metacoin"
			"github.com/ethereum/go-ethereum/ethclient"
			"github.com/ethereum/go-ethereum/accounts/abi/bind"
			"github.com/ethereum/go-ethereum/common"
			// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
		  // "github.com/ethereum/go-ethereum/rpc"
			"strings"
			"fmt"
			"math/big"
			"time"
			"sync"
)

//const key  = `paste the contents of your JSON key file here`
// const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`



func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	tokenContractAddr := "0xAE6313a252d905cdc0d8e9116fE31696CC832145"
	recvAddr := "0x7303040c37df72cc4410511d6ccb51e7ab7a42d5"

	go func(){
	  fmt.Println("Start thread 1")
		const key  = `{"address":"eb80964e1567064ba810b45300fd2ce3193d1684","crypto":{"cipher":"aes-128-ctr","ciphertext":"b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff","cipherparams":{"iv":"3a5118ff590d1a1b435389e754c007e6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"82fe2f19e0715fbdca86cf864ea0261e6593ca34bb9b9f9cc34ac2b1f5f056ec"},"mac":"2e5746c10e257ef3c9c434333cf40f78c73587fc3832bbdd8ce08808309c3865"},"id":"71fee6ee-4c06-4eee-8739-2c00701e0726","version":3}`

		callerContractAddr := "0xa341d306c6c90d19f7be11d4d347bd4206854b5c"

		transferToken("Thread 1", key, callerContractAddr, tokenContractAddr, recvAddr)

		defer wg.Done()
	}()

	go func(){
	  fmt.Println("Start thread 2")
		const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
		callerContractAddr := "0x27af2a27850e185ef3ccab0c62fdea25ace9a8f9"
		transferToken("Thread 2", key, callerContractAddr, tokenContractAddr, recvAddr)

		defer wg.Done()
		}()


	wg.Wait()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)
}
func transferToken(threadName string,key string, callerContractAddr string, tokenContractAddr string, recvAddr string ){
				fmt.Println("Start thread: ",threadName)

				// Get credentials for the account to charge for contract deployments
				auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
				if err != nil {
					log.Fatalf("Failed to create authorized transactor: %v", err)
				}

				//Create transation
				auth.GasPrice = big.NewInt(1)
				auth.GasLimit = 100000000000000

				// connect to an ethereum node  hosted by infura
				// client1, err  := ethclient.Dial("http://localhost:8502")
				client, err  := ethclient.Dial("ws://localhost:8546")
				if err != nil {
					log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
				}

				contractAddress := common.HexToAddress(callerContractAddr)
				instance, err := metacoin.NewCoinCaller(contractAddress, client)

				if err != nil {
					log.Fatalf("Unable to connect to network:%v\n", err)
				}
				coinContractAddress := common.HexToAddress(tokenContractAddr)

				receiver := common.HexToAddress(recvAddr)
				amount := big.NewInt(1)

				for i:=0 ; i< 10000; i++ {
					tx, err := instance.SendCoin(auth,coinContractAddress,receiver,amount)
					if err != nil {
							log.Fatalf(threadName," Transaction ",i," create error: ", err)
					}
					fmt.Println(threadName," Transaction ",i," , tx =",tx.Hash().Hex())
				}

	}
