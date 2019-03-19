package main

import (
	"os"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"math"
	"math/big"
)

func main(){
	if len(os.Args) <2 {
		 fmt.Println("Please use syntax: go run view_eth_balance.go  webserver  account")
		 return
	}
	webserver := os.Args[1]
	accountaddr := os.Args[2]

	client, err  := ethclient.Dial(webserver)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

//	account := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
	account := common.HexToAddress(accountaddr)

	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance: ",balance) // 25893180161173005034

	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue:", ethValue) // 25.729324269165216041
}
