package utils

import (
  "github.com/go-redis/redis"
  "fmt"
    "encoding/json"
  "time"
)

type RedisPool struct {
   Clients []*redis.Client
   Current int
   TxCh chan *Transaction
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
   txCh := make(chan *Transaction,100)
   Rclients =  &RedisPool{
        Clients:clients,
        Current:0,
        TxCh: txCh,
   }
   return Rclients
}
func (rp *RedisPool) getClient() *redis.Client {
  if rp.Current >= len(rp.Clients) {
      rp.Current =  rp.Current % len(rp.Clients)
  }
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
    rp.TxCh <- trans
    return true
}
func (rp *RedisPool) Loop() {
  for {
      select {
            case  tx:= <- rp.TxCh:
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

        }
    }
}
