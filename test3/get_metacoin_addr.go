package main

import (
		 "context"
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
)


func main(){

	start := time.Now()
	fmt.Println("Start transfer cash")
	const key  = `{"address":"eb80964e1567064ba810b45300fd2ce3193d1684","crypto":{"cipher":"aes-128-ctr","ciphertext":"b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff","cipherparams":{"iv":"3a5118ff590d1a1b435389e754c007e6"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"82fe2f19e0715fbdca86cf864ea0261e6593ca34bb9b9f9cc34ac2b1f5f056ec"},"mac":"2e5746c10e257ef3c9c434333cf40f78c73587fc3832bbdd8ce08808309c3865"},"id":"71fee6ee-4c06-4eee-8739-2c00701e0726","version":3}`
	coinSpawnContractAddr := "0xd6a066a476f73c239e532362e99728c5df35ed59"

	getAddress(key,coinSpawnContractAddr)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
func getAddress(key string, contractAddr string){
				fmt.Println("Start get token address ")

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

				instance, err := metacoin.NewCoinSpawn(contractAddress, client)

				if err != nil {
					log.Fatalf("Unable to connect to network:%v\n", err)
				}
				evtIterator,err := instance.FilterCreateCoinEvt(&bind.FilterOpts{Start: 0, End: nil,  Context: context.Background()})
				if err != nil {
						fmt.Println("Failed to execute a filter query command", "err", err)
						return
				}
				for  evtIterator.Next() {
					 fmt.Println("Addr: ", evtIterator.Event.Addr.Hex())
				}
				fmt.Println("Finished")
	}
