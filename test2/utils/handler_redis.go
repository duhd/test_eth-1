package utils

import (
  "time"
  // "math/big"
  // "strings"
   "fmt"
  "encoding/json"
  // "errors"
  // "strings"
  // "github.com/ethereum/go-ethereum/crypto"
    // _ "github.com/jinzhu/gorm/dialects/mysql"
    // "encoding/hex"
)

type RedisHandler struct {
    redisCache *RedisPool
}

func NewRedisHandler(rcache *RedisPool)  *RedisHandler{
      return &RedisHandler{
        redisCache:rcache,
      }
}
func (rh *RedisHandler) Report() string {
      client := rh.redisCache.getClient()
      keys, err  := client.Keys("transaction:*").Result()
      if err != nil {
        // handle error
        fmt.Println(time.Now()," Cannot get keys ")
      }
      vals, err1 := client.MGet(keys...).Result()
      if err1 != nil {
        // handle error
        fmt.Println(time.Now()," Cannot get values of  keys: ", keys)
      }

      fmt.Println("Elements: ", len(keys))
      diff_arr1 := []int64{}
      diff_arr := []int64{}

      for _, element := range vals {
          data := &Transaction{}
          err2 := json.Unmarshal([]byte(element.(string)), data)
          if err2 != nil {
              fmt.Println(time.Now()," Element:", element, ", Error:", err2)
              continue
          }
          fmt.Println("ID:",data.Id,"RequestTime:",data.RequestTime,
            "TxReceiveTime:",data.TxReceiveTime,"TxConfirmedTime:",data.TxConfirmedTime)

          var max int64 = 0
          if data.TxConfirmedTime != nil {
              for _,value := range data.TxConfirmedTime {
                  if value > max {
                     max = value
                  }
              }
              diff1 := data.TxReceiveTime - data.RequestTime
              diff_arr1 = append(diff_arr1,diff1)
          }
          // else {
          //     max = time.Now().UnixNano()
          // }
          if max >0 {
              diff := max  - data.TxReceiveTime
              diff_arr = append(diff_arr,diff)
          }
      }
      var total1 int64 = 0
    	for _, value1:= range diff_arr1 {
    		total1 += value1
    	}
      len1 := int64(len(diff_arr1))
      var avg1 int64 = 0
      if len1 >0 {
        	avg1 = total1/(len1 *1000)
      }

      var total int64 = 0
    	for _, value:= range diff_arr {
    		total += value
    	}
      len2 := int64(len(keys))
      len := int64(len(diff_arr))
      var avg int64 = 0
      if len >0 {
        	avg = total/(len *1000)
      }
      return fmt.Sprintf("Total Tx: %v , Total Complete TX: %v ,Avg RequestTime: %v , Avg Onchain: %v ", len2, len,avg1, avg)
}
