package main

import (
	"os"
	// "strings"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/rpc"
	// "github.com/ethereum/go-ethereum/p2p"
	"fmt"
	// "test_eth/test2/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	// "github.com/ethereum/go-ethereum/core/rawdb"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/core/vm"
	// "github.com/ethereum/go-ethereum/crypto"
	// "github.com/ethereum/go-ethereum/log"
	// "github.com/ethereum/go-ethereum/p2p"
	// "github.com/ethereum/go-ethereum/params"
	// "github.com/ethereum/go-ethereum/rlp"
	// "github.com/ethereum/go-ethereum/rpc"

)


type RPCTransaction struct {
	BlockHash        common.Hash     `json:"blockHash"`
	BlockNumber      *hexutil.Big    `json:"blockNumber"`
	From             common.Address  `json:"from"`
	Gas              hexutil.Uint64  `json:"gas"`
	GasPrice         *hexutil.Big    `json:"gasPrice"`
	Hash             common.Hash     `json:"hash"`
	Input            hexutil.Bytes   `json:"input"`
	Nonce            hexutil.Uint64  `json:"nonce"`
	To               *common.Address `json:"to"`
	TransactionIndex hexutil.Uint    `json:"transactionIndex"`
	Value            *hexutil.Big    `json:"value"`
	V                *hexutil.Big    `json:"v"`
	R                *hexutil.Big    `json:"r"`
	S                *hexutil.Big    `json:"s"`
}


func main(){
		if len(os.Args) != 2 {
			 fmt.Println("Please use syntax: go run txpool.go  server ")
			 return
		}
	  server := os.Args[1]
		ctx := context.Background()
		method := "txpool_content"
		client, err := rpc.DialContext(ctx, server)
		if err != nil {
			fmt.Println("Unable to connect to network:%v\n", err)
		}
		var result map[string]map[string]map[string][]*RPCTransaction
		client.CallContext(ctx, &result, method)

		for key, value := range result {
	    fmt.Println("Key:", key, "Value:", value)
		}
}
