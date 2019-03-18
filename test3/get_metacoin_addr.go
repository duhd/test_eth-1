package main

import (
		 "context"
			"log"
			"test_eth/contracts/metacoin"
			"github.com/ethereum/go-ethereum/ethclient"
			"github.com/ethereum/go-ethereum/accounts/abi/bind"
			"github.com/ethereum/go-ethereum/common"
			util "test_eth/utils"
			"strings"
			"fmt"
			"math/big"
			"time"
)


func main(){
	if len(os.Args) <1 {
		 fmt.Println("Please use syntax: go run deploy_metacoin.go  coinspawn_address")
		 return
	}
	coinSpawnContractAddr := os.Args[1]

	start := time.Now()
	fmt.Println("Start transfer cash")

	getAddress(coinSpawnContractAddr)
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
func getAddress(contractAddr string){
			var c util.Config
			c.GetConf()

			fmt.Println("Start get token address ")

			// Get credentials for the account to charge for contract deployments
			auth, err := bind.NewTransactor(strings.NewReader(c.Key), c.Password)
			if err != nil {
				log.Fatalf("Failed to create authorized transactor: %v", err)
			}

			//Create transation
			auth.GasPrice = big.NewInt(1)
			auth.GasLimit = 100000000000000


			client, err  := ethclient.Dial(c.Webservice)
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
