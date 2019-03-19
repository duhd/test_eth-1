package main

import (
	"os"
	// "log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"math/big"
)

func main(){
	if len(os.Args) <3 {
		 fmt.Println("Please use syntax: go run view_token_account.go  webserver  contractAddress accountaddr ")
		 return
	}
	webserver := os.Args[1]
	contractAddress := os.Args[2]
  accountaddr := os.Args[3]

	client, err  := ethclient.Dial(webserver)

	if err != nil {
		fmt.Println("Unable to connect to network:%v\n", err)
	}

  instance, err := contracts.NewVNDWallet(common.HexToAddress(contractAddress), client)
	if err != nil {
		fmt.Println("Unable to bind to deployed instance of contract:%v\n")
	}


	// address := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
	address := common.HexToAddress(accountaddr)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		fmt.Println("Get balanceof: ", err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	fmt.Printf("balance: %f", bal) // "balance: 74605500.647409"

	// value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	// fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
