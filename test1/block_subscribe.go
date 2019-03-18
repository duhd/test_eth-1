package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("ws://localhost:8546/ws")
	if err != nil {
		log.Fatal(err)
	}

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
    fmt.Println("Errror: ",err)
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
        fmt.Println("Errror blockbyhash: ",err)
				log.Fatal(err)
			}
			fmt.Println("time:",block.Time().Uint64(),"hex: ", block.Hash().Hex(),", number:",block.Number().Uint64(),",None:", block.Nonce(),",number of transaction",len(block.Transactions()))
		}
	}
}
