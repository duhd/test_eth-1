package main

import (
			"os"
			"log"
			"test_eth/contracts/metacoin"
			util "test_eth/utils"
			"github.com/ethereum/go-ethereum/ethclient"
			"github.com/ethereum/go-ethereum/accounts/abi/bind"
			"github.com/ethereum/go-ethereum/common"
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
  spawnContractAddr := os.Args[1]

	var c util.Config
	c.GetConf()
	fmt.Println(c)

	start := time.Now()
	fmt.Println("Start transfer cash")


	fmt.Println("Start init token ")

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(c.Key),c.Password)
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

	contractAddress := common.HexToAddress(spawnContractAddr)

	instance, err := metacoin.NewCoinSpawn(contractAddress, client)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	initialBalance := big.NewInt(1000000000)

	tx,err := instance.CreateCoin(auth,initialBalance)
	if err != nil {
			log.Fatalf("Transaction create error: ", err)
	}
	fmt.Println("Transaction hash:",tx.Hash().Hex()) // 1
	fmt.Println("Finished")

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
