package main

import (
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8546/ws")
	if err != nil {
		log.Fatal(err)
	}

	tx_hash := common.HexToHash("0xb9a6bfd71501b9129a6ec128f1335748e50c3ac1cee08e2cbe52c7d8ea0b18e3")
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
