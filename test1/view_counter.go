package main

import (
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
  // "github.com/ethereum/go-ethereum/rpc"
		"github.com/ethereum/go-ethereum/common"
	// "strings"
	"fmt"
	// "time"
	// "os"
)

//const key  = `paste the contents of your JSON key file here`
// const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
func main(){
	// connect to an ethereum node  hosted by infura
	client, err  := ethclient.Dial("http://localhost:8502")

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	contractAddress := "0xbb7dcbfb547171d1d619ba4f83fde03da3cbc5bc"
  instance, err := contracts.NewCounter(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}

	//Create transation
	count, err := instance.Count(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Error: " , err)
		return
	}
	fmt.Printf("count:", count) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
}
