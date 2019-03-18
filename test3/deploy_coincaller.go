package main

import (
	"log"
	"test_eth/contracts/metacoin"
	util "test_eth/utils"
		// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
  // "github.com/ethereum/go-ethereum/rpc"
		// "math/big"
	"strings"
	"fmt"
	// "time"
	// "os"
)

//const key  = `paste the contents of your JSON key file here`
func main(){
	var c util.Config
	c.GetConf()

	fmt.Println(c)

	// connect to an ethereum node  hosted by infura
	blockchain, err  := ethclient.Dial(c.Server)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(c.Key), c.Password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	address, tx, _, err:= metacoin.DeployCoinCaller(auth,blockchain)

	if err != nil {
    log.Fatalf("Failed to deploy new trigger contract: %v", err)
  }
	fmt.Printf("Contract address deploy:0x%x\n", address)
	fmt.Printf("Transaction: ", tx.Hash())
}
