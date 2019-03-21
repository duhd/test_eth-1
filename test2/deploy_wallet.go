package main

import (
	"os"
	"bytes"
	"log"
	"test_eth/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"io/ioutil"
	"fmt"
	"test_eth/test2/utils"
	"gopkg.in/yaml.v2"
)

func main(){
				config_file := "config.yaml"
				if len(os.Args) == 2 {
					 config_file = os.Args[1]
				}
				cfg := utils.LoadConfig(config_file)
				// fmt.Println("Config", cfg)

				node := cfg.Networks[0]

				blockchain, err  := ethclient.Dial("http://"+node.Http)

				if err != nil {
					log.Fatalf("Unable to connect to network:%s with %v\n",node.Http, err)
				}

				keyjson := utils.LoadKey(cfg.Keys.Keystore,cfg.Contract.Owner)

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
