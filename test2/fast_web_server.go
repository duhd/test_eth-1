package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
  "strings"
  // "github.com/go-redis/redis"
  // "encoding/json"
  "test_eth/test2/utils"
)


var cfg *utils.Config

var clientPool *utils.ClientPool

func init() {
  config_file := "config.yaml"
  if len(os.Args) == 2 {
      config_file = os.Args[1]
   }

   println("init function")
   cfg = utils.LoadConfig(config_file)


	 //Creat redis connection
   println("Initialize redis")
   utils.NewRedisPool()

   println("Delete old data in redis ")
   //utils.DeleteData("transaction*")
   //utils.DeleteData("nonce*")

   // sha = sha1.New()
   println("Load key in account array ")
   utils.LoadKeyStores(cfg.Keys.Keystore)

    //Load all wallets in hosts
    println("Create rpc connection pool ")
    clientPool = utils.NewClientPool()

		//Sync nonce of account
		println("sync nonce of account from ethereum ")
		utils.SyncNonce(clientPool.GetClient().Client)

}

func main() {
	router := routing.New()
  api := router.Group("/api/v1/wallet")

	api.Get("/<method>/<p1>/<p2>/<p3>/<p4>", processCall)
  api.Get("/<method>/<p1>/<p2>/<p3>", processCall)
  api.Get("/<method>/<p1>/<p2>", processCall)
  api.Get("/<method>/<p1>", processCall)
  api.Get("/<method>", processCall)

  fmt.Println("Start listening")
	panic(fasthttp.ListenAndServe(":"+ cfg.Webserver.Port, router.HandleRequest))
}
// createTodo add a new todo
func processCall(c *routing.Context) error {
  method := c.Param("method")
  switch method {
      case "transfer":
           fmt.Println("call transfer")
           transfer(c)
           return  nil
       case "balance":
           fmt.Println("call balance")
           balance(c)
           return nil
       case "report":
           fmt.Println("call report")
           report(c)
           return nil
       case "accounts":
           fmt.Println("call accounts")
           accounts(c)
           return nil
       case "key":
           fmt.Println("call key")
           getKey(c)
           return nil
			 case "test":
           fmt.Println("call test")
           fmt.Fprintf(c, "data=test")
           return nil
   }

   fmt.Fprintf(c, "URL not found ")
   return nil
 }


 // call transfer token
 func transfer(c *routing.Context){

     from := c.Param("p1")
     to := c.Param("p2")
     amount := c.Param("p3")
     append := c.Param("p4")

     if from == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return
     }
     if to == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return
     }
     from = strings.TrimPrefix(from,"0x")
     to = strings.TrimPrefix(to,"0x")

  	 result, err := utils.TransferToken(from,to,amount,append)
     if err != nil {
           fmt.Println("Error to transfer token: ", err)
           fmt.Fprintf(c,"Error to transfer token: ", err)
           return
     }
		 fmt.Fprintf(c,"transaction: ", result)
     // fmt.Fprintf(c,"transaction: penđing")
 }
 // call transfer token
 func balance(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")

     bal, err := utils.BalaneOf(account)
     if err != nil {
         fmt.Fprintf(c,"error:",err)
         return
     }
     fmt.Fprintf(c,"balance:",bal)
 }
 // call transfer token
 func report(c *routing.Context){
     fmt.Println("Start report")
     report := utils.Report()

     fmt.Fprintf(c,"data:" + report)
 }
 func accounts(c *routing.Context){
     accounts, err := utils.GetAccountList()
     if err != nil {
       // handle error
			 fmt.Fprintf(c,"error:",err )
			 return
     }
    fmt.Fprintf(c,"accounts",accounts )
 }
 func getKey(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")
		 
		 val, err := utils.GetRedisAccountKey(account)
     if err != nil {
         fmt.Fprintf(c,"error:",err)
         return
     }
    fmt.Fprintf(c,"key:",val )
 }
