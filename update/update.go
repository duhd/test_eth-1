package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"eth/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	// "os"
)

const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`

func main(){
	// connect to an ethereum node  hosted by infura
	blockchain, err := ethclient.Dial("http://localhost:8502")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(key), "123456")

	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

  contract, err := contracts.NewInbox(common.HexToAddress("0xcc94f30BefDa41ae42F883e1Ae7f6291F1F3698F"), blockchain)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}
	contract.SetMessage(&bind.TransactOpts{
		From:auth.From,
		Signer:auth.Signer,
		Value: nil,
	}, "Hello From Mars")
	fmt.Println("End")
}
