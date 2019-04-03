package utils

import (
  "github.com/go-redis/redis"
  "fmt"
  "encoding/json"
  "time"
  "sync"
)

type Transaction struct {
        Id                string  `json:"Id"`
        CoinBase          string  `json:"Coinbase"`
        TxNonce           uint64  `json:"TxNonce"`
        RequestTime       int64   `json:"RequestTime"`
        TxReceiveTime     int64   `json:"TxReceiveTime"`
        TxConfirmedTime   []int64 `json:"TxConfiredTime"`
}


type RedisPool struct {
   Clients []*redis.Client
   Current int
   StartTxCh chan *Transaction
   EndTxCh chan *Transaction
   mutex sync.Mutex
}

var Rclients *RedisPool

func NewRedisPool() *RedisPool{
  max_connection := cfg.Redis.MaxConn

  clients := []*redis.Client{}
  for i:=0 ; i<max_connection; i++ {
        //Creat redis connection
        cl := redis.NewClient(&redis.Options{
          Addr:     cfg.Redis.Host,
          Password: cfg.Redis.Password, // no password set
          DB:       cfg.Redis.Db,  // use default DB
        })
       clients = append(clients,cl)
   }
   startTxCh := make(chan *Transaction,cfg.Channel.LogQueue)
   endTxCh := make(chan *Transaction,cfg.Channel.LogQueue)

   Rclients =  &RedisPool{
        Clients:clients,
        Current:0,
        StartTxCh: startTxCh,
        EndTxCh: endTxCh,
   }
   return Rclients
}
func (rp *RedisPool) Process() {
  for {
      select {
            case  tx:= <- rp.StartTxCh:
              go func() {
                fmt.Println("Write transation:",tx.Id, " to redis")
                client := Rclients.getClient()
                value, err := json.Marshal(tx)
                if err != nil {
                    fmt.Println(err)
                }
                err = client.Set("transaction:" + tx.Id,string(value), 0).Err()
                if err != nil {
                  fmt.Println(time.Now()," Write transaction to redis error: ", err)
                }
              }()
            case  tx:= <- rp.EndTxCh:
              go func(){
                fmt.Println("Get transation:",tx.Id, " from redis")
                client := Rclients.getClient()
                val, err2 := client.Get("transaction:" + tx.Id).Result()
                if err2 != nil {
                    fmt.Println(time.Now()," Cannot find transaction: ", tx.Id)
                    return
                }
                data := &Transaction{}
                err := json.Unmarshal([]byte(val), data)
                if err != nil {
                    fmt.Println(time.Now()," Cannot parse data ", err)
                    return
                }
                data.CoinBase = tx.CoinBase
                TxConfirmedTime := time.Now().UnixNano()
                data.TxConfirmedTime = append(data.TxConfirmedTime,TxConfirmedTime )

                fmt.Println("Update transation:",tx.Id, " to redis")
                value, err := json.Marshal(data)
                err = client.Set("transaction:" + tx.Id,string(value), 0).Err()
                if err != nil {
                  fmt.Println(time.Now()," Cannot update transaction: ",tx.Id,",Error:", err)
                }
                //
                // if data.TxNonce != nonce {
                //   fmt.Println(time.Now()," nonce:",data.TxNonce," tx:",key," request:",data.RequestTime," receive:", data.TxReceiveTime, " error:",nonce)
                // }

                time_receive_ms := (data.TxReceiveTime - data.RequestTime)/1000000
                time_confirm_ms := (TxConfirmedTime  - data.RequestTime )/1000000
                fmt.Println(time.Now()," nonce:",data.TxNonce," tx:",data.Id," request:",data.RequestTime," receive:", time_receive_ms, " confirm:",time_confirm_ms)
              }()
        }
    }
}


func (rp *RedisPool) getClient() *redis.Client {
  rp.mutex.Lock()
  defer rp.mutex.Unlock()

  len := len(rp.Clients)
  if rp.Current >= len {
      rp.Current =  rp.Current % len
  }
  fmt.Println("Current Redis connect: ",rp.Current," in ",len)
  client := rp.Clients[rp.Current]
  rp.Current = rp.Current + 1

  return client
}

func (rp *RedisPool) LogStart(key string, nonce uint64, requesttime int64) bool {
    trans :=  &Transaction{
                Id: key,
                TxNonce: nonce,
                RequestTime: requesttime,
                TxReceiveTime: time.Now().UnixNano()}
    rp.StartTxCh <- trans
    return true
}

func (rp *RedisPool) LogEnd(key string, nonce uint64, coinbase string) bool {
    trans :=  &Transaction{
                Id: key,
                TxNonce: nonce,
                CoinBase: coinbase,
                }
    rp.EndTxCh <- trans
    return true
}
