package utils

import (
  "os"
  "fmt"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "strings"
  "path/filepath"
)

type Config struct {
    Channel struct {
        TransferQueue int `yaml:"transferqueue"`
        LogQueue int `yaml:"logqueue"`
    } `yaml:"channel"`
    Webserver struct {
			  Port string `yaml:"port"`
        MaxRpcConnection int `yaml:"maxrpc"`
        MaxListenRpcConnection int `yaml:"maxlistenrpc"`
        Mode int `yaml:"mode"`
		} `yaml:"webserver"`
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
        MaxConn int  `yaml:"maxconn"`
			  Host string `yaml:"host"`
			  Password string `yaml:"password"`
			  Db int `yaml:"db"`
		} `yaml:"redis"`
		Contract struct {
        GasPrice string `yaml:"gasprice"`
        GasLimit uint64 `yaml:"gaslimit"`
				Owner string `yaml:"owner"`
				InitialToken int64 `yaml:"initialToken"`
				MasterKey1 string `yaml:"masterkey1"`
				MasterKey2 string `yaml:"masterkey2"`
				Address string `yaml:"address"`
		} `yaml:"contract"`
}

var cfg *Config

func LoadConfig(file string) *Config {
     cfg = &Config{}

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

func LoadKey(root string,addr string) []byte {
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
