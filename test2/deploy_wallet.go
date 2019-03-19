package main

import (
	"bytes"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strings"
	"fmt"
)
// const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`
// const webserver = "http://172.101.0.17:8501"
// const password = "123456"
// const masterkey1 = "0xbea868edea1c167aab5a0eef99496e2a690f3fae"
// const masterkey1 = "0x0e18db9aeea79d71b4c91c8375f1ef7fd0aaa594"

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
	keyjson, err := ioutil.ReadFile(file)


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
