package main

import (
	"os"
	"strings"
	"context"
	"log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/p2p"
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

type Config struct {
			Keys  struct {
				  Keystore string `yaml:"keystore"`
					Password string `yaml:"password"`
			} `yaml:"keys"`
			Networks []struct {
					Name string `yaml:"name"`
					Http string `yaml:"http"`
					WebSocket string `yaml:"websocket"`
					LocalAddr string `yaml:"local"`
			} `yaml:"networks"`
			Redis struct {
				  Host string `yaml:"host"`
				  Password string `yaml:"password"`
				  Db int `yaml:"db"`
			} `yaml:"redis"`
			Contract struct {
					Owner string `yaml:"owner"`
					InitialToken int64 `yaml:"initialToken"`
					MasterKey1 string `yaml:"masterkey1"`
					MasterKey2 string `yaml:"masterkey2"`
					Address string `yaml:"address"`
			} `yaml:"contract"`
	}

var cfg *Config

func loadConfig(file string) *Config {
     cfg := &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err   #%v ", err)
     }

     err = yaml.Unmarshal(yamlFile, cfg)
     if err != nil {
         fmt.Println("Unmarshal: %v", err)
     }
     return cfg
}

func main(){
	config_file := "config.yaml"
	if len(os.Args) == 2 {
		 fmt.Println("Please use syntax: go run addpeers.go  configfile ")
		 config_file = os.Args[1]
	}
	cfg = loadConfig(config_file)


  ctx := context.Background()
	method := "admin_nodeInfo"

	clients :=  make([]*rpc.Client,0)
	enodes := make([]string,0)

	for _, node := range cfg.Networks {
				fmt.Println("Connect host: ",node.Http)
				var result p2p.NodeInfo
				client, err := rpc.DialContext(ctx, "http://"+ node.Http)
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

	fmt.Println("Start add peers for each client ")
	method = "admin_addPeer"
	for id, client := range clients {
		 for id1, enode := range enodes {
			 if id != id1 {
				 var result  bool
				 client.CallContext(ctx, &result, method,enode)
				 if result {
					 	 fmt.Printf("Add peer: ", enode)
				 }else{
					 	 fmt.Printf("Failed adding peer: ", enode)
				 }
			 }
		 }
	}










}
