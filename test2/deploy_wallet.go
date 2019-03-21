package main

import (
	"os"
	"strings"
	"bytes"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"path/filepath"
	"io/ioutil"
	"fmt"
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
func loadConfig(file string) *Config {
     var cfg  Config

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err   #%v ", err)
     }

     err = yaml.Unmarshal(yamlFile, &cfg)
     if err != nil {
         fmt.Println("Unmarshal: %v", err)
     }
     return &cfg
}

func loadKey(root string,addr string) []byte {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
         panic(err)
    }
    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) == 3 {
					   account := list[2]
						 if account == strings.TrimPrefix(addr,"0x") {
							 keyjson, err := ioutil.ReadFile(file)
							 if err != nil {
										fmt.Println("Error in read file: ", file )
										return nil
							 }
							 return keyjson
						 }
         }
    }
		return nil
}
func main(){
				config_file := "config.yaml"
				if len(os.Args) == 2 {
					 config_file = os.Args[1]
				}
				cfg := loadConfig(config_file)
				// fmt.Println("Config", cfg)

				node := cfg.Networks[0]

				blockchain, err  := ethclient.Dial("http://"+node.Http)

				if err != nil {
					log.Fatalf("Unable to connect to network:%s with %v\n",node.Http, err)
				}

				keyjson := loadKey(cfg.Keys.Keystore,cfg.Contract.Owner)

				auth, err := bind.NewTransactor(bytes.NewReader(keyjson), cfg.Keys.Password)
				if err != nil {
					log.Fatalf("Failed to create authorized transactor: %v", err)
				}

				//triggerAddr, _, trigger, err := DeployTrigger(auth, backends.NewRPCBackend(conn))
				var initialSupply *big.Int = big.NewInt(cfg.Contract.InitialToken)
				tokenName := "Vietnam Dong"
				decimalUnits := uint8(1)
				tokenSymbol := "VND"
				initMasterKey1 := common.HexToAddress(cfg.Contract.MasterKey1)
				initMasterKey2 :=  common.HexToAddress(cfg.Contract.MasterKey2)

				address, tx, _, err:= contracts.DeployVNDWallet(auth,blockchain,initialSupply, tokenName,
					 decimalUnits , tokenSymbol , initMasterKey1, initMasterKey2)

				if err != nil {
			    log.Fatalf("Failed to deploy new trigger contract: %v", err)
			  }
				fmt.Println("Transaction: ", tx.Hash())
				fmt.Println("Contract address deploy:", address.Hex())
				cfg.Contract.Address = address.Hex()
				newcfg, err1 := yaml.Marshal(&cfg)
			  if err1 != nil {
			      fmt.Println("yaml.Marshal error: %v", err)
			  }
			  fmt.Printf("---\n%s", string(newcfg))
				err = ioutil.WriteFile(config_file, newcfg, 0644)
				if err != nil {
					fmt.Println("Write file error:",err)
				}
}
