package main

import (
	"os"
	"strings"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/p2p"
	"fmt"
)
func main(){
	if len(os.Args) <1 {
		 fmt.Println("Please use syntax: go run addpeers.go  webservers")
		 return
	}
	webservers := os.Args[1]
	hosts := strings.Split(webservers,",")

  ctx := context.Background()
	method := "admin_nodeInfo"

	clients :=  make([]*rpc.Client,0)
	enodes := make([]string,0)
	for _, url := range hosts {
		fmt.Println("Connect host: ",url)
		var result p2p.NodeInfo
		client, err := rpc.DialContext(ctx, url)
	  if err != nil {
	    log.Fatalf("Unable to connect to network:%v\n", err)
	  }
		client.CallContext(ctx, &result, method)
		fmt.Printf("Enode: ", result.Enode)

		 enodes = append(enodes,result.Enode)
		 clients = append(clients,client)
	}

	fmt.Println("Start add peers for each client ")
	method = "admin_addPeer"
	for id, client := range clients {
		 for id1, enode := range enodes {
			 if id != id1 {
				 var result  bool
				 params := strings.Replace(enode,"127.0.0.1:30301?discport=0","172.101.0.17:30301",-1)
				 params = strings.Replace(params,"127.0.0.1:30302?discport=0","172.101.0.18:30302",-1)
				 client.CallContext(ctx, &result, method,params)
				 if result {
					 	 fmt.Printf("Add peer: ", params)
				 }else{
					 	 fmt.Printf("Failed adding peer: ", params)
				 }

			 }
		 }
	}










}
