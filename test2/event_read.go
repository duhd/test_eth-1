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

const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`
const webserver = "http://172.101.0.17:8501"
const password = "123456"

func main(){
  	// connect to an ethereum node  hosted by infura
		if len(os.Args) <3 {
			 fmt.Println("Please use syntax: go run deploy_metacoin.go  contract_addr from to")
			 return
		}
		contractAddr := os.Args[1]
		fromAddr := os.Args[2]
		toAddr := os.Args[3]

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
