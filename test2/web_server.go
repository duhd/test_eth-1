package main

import (
        // "strconv"
        // "math/rand"
        // "crypto/sha1"
        // "encoding/base64"
        // "hash"
        "os"
        // "strings"
       "github.com/gin-gonic/gin"
       // "github.com/go-redis/redis"
        // "net/http"
        // "encoding/json"
        "fmt"
        "test_eth/test2/utils"
        "sync"

)

var cfg *utils.Config
var rpcrouting *utils.RpcRouting
var redisCache *utils.RedisPool

func init() {
   config_file := "config.yaml"
   if len(os.Args) == 2 {
       config_file = os.Args[1]
    }

    println("init function")
    cfg = utils.LoadConfig(config_file)
}

func main() {
  //Creat redis connection
	println("Initialize redis")
	redisCache := utils.NewRedisPool()

	println("Delete old data in redis ")
	//utils.DeleteData("transaction*")
	//utils.DeleteData("nonce*")

	// sha = sha1.New()
  //Load all wallets in hosts
  println("Create rpc connection pool ")
  rpcrouting = utils.NewRouting()

	 var wg sync.WaitGroup
   wg.Add(3)

	 go func (){
				println("Loop Routing process message ")
				defer wg.Done()
				rpcrouting.Process()
	 }()

	 go func (){
		 	  println("Loop webservice ")
			  defer wg.Done()
				redisCache.Process()
	 }()

	 go func (){
			 println("Loop webservice ")
			defer wg.Done()
			httpServer()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}
func httpServer(){
  router := gin.Default()
  // Simple group: v1

  gin_api := utils.NewApiGin(cfg,rpcrouting,redisCache)

  v1 := router.Group("/api/v1")
  {
      v1.GET("/wallet/:method/:p1/:p2/:p3/:p4", gin_api.ProcessCall)
      v1.GET("/wallet/:method/:p1", gin_api.ProcessCall)
      v1.GET("/wallet/:method", gin_api.ProcessCall)
   }
   router.Run(":"+ cfg.Webserver.Port)
}
