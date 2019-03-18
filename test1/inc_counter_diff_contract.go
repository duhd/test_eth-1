package main

import (
			"log"
			"test_eth/contracts"
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
	go func(){
					fmt.Println("Start thread 1")
					const key  = `{"address":"eb80964e1567064ba810b45300fd2ce3193d1684","crypto":{"cipher":"aes-128-ctr","ciphertext":"b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff","cipherparams":{"iv":"3a5118ff590d1a1b435389e754c007e6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"82fe2f19e0715fbdca86cf864ea0261e6593ca34bb9b9f9cc34ac2b1f5f056ec"},"mac":"2e5746c10e257ef3c9c434333cf40f78c73587fc3832bbdd8ce08808309c3865"},"id":"71fee6ee-4c06-4eee-8739-2c00701e0726","version":3}`
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
					client1, err  := ethclient.Dial("ws://localhost:8546")
					if err != nil {
						log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
					}

					contractAddress := common.HexToAddress("0x7b3d4b04434315a9b54249b6175a8feaa1d45994")
					instance1, err := contracts.NewCounter(contractAddress, client1)

					if err != nil {
						log.Fatalf("Unable to connect to network:%v\n", err)
					}
					for i:=0 ; i< 100000; i++ {
									tx, err := instance1.Increase(auth)
									if err != nil {
										log.Fatal(err)
									}
									fmt.Println("Thread 1: tx: ",i," , sent:", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
					}
			  	defer wg.Done()
		}()
	go func(){
				  fmt.Println("Start thread 2")
					const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
					// Get credentials for the account to charge for contract deployments
					auth, err := bind.NewTransactor(strings.NewReader(key), "123456")
					if err != nil {
						log.Fatalf("Failed to create authorized transactor: %v", err)
					}

					//Create transation
					auth.GasPrice = big.NewInt(1)
					auth.GasLimit = 100000000000000

					// connect to an ethereum node  hosted by infura
					// client2, err  := ethclient.Dial("http://localhost:8502")
					client2, err  := ethclient.Dial("ws://localhost:8546")

					if err != nil {
						log.Fatalf("Unable to connect to network:%v\n", err)
					}

					contractAddress := common.HexToAddress("0xbb7dcbfb547171d1d619ba4f83fde03da3cbc5bc")
					instance2, err := contracts.NewCounter(contractAddress, client2)

					if err != nil {
						log.Fatalf("Unable to connect to network:%v\n", err)
					}

					for i:=0 ; i< 100000; i++ {
								tx, err := instance2.Increase(auth)
								if err != nil {
									log.Fatal(err)
								}
								fmt.Println("Thread 2: tx: ",i," , sent:", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
					}
			  	defer wg.Done()
		}()
	wg.Wait()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
