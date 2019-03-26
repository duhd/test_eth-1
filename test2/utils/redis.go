package utils

import (
  "os"
  "fmt"
  "io/ioutil"
  "strings"
  "path/filepath"
  "encoding/json"
  "time"
  "github.com/go-redis/redis"
  "strconv"
)

type Transaction struct {
        Id                string  `json:"Id"`
        RequestTime       int64   `json:"RequestTime"`
        TxReceiveTime     int64   `json:"TxReceiveTime"`
        TxConfirmedTime    []int64 `json:"TxConfiredTime"`
   }

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
         }
    }
}

func LogStart(key string,requesttime int64){
  trans :=  &Transaction{
              Id: key,
              RequestTime: requesttime,
              TxReceiveTime: time.Now().UnixNano()}
  value, err := json.Marshal(trans)
  if err != nil {
      fmt.Println(err)
      return
  }
  err = Redis_client.Set("transaction:" + key,string(value), 0).Err()
	if err != nil {
		panic(err)
	}
}

func LogEnd(key string){
      val, err2 := Redis_client.Get("transaction:" + key).Result()
      if err2 != nil {
          fmt.Println("Cannot find transaction: ", key)
          return
      }
      data := &Transaction{}
      err := json.Unmarshal([]byte(val), data)
      if err != nil {
          fmt.Println("Cannot parse data ", err)
          return
      }
      data.TxConfirmedTime = append(data.TxConfirmedTime, time.Now().UnixNano())
      value, err := json.Marshal(data)

      err = Redis_client.Set("transaction:" + key,string(value), 0).Err()
    	if err != nil {
    	     fmt.Println("Cannot set data ", err)
    	}
      fmt.Println("Finish write transaction: ", key)
}
func GetNonce(account string) uint64 {
  val, err := Redis_client.Get("nonce:" + account).Result()
  if err != nil {
      fmt.Println("Cannot find nonce of account: ", account)
      return uint64(0)
  }
  value , err := strconv.ParseUint(val, 10, 64)
  if err != nil {
      fmt.Println("Cannot parce nonce of ", val)
      return uint64(0)
  }
  return value
}
func CommitNonce(account string, nonce uint64) bool {
  err := Redis_client.Set("nonce:" + account,uint64(nonce), 0).Err()
  if err != nil {
       fmt.Println("Cannot set nonce  ", err)
       return false
  }
  return true
}
func NoneIncr(account string) bool {
  _, err := Redis_client.Incr("nonce:" + account).Result()
	if err != nil {
    fmt.Println("Cannot increase nonce  ", err)
    return false
	}
  return true
}
