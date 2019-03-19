package main

import (
	"os"
  "context"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var logger = log.New()


func main(){
  	// connect to an ethereum node  hosted by infura
		if len(os.Args) <5 {
			 fmt.Println("Please use syntax: go run deploy_metacoin.go  webserver  password contract_addr from to")
			 return
		}
		webserver := os.Args[1]
		password := os.Args[2]

		contractAddr := os.Args[3]
		fromAddr := os.Args[4]
		toAddr := os.Args[5]

  	client, err  := ethclient.Dial(webserver)

  	if err != nil {
  		logger.Crit("Unable to connect to network:%v\n", err)
  	}

		fmt.Println("Start read event")
		contractAddress := common.HexToAddress(contractAddr)
		instance, err := contracts.NewVNDWallet(contractAddress, client)
		if err != nil {
			  fmt.Println("Unable to bind to deployed instance of contract")
		}
		from :=  []common.Address{common.HexToAddress(fromAddr),}
		to :=   []common.Address{common.HexToAddress(toAddr),}

		evtIterator,err := instance.FilterTransfer(&bind.FilterOpts{Start: 0, End: nil,  Context: context.Background()},from,to)
		if err != nil {
				fmt.Println("Failed to execute a filter query command", "err", err)
				return
		}
		for  evtIterator.Next() {
			 fmt.Println("From: ", evtIterator.Event.From.Hex(), ", To: ",  evtIterator.Event.To.Hex(), ", Value: ",evtIterator.Event.Value, " Data: ",string(evtIterator.Event.Data))
		}
		fmt.Println("Finished")
}
