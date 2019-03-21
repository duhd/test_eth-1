package main

import (
	"os"
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// connect to an ethereum node  hosted by infura
	if len(os.Args) <2 {
		 fmt.Println("Please use syntax: go run get_transaction_state.go  webserver  transaction_hash")
		 return
	}
	webserver := os.Args[1]
	transHash := os.Args[2]

	client, err  := ethclient.Dial(webserver)

	if err != nil {
		log.Fatal(err)
	}

	tx_hash := common.HexToHash(transHash)
	tx, isPending, err := client.TransactionByHash(context.Background(), tx_hash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction status:",isPending,", hash: ", tx.Hash()) // 1

	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("GasUsed:",receipt.GasUsed) // 1
}
