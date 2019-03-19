package main

import (
	"os"
	"bytes"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"io/ioutil"
	"fmt"
)
func main(){
	if len(os.Args) <5 {
		 fmt.Println("Please use syntax: go run deploy_wallet.go  keyfile webserver password masterkey1 masterkey2")
		 return
	}

	keyfile := os.Args[1]
	webserver := os.Args[2]
	password := os.Args[3]
	masterkey1 := os.Args[4]
	masterkey2 := os.Args[5]

	blockchain, err  := ethclient.Dial(webserver)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	keyjson, err := ioutil.ReadFile(keyfile)


	auth, err := bind.NewTransactor(bytes.NewReader(keyjson), password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	//triggerAddr, _, trigger, err := DeployTrigger(auth, backends.NewRPCBackend(conn))
	var initialSupply *big.Int = big.NewInt(100000000000000)
	tokenName := "Vietnam Dong"
	decimalUnits := uint8(1)
	tokenSymbol := "VND"
	initMasterKey1 := common.HexToAddress(masterkey1)
	initMasterKey2 :=  common.HexToAddress(masterkey2)

	address, tx, _, err:= contracts.DeployVNDWallet(auth,blockchain,initialSupply, tokenName,
		 decimalUnits , tokenSymbol , initMasterKey1, initMasterKey2)

	if err != nil {
    log.Fatalf("Failed to deploy new trigger contract: %v", err)
  }
	fmt.Printf("Contract address deploy:0x%x\n", address)
	fmt.Printf("Transaction: ", tx.Hash())
}
