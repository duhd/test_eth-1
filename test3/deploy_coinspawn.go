package main

import (
	"log"
	"test_eth/contracts/metacoin"
	util "test_eth/utils"
		// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
  // "github.com/ethereum/go-ethereum/rpc"
		// "math/big"
	"strings"
	"fmt"
	// "time"
	// "os"
)

//const key  = `paste the contents of your JSON key file here`
// const key  = `{"address":"d95f832f5296037df962ad33da618cbf0a52e192","crypto":{"cipher":"aes-128-ctr","ciphertext":"f999d122f6edf0c3664adb25a0cb5cd91405592f36518c42684ab7db9b565d4d","cipherparams":{"iv":"ef2f1eb65573db114d5c9e6f2ac5edd2"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"f6b2cddd480c5d496f1e786c1e3705dd6362b65e96201749eb5f7bd08232bb46"},"mac":"e7111e5645875bdc1f8a21f6a33aa318c34a0df6f49c5007c427c05987dfbd85"},"id":"9cae0855-92f6-4e35-9ca1-4544a6d66b52","version":3}`
// const key  = `{"address":"ffbcd481c1330e180879b4d2b9b50642eea43c02","crypto":{"cipher":"aes-128-ctr","ciphertext":"351950aa30a37e4b385ae27ff2139c4151a6021333bd986602e80c2288f9e8fe","cipherparams":{"iv":"aec5c52378134e49a6037a5b77bec309"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8b0640866e9dbbba9f4a5da4348905b6f332a1b44a614ceadd0e9bd4ea7cdd7d"},"mac":"247c67172dbcdc48f031394ef1a25547f720b769be433a382a48028137f34002"},"id":"e52f52a5-cea4-459d-9e9b-ad8c76d7a562","version":3}`
// const rpc_server = "http://localhost:8502"
// const password = "123456"


func main(){
	var c util.Config
	c.GetConf()

	fmt.Println(c)

	// connect to an ethereum node  hosted by infura
	blockchain, err  := ethclient.Dial(c.Server)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}
	// Get credentials for the account to charge for contract deployments
	auth, err := bind.NewTransactor(strings.NewReader(c.Key), c.Password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	address, tx, _, err:= metacoin.DeployCoinSpawn(auth,blockchain)

	if err != nil {
    log.Fatalf("Failed to deploy new trigger contract: %v", err)
  }
	fmt.Printf("Contract address deploy:0x%x\n", address)
	fmt.Printf("Transaction: ", tx.Hash())
}
