package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"eth/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)


func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("http://localhost:8502")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Create a new instance of the Inbox contract bound to a specific deployed contract
  contract, err :=contracts.NewInbox(common.HexToAddress("0xcc94f30BefDa41ae42F883e1Ae7f6291F1F3698F"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}
	fmt.Println(contract.Message(nil))
}
