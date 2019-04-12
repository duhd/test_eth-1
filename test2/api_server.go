package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
  // "strings"
  // "github.com/go-redis/redis"
  // "encoding/json"
  "test_eth/test2/utils"
	// "github.com/savsgio/go-logger"
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

	 //Load all wallets in hosts
	 println("Create rpc connection pool ")
	 rpcrouting = utils.NewRouting()

	 println("Load wallets ")
 	 utils.LoadWallets(rpcrouting)

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
			ApiServer()
	 }()

	 wg.Wait()
	 fmt.Println("Finished webserver")
}

func ApiServer(){

	// router := fasthttprouter.New()
	router := routing.New()

  router.Get("/login", utils.Api_login_get)
	router.Post("/login", utils.Api_login)

	api := router.Group("/api/v1")

	if cfg.Jwt.Enable {
			fmt.Println("Using jwt")
	}

	api.Put("/cash/credit/<address>/<amount>/<traceid>", utils.JWTMiddleware(utils.Api_cash_credit))
	api.Put("/cash/debit/<address>/<amount>/<traceid>", utils.JWTMiddleware(utils.Api_cash_debit))
	api.Put("/cash/transfer/<from>/<to>/<amount>/<note>/<traceid>", utils.JWTMiddleware(utils.Api_cash_transfer))


	api.Get("/balance/<address>", utils.JWTMiddleware(utils.Api_balance))
	api.Get("/balance/all", utils.JWTMiddleware(utils.Api_balance_all))

	api.Post("/account/new", utils.JWTMiddleware(utils.Api_account_new))
	api.Get("/account/total", utils.JWTMiddleware(utils.Api_account_total))
	api.Get("/account/list/active", utils.JWTMiddleware(utils.Api_account_list_active))
	api.Get("/account/list/inactive", utils.JWTMiddleware(utils.Api_account_list_inactive))
	api.Get("/account/lock/<address>/<traceid>", utils.JWTMiddleware(utils.Api_account_lock))
	api.Get("/account/status/<address>", utils.JWTMiddleware(utils.Api_account_status))

	api.Get("/transaction/<txhash>",utils.JWTMiddleware(utils.Api_transaction))
	api.Get("/transaction/list/<account>/<fromdate>/<todate>", utils.JWTMiddleware(utils.Api_transaction_list))
	api.Get("/transaction/lock/<account>/<fromdate>/<todate>", utils.JWTMiddleware(utils.Api_transaction_lock))

	server := &fasthttp.Server{
		Name:    "JWT API Server",
		Handler: router.HandleRequest,
	}

	fmt.Println("Start listening")

	if cfg.Webserver.Tls {
		fmt.Println("Start server using TLS ")
		panic(server.ListenAndServeTLS(":"+ cfg.Webserver.Port,cfg.Webserver.CertificateFile,cfg.Webserver.KeyFile))
	} else {
		fmt.Println("Start server without TLS  ")
		panic(server.ListenAndServe(":"+ cfg.Webserver.Port))
	}
}
