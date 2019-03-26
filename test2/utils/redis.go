package utils

import (
  "os"
  "fmt"
  "io/ioutil"
  "strings"
  "path/filepath"
  // "encoding/json"
  // "time"
  "github.com/go-redis/redis"
  // "strconv"
  "github.com/ethereum/go-ethereum/crypto"
  "github.com/ethereum/go-ethereum/accounts/keystore"
)


var Redis_client *redis.Client

func DeleteData(pattern string){
  keys, err  := Redis_client.Keys(pattern).Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get keys ")
  }
  res := Redis_client.Del(keys...)
  fmt.Println("Redis delete: ", res)
}
func LoadKeyStores(root string){
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
             account := "account:" + list[2]
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }
             //Set key in redis
              err = Redis_client.Set(account,string(keyjson), 0).Err()
              if err != nil {
                panic(err)
              }

              accountKey, err := keystore.DecryptKey( []byte(keyjson), cfg.Keys.Password)
              if err != nil {
                  fmt.Println("Cannot decrypt key file: ", err)
                  return
              }
              privateKey := accountKey.PrivateKey

              private := "private:" + list[2]
              //Set key in redis
              err = Redis_client.Set(private,string(crypto.FromECDSA(privateKey)), 0).Err()
              if err != nil {
                 panic(err)
              }
         }
    }
}
