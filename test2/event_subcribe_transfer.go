package main

import (
		"os"
		"strings"
	  "context"
		"test_eth/contracts"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"fmt"
		"time"
		"github.com/ethereum/go-ethereum/common"
)
func str2addr(str string) []common.Address {
	 list := strings.Split(str,",")
	 ret :=  make([]common.Address,0)
	 for _, element := range list {
		 fmt.Println("Addr: ",element)
		  addr := common.HexToAddress(element)
      ret = append(ret,addr)
	 }
	 return ret
}
func main(){
	if len(os.Args) <5 {
		 fmt.Println("Please use syntax: go run event_subcribe_transfer.go  webserver contract_addr froms to")
		 return
	}
	webserver := os.Args[1]
	contractAddr := os.Args[2]
	fromAddrs := os.Args[3]
	toAddrs := os.Args[4]

	client, err  := ethclient.Dial(webserver)
	if err != nil {
		fmt.Println("Unable to connect to network:%v\n", err)
	}


	contractAddress := common.HexToAddress(contractAddr)
	instance, err := contracts.NewVNDWallet(contractAddress, client)
	if err != nil {
		  fmt.Println("Unable to bind to deployed instance of contract")
	}
	fmt.Println("Start listening")

	eventCh := make(chan *contracts.VNDWalletTransfer,10)

	from :=  str2addr(fromAddrs)
	to :=   str2addr(toAddrs)

	sub,err := instance.WatchTransfer(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh, from, to )
	defer sub.Unsubscribe()
	//
	// var event *contracts.CounterCounterIncreasedEvt
   for {
       select {
						 case event := <-eventCh:
							  fmt.Println("time:",time.Now(),", From: ", event.From.Hex(),", To: ", event.To.Hex(), ", Value: ", event.Value,",Data: ",string(event.Data) )
	        }
   }

}
