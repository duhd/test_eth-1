package main

import (
		 // "context"
			"log"
			"test_eth/contracts/metacoin"
			"github.com/ethereum/go-ethereum/ethclient"
			"github.com/ethereum/go-ethereum/accounts/abi/bind"
			"github.com/ethereum/go-ethereum/common"
			// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
		  // "github.com/ethereum/go-ethereum/rpc"
			"strings"
			"fmt"
			"math/big"
			"time"
)

// const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`
// const ws_server = "ws://localhost:8546"
// const password = "123456"
// const spawnContractAddr = "0xffe2524adee706539e682152dcd5eefd819fd7cd"

func main(){
	var c config
	c.getConf()

	fmt.Println(c)

	start := time.Now()
	fmt.Println("Start transfer cash")


	fmt.Println("Start init token ")

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(c.key),c.password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//Create transation
	auth.GasPrice = big.NewInt(1)
	auth.GasLimit = 100000000000000

	// connect to an ethereum node  hosted by infura
	// client1, err  := ethclient.Dial("http://localhost:8502")
	client, err  := ethclient.Dial(c.webservice)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	contractAddress := common.HexToAddress(spawnContractAddr)

	instance, err := metacoin.NewCoinSpawn(contractAddress, client)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	initialBalance := big.NewInt(1000000000)

	tx,err := instance.CreateCoin(auth,initialBalance)
	if err != nil {
			log.Fatalf("Transaction create error: ", err)
	}
	fmt.Println("Transaction hash:",tx.Hash().Hex()) // 1
	fmt.Println("Finished")

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
