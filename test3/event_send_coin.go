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
			 fmt.Println("Please use syntax: go run deploy_metacoin.go  coincaller_address")
			 return
		}
		coinCallerContractAddr := os.Args[1]

		var c util.Config
		c.GetConf()
  	// connect to an ethereum node  hosted by infura
  	client, err  := ethclient.Dial(c.Webservice)
		// client, err  := ethclient.Dial("http://localhost:8502")

  	if err != nil {
  		fmt.Println("Unable to connect to network:%v\n", err)
  	}

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
