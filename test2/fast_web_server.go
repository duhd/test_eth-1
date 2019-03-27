package main

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
	"github.com/valyala/fasthttp"
  "os"
  "strings"
  "github.com/go-redis/redis"
  "encoding/json"
  "test_eth/test2/utils"
)


var cfg *utils.Config

var clients []*utils.EthClient
var current int = 0

func init() {
  config_file := "config.yaml"
  if len(os.Args) == 2 {
      config_file = os.Args[1]
   }

   println("init function")
   cfg = utils.LoadConfig(config_file)

   //Creat redis connection
   utils.Redis_client = redis.NewClient(&redis.Options{
     Addr:     cfg.Redis.Host,
     Password: cfg.Redis.Password, // no password set
     DB:       cfg.Redis.Db,  // use default DB
   })

   utils.DeleteData("transaction*")
   utils.DeleteData("nonce*")

   // sha = sha1.New()
   utils.LoadKeyStores(cfg.Keys.Keystore)


    //Load all wallets in hosts
    max_connection := cfg.Webserver.MaxRpcConnection
    for i:=0 ; i<max_connection; i++ {
         for _,host := range cfg.Networks {
              ethclient, err := utils.NewEthClient(host.Http)
              if err != nil {
                continue
              }
              clients = append(clients,ethclient)
            }
     }

     //Sync nonce of account
     utils.SyncNonce(clients[0].Client)

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

     //fmt.Println("Transfer: ", current," from ",from," to ",to, " amount: ",amount, " note:",append)
     client := clients[current]

     // go func() {
     //     result, err := client.TransferToken(from,to,amount,append)
     //     if err != nil {
     //         fmt.Println("Error to transfer token: ", err)
     //         return
     //     }
     //     fmt.Println("Transaction: ", result)
     //   }()


     result, err := client.TransferToken(from,to,amount,append)
     if err != nil {
           fmt.Println("Error to transfer token: ", err)
           fmt.Fprintf(c,"Error to transfer token: ", err)
           return
     }
     current = current + 1
     current = current % len(clients)
     fmt.Fprintf(c,"transaction: ", result)
     // fmt.Fprintf(c,"transaction: penđing")
 }
 // call transfer token
 func balance(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")

     client := clients[current]
     bal, err := client.BalaneOf(account)

     if err != nil {
         fmt.Fprintf(c,"error:",err)
         return
     }
     fmt.Fprintf(c,"balance:",bal)
 }
 // call transfer token
 func report(c *routing.Context){
     fmt.Println("Start report")
     keys, err  := utils.Redis_client.Keys("transaction:*").Result()
     if err != nil {
       // handle error
       fmt.Println(" Cannot get keys ")
     }
     vals, err1 := utils.Redis_client.MGet(keys...).Result()
     if err1 != nil {
       // handle error
       fmt.Println(" Cannot get values of  keys: ", keys)
     }

     fmt.Println("Elements: ", len(keys))
     diff_arr1 := []int64{}
     diff_arr := []int64{}

     for _, element := range vals {
         data := &utils.Transaction{}
         err2 := json.Unmarshal([]byte(element.(string)), data)
         if err2 != nil {
             fmt.Println("Element:", element, ", Error:", err2)
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

     fmt.Fprintf(c,"Total Tx",len2 , "Total Complete TX",len, "Avg RequestTime", avg1, "Avg Onchain",avg)
 }
 func accounts(c *routing.Context){
     keys, err  := utils.Redis_client.Keys("account*").Result()
     if err != nil {
       // handle error
       fmt.Println(" Cannot get keys ")
     }
     accounts := []string{}
     for _, element := range keys {
        account := strings.TrimPrefix(element,"account:")
        accounts = append(accounts,account)
     }
    fmt.Fprintf(c,"accounts",accounts )
 }
 func getKey(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")
     val, err := utils.Redis_client.Get("account:"+account).Result()
     if err != nil {
         fmt.Fprintf(c,"error:",err)
         return
     }
    fmt.Fprintf(c,"key:",val )
 }
