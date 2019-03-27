package main

import (
	"os"
	"strings"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/p2p"
	"fmt"
	"test_eth/test2/utils"
)

var cfg *utils.Config

func main(){
	if len(os.Args) == 1 {
		 fmt.Println("Please use syntax: go run txpool.go  server ")
		 server = os.Args[1]
	}
	ctx := context.Background()
	method := "txpool_content"
	client, err := rpc.DialContext(ctx, server)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	client.CallContext(ctx, &result, method)
	enode := strings.Replace(result.Enode,"127.0.0.1",node.LocalAddr,-1)
	enode = strings.Replace(enode,"?discport=0","",-1)
	fmt.Printf("Enode: ", enode)
	enodes = append(enodes,enode)
	clients = append(clients,client)
}
