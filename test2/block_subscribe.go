package main

import (
	"os"
	"strings"
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
  "test_eth/test2/utils"
	"github.com/go-redis/redis"
)

func main() {
	if len(os.Args) <1 {
		 fmt.Println("Please use syntax: go run block_subscribe.go  webserver ")
		 return
	}
	webserver := os.Args[1]

	config_file := "config.yaml"
	cfg := utils.LoadConfig(config_file)

	//Creat redis connection
	utils.Redis_client = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password, // no password set
		DB:       cfg.Redis.Db,  // use default DB
	})


	client, err := ethclient.Dial(webserver)
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
			for _, transaction := range block.Transactions(){
				   fmt.Println("Transaction: ",transaction.Hash().Hex())
					 key := strings.TrimPrefix(transaction.Hash().Hex(),"0x")
					 utils.LogEnd(key)
			}
		}
	}
}
