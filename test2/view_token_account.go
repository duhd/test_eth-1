package main

import (
	"os"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"math/big"
	// "time"
	// "os"
)

//const key  = `paste the contents of your JSON key file here`
// const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
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
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

  instance, err := contracts.NewVNDWallet(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatalf("Unable to bind to deployed instance of contract:%v\n")
	}


	// address := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
	address := common.HexToAddress(accountaddr)


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
