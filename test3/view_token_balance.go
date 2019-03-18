package main

import (
	"log"
	"test_eth/contracts/metacoin"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"math/big"
)
func main(){
	// connect to an ethereum node  hosted by infura
	client, err  := ethclient.Dial("http://localhost:8502")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	contractAddress := "0xAE6313a252d905cdc0d8e9116fE31696CC832145"
  instance, err := metacoin.NewMetaCoin(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}


	// address := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
	// accountAddr := "0xeb80964e1567064ba810b45300fd2ce3193d1684"
	accountAddr := "d95f832f5296037df962ad33da618cbf0a52e192"
	address := common.HexToAddress(accountAddr)

	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)
	fbal.SetString(bal.String())
	fmt.Printf("balance: %f", bal) // "balance: 74605500.647409"

	// value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(18))))
	// fmt.Printf("balance: %f", value) // "balance: 74605500.647409"
}
