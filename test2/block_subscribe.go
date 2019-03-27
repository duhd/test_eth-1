package main

import (
	"os"
	// "strings"
	"context"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
  "test_eth/test2/utils"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var cfg *utils.Config

func main() {
	var wg sync.WaitGroup

	 config_file := "config.yaml"
	 if len(os.Args) == 2 {
			config_file = os.Args[1]
	 }

	 	cfg = utils.LoadConfig(config_file)

		//Creat redis connection
		utils.Redis_client = redis.NewClient(&redis.Options{
			Addr:     cfg.Redis.Host,
			Password: cfg.Redis.Password, // no password set
			DB:       cfg.Redis.Db,  // use default DB
		})

		len := len(cfg.Networks)
		wg.Add(len)

		for _,host := range cfg.Networks {
			  go func(httpUrl string, socketUrl string ){
						defer wg.Done()
						listening_block_from_host(httpUrl,socketUrl)
				}(host.Http,host.WebSocket)
	 	}
	  wg.Wait()
}
func listening_block_from_host(httpUrl string, socketUrl string ){
		fmt.Println("Listening from: ", socketUrl)
		websocket, err := ethclient.Dial("ws://" + socketUrl)
		if err != nil {
				log.Fatal("Cannot connect to websocket", err)
				return
		}

		max_client := cfg.Webserver.MaxListenRpcConnection

		var clients []*utils.EthClient
		var current int = 0

		for i:=0; i< max_client; i++ {
			ethclient, err := utils.NewEthClient(httpUrl)
			if err != nil {
				log.Fatal("Cannot connect to: ",httpUrl," error:", err)
				continue
			}
			clients = append(clients,ethclient)
		}

		headers := make(chan *types.Header)
		sub, err := websocket.SubscribeNewHead(context.Background(), headers)
		if err != nil {
		    fmt.Println("Cannot SubscribeNewHead to host: ", socketUrl ," Error: ",err)
				return
		}
	  fmt.Println("Start listening from ",socketUrl," : ")
		for {
					select {
								case err := <-sub.Err():
										fmt.Println("Error from: ",socketUrl," Error: ",err)
										log.Fatal(err)
								case header := <-headers:
										t := time.Now()
										fmt.Println(t.Format(time.RFC822),"Block Number: ", header.Number.String(), " header hash: " , header.Hash().Hex())
										client := clients[current]
										go func(){
												client.UpdateReceipt(header)
										}()
										current = current + 1
										current = current % len(clients)
						}
		}
}
