package main

import (
	  "context"
		"test_eth/contracts/metacoin"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"fmt"
		"time"
		"github.com/ethereum/go-ethereum/common"
		util "test_eth/utils"
)

func main(){
		if len(os.Args) <1 {
			 fmt.Println("Please use syntax: go run deploy_metacoin.go  coinspawn_address")
			 return
		}
	  spawnContractAddr := os.Args[1]

		var c util.Config
		c.GetConf()
  	client, err  := ethclient.Dial(c.Webservice)

  	if err != nil {
  		fmt.Println("Unable to connect to network:%v\n", err)
  	}

		contractAddress := common.HexToAddress(spawnContractAddr)
		instance, err := metacoin.NewCoinSpawn(contractAddress, client)
		if err != nil {
			  fmt.Println("Unable to bind to deployed instance of contract")
		}
		fmt.Println("Start listening")

		eventCh := make(chan *metacoin.CoinSpawnCreateCoinEvt,10)

		sub,err := instance.WatchCreateCoinEvt(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh )
		defer sub.Unsubscribe()
     for {
         select {
							 case event := <-eventCh:
								  fmt.Println("time:",time.Now(),", Addr: ", event.Addr.Hex()," Initial Value: ", event.InitialBalance )
		        }
     }
}
