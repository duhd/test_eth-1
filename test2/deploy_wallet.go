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
	"gopkg.in/yaml.v2"
)

type Config struct {
     Keystore string `json:"keystore"`
     Server string `yaml:"server"`
     Webservice string `yaml:"webservice"`
     Password string `yaml:"password"`
     RedisHost string `yaml:"redis_host"`
     RedisPassword string `yaml:"redis_password"`
     RedisDb int `yaml:"redis_db"`
     ContractAddr string `yaml:"contract_address"`
		 KeyFile string `yaml:"keyfile"`
     MasterKey1 string `yaml:"masterkey1"`
     MasterKey2 string `yaml:"masterkey2"`
		 InitialToken int64 `yaml:"initialToken"`

}
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
		 config_file = os.Args[1]
	}
	cfg := loadConfig(config_file)

	blockchain, err  := ethclient.Dial(cfg.Server)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	// Get credentials for the account to charge for contract deployments
	keyfile := cfg.Keystore + "/" + cfg.KeyFile
	keyjson, err := ioutil.ReadFile(keyfile)

	auth, err := bind.NewTransactor(bytes.NewReader(keyjson), cfg.Password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	//triggerAddr, _, trigger, err := DeployTrigger(auth, backends.NewRPCBackend(conn))
	var initialSupply *big.Int = big.NewInt(cfg.InitialToken)
	tokenName := "Vietnam Dong"
	decimalUnits := uint8(1)
	tokenSymbol := "VND"
	initMasterKey1 := common.HexToAddress(cfg.MasterKey1)
	initMasterKey2 :=  common.HexToAddress(cfg.MasterKey2)

	address, tx, _, err:= contracts.DeployVNDWallet(auth,blockchain,initialSupply, tokenName,
		 decimalUnits , tokenSymbol , initMasterKey1, initMasterKey2)

	if err != nil {
    log.Fatalf("Failed to deploy new trigger contract: %v", err)
  }
	fmt.Println("Transaction: ", tx.Hash())
	fmt.Println("Contract address deploy:", address.Hex())
	cfg.ContractAddr = address.Hex()
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
