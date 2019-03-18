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

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	start := time.Now()

	recvAddr := "0x7303040c37df72cc4410511d6ccb51e7ab7a42d5"

	go funct(){
	  fmt.Println("Start thread 1")
		const key  = `{"address":"eb80964e1567064ba810b45300fd2ce3193d1684","crypto":{"cipher":"aes-128-ctr","ciphertext":"b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff","cipherparams":{"iv":"3a5118ff590d1a1b435389e754c007e6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"82fe2f19e0715fbdca86cf864ea0261e6593ca34bb9b9f9cc34ac2b1f5f056ec"},"mac":"2e5746c10e257ef3c9c434333cf40f78c73587fc3832bbdd8ce08808309c3865"},"id":"71fee6ee-4c06-4eee-8739-2c00701e0726","version":3}`
		contractAddr := "0x382d559d774299a8e2bf48d54a41e54b7a3991b4"
		transferToken("Thread 1",key,contractAddr,recvAddr)
		defer wg.Done()
	}()
	go func(){
	  fmt.Println("Start thread 2")
		const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
		contractAddr := "0x0c5375154e912135356c8320b26337a044c10852"
		transferToken("Thread 2",key,contractAddr,recvAddr)
		defer wg.Done()
		}()
	wg.Wait()
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)
}
func transferToken(threadName string,key string, contractAddr string, recvAddr string ){
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

				contractAddress := common.HexToAddress(contractAddr)
				instance, err := contracts.NewVNDWallet(contractAddress, client)

				if err != nil {
					log.Fatalf("Unable to connect to network:%v\n", err)
				}

				address := common.HexToAddress(recvAddr)
				value := big.NewInt(1)

				for i:=0 ; i< 100000; i++ {
					note :=  fmt.Sprintf(",%s, Transaction: %d",threadName,i)
					tx, err := instance.Transfer(auth, address, value, []byte(note))
					if err != nil {
							log.Fatalf(threadName," Transaction ",i," create error: ", err)
					}
					fmt.Println(threadName," Transaction ",i," , tx =",tx.Hash().Hex())
				}

	}
