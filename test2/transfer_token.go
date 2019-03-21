package main

import (
	"os"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	// "strings"
	"bytes"
	"io/ioutil"
	"fmt"
	"math/big"
	"time"
)


func main(){
	if len(os.Args) <7 {
		 fmt.Println("Please use syntax: go run transfer_token.go keyfile  websocket password contractAddr recvAddr amount note")
		 return
	}
	keyfile := os.Args[1]
	websocket := os.Args[2]
	password := os.Args[3]
	contractAddr :=  os.Args[4]
	recvAddr := os.Args[5]
	amount := os.Args[6]
	append := os.Args[7]


	start := time.Now()
	fmt.Println("Start transfer token")
	keyjson, err := ioutil.ReadFile(keyfile)

	auth, err := bind.NewTransactor(bytes.NewReader(keyjson), password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//Create transation
	//auth.GasPrice = big.NewInt(1)
	//auth.GasLimit = 100000000000000

	client, err  := ethclient.Dial(websocket)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	contractAddress := common.HexToAddress(contractAddr)
	instance, err := contracts.NewVNDWallet(contractAddress, client)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	address := common.HexToAddress(recvAddr)
	value := new(big.Int)
	value, ok := value.SetString(amount, 10)
	 if !ok {
			 fmt.Println("SetString: error")
			 return
	 }

	note :=  fmt.Sprintf("Transaction:  %s", append)
	tx, err := instance.Transfer(auth, address, value, []byte(note))
	if err != nil {
			log.Fatalf(" Transaction create error: ", err)
	}
	fmt.Println(" Transaction =",tx.Hash().Hex())

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Println("All times: ",elapsed)

}
