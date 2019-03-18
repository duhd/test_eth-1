package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/core/types"
	// "github.com/ethereum/go-ethereum/crypto"
	// "crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

)

func main(){
		// connect to an ethereum node  hosted by infura
		client, err  := ethclient.Dial("http://localhost:8502")

		if err != nil {
			log.Fatalf("Unable to connect to network:%v\n", err)
		}

	//	account := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
		account := common.HexToAddress("0xd95f832f5296037df962ad33da618cbf0a52e192")

		balance, err := client.BalanceAt(context.Background(), account, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Balance: ",balance) // 25893180161173005034

		// pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
		// fmt.Println("pendingBalance:", pendingBalance/1000000000000000000) // 25729324269165216042
		fbalance := new(big.Float)
		fbalance.SetString(balance.String())
		ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
		fmt.Println("ethValue:", ethValue) // 25.729324269165216041
}
