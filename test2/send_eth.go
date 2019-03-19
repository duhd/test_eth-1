package main

import (
	"context"
	"log"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/keystore"
	"io/ioutil"
	// "crypto/ecdsa"
	// crand "crypto/rand"
	"fmt"
	"math/big"

)
func main(){
		// connect to an ethereum node  hosted by infura
		client, err  := ethclient.Dial("http://localhost:8502")

		if err != nil {
			log.Fatalf("Unable to connect to network:%v\n", err)
		}

		file := "/Users/nguyenthanhbinh/Blockchain/test/ethereum/my-eth-chain-poa/node0/datadir/keystore/UTC--2019-02-11T11-39-39.949347000Z--eb80964e1567064ba810b45300fd2ce3193d1684"
    password := "123456"

    keyjson, err := ioutil.ReadFile(file)

    key, err := keystore.DecryptKey(keyjson, password)
    if err != nil {
        fmt.Println("json key failed to decrypt: %v", err)
    }

    fmt.Println("Private Key: ", key.PrivateKey)
		fmt.Printf("Private Key: %x", crypto.FromECDSA(key.PrivateKey))
		privateKey :=  key.PrivateKey
		// privateKey, err := crypto.HexToECDSA("b53c25d092b3eb50059b52b983f73c2fb36838ea4c69f372976dcada11fa8dff")
		// privateKey, err := crypto.LoadECDSA("/Users/nguyenthanhbinh/Blockchain/test/ethereum/my-eth-chain-poa/node0/datadir/keystore/UTC--2019-02-11T11-39-39.949347000Z--eb80964e1567064ba810b45300fd2ce3193d1684")
		if err != nil {
			log.Fatal(err)
		}

		// publicKey := privateKey.Public()
		// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		// if !ok {
		// 	log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		// }
		//
		// fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

		fmt.Println("Address: ", fromAddress.Hex(), ",Hash: ",fromAddress.Hash())
		// fromAddress := common.HexToAddress("0xeb80964e1567064ba810b45300fd2ce3193d1684")
		//
		// balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		//
		// fmt.Println("Balance: ",balance) // 25893180161173005034

		// pendingBalance, err := client.PendingBalanceAt(context.Background(), fromAddress)
		// fmt.Println("pendingBalance:", pendingBalance) // 25729324269165216042

		nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}


		value := big.NewInt(10000) // in wei (1 eth)
		gasLimit := uint64(21000)                // in units
		// gasPrice, err := client.SuggestGasPrice(context.Background())
		// fmt.Println("gasPrice: ",gasPrice)
		gasPrice := big.NewInt(10000000000000000) // in wei (30 gwei
		if err != nil {
			log.Fatal(err)
		}


		//Create a transaction
		toAddress := common.HexToAddress("0xd95f832f5296037df962ad33da618cbf0a52e192")
		var data []byte

		tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

		//Sign the transaction
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}


		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}
		//Send transaction to blockchain
		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
