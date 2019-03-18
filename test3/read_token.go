package main

import (
		"test_eth/contracts/metacoin"
		"github.com/ethereum/go-ethereum/ethclient"
		"fmt"
		"github.com/ethereum/go-ethereum/common"
		util "test_eth/utils"
)

func main(){
		if len(os.Args) <1 {
			 fmt.Println("Please use syntax: go run deploy_metacoin.go  metacoin_address")
			 return
		}
		tokenContractAddr := os.Args[1]

		var c util.Config
		c.GetConf()

		client, err  := ethclient.Dial(c.Webservice)
		if err != nil {
			fmt.Println("Unable to connect to network:%v\n", err)
		}

		contractAddress := common.HexToAddress(tokenContractAddr)
		instance, err := metacoin.NewMetaCoin(contractAddress, client)
		if err != nil {
			  fmt.Println("Unable to bind to deployed instance of contract")
		}
		fmt.Println("Start reading: ", metacoin.balances)
}
