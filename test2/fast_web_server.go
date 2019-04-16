package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
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
	redisCache = utils.NewRedisPool()

	println("Delete old data in redis ")
	//utils.DeleteData("transaction*")
	//utils.DeleteData("nonce*")

	 //Load all wallets in hosts
	 println("Create rpc connection pool ")
	 rpcrouting = utils.NewRouting(cfg)


	 var wg sync.WaitGroup

	 wg.Add(3)

	 go func (){
				println("Loop Routing process message ")
				defer wg.Done()
				rpcrouting.Process()
	 }()

	 go func (){
		 	  println("Loop redisPool ")
			  defer wg.Done()
				redisCache.Process()
	 }()

	 go func (){
			 println("Loop httpServer ")
			defer wg.Done()
			httpServer()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}

func httpServer(){
	router := routing.New()

	//
	// api_v1 := router.Group("/api/v1/wallet")
	// fast_api := utils.NewApiFast(cfg,rpcrouting,redisCache)
  // api_v1.Get("/<method>/<p1>/<p2>/<p3>/<p4>/<p5>/<p6>", fast_api.ProcessCall)
	// api_v1.Get("/<method>/<p1>/<p2>/<p3>/<p4>/<p5>", fast_api.ProcessCall)
	// api_v1.Get("/<method>/<p1>/<p2>/<p3>/<p4>", fast_api.ProcessCall)
	// api_v1.Get("/<method>/<p1>/<p2>/<p3>", fast_api.ProcessCall)
	// api_v1.Get("/<method>/<p1>/<p2>",  fast_api.ProcessCall)
	// api_v1.Get("/<method>/<p1>", fast_api.ProcessCall)
	// api_v1.Get("/<method>",  fast_api.ProcessCall)

	api_v2 := router.Group("/api/v2/wallet")
	fast_api_v2 := utils.NewApiFastV2(cfg,rpcrouting,redisCache)
	api_v2.Get("/<method>/<p1>/<p2>/<p3>/<p4>/<p5>/<p6>", fast_api_v2.ProcessCall)
	api_v2.Get("/<method>/<p1>/<p2>/<p3>/<p4>/<p5>", fast_api_v2.ProcessCall)
	api_v2.Get("/<method>/<p1>/<p2>/<p3>/<p4>", fast_api_v2.ProcessCall)
	api_v2.Get("/<method>/<p1>/<p2>/<p3>", fast_api_v2.ProcessCall)
	api_v2.Get("/<method>/<p1>/<p2>",  fast_api_v2.ProcessCall)
	api_v2.Get("/<method>/<p1>", fast_api_v2.ProcessCall)
	api_v2.Get("/<method>",  fast_api_v2.ProcessCall)

	fmt.Println("Start listening")
	panic(fasthttp.ListenAndServe(":"+ cfg.Webserver.Port, router.HandleRequest))
}
