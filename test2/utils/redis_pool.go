package utils

import (
  "github.com/go-redis/redis"
)

type RedisPool struct {
   Clients []*redis.Client
   Current int
}

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
   Rclients =  &RedisPool{
      Clients:clients,
      Current:0,
   }
   return Rclients
}
func (rp *RedisPool) getClient() *redis.Client {
  client := rp.Clients[rp.Current]
  rp.Current = rp.Current + 1
  rp.Current = rp.Current % len(rp.Clients)
  return client 
}

var Rclients *RedisPool
