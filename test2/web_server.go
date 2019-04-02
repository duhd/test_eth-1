package main

import (
        // "strconv"
        // "math/rand"
        // "crypto/sha1"
        // "encoding/base64"
        // "hash"
        "os"
        "strings"
       "github.com/gin-gonic/gin"
       // "github.com/go-redis/redis"
        "net/http"
        // "encoding/json"
        "fmt"
        "test_eth/test2/utils"
        "sync"

)

var cfg *utils.Config

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
	redisPool := utils.NewRedisPool()

	println("Delete old data in redis ")
	//utils.DeleteData("transaction*")
	//utils.DeleteData("nonce*")

	// sha = sha1.New()
	println("Load key in account array ")
	utils.LoadKeyStores(cfg.Keys.Keystore)

	 //Load all wallets in hosts
	 println("Create rpc connection pool ")
	 clientPool := utils.NewClientPool()

	 //Sync nonce of account
	 println("sync nonce of account from ethereum ")
	 utils.SyncNonce(clientPool.GetClient().Client)


	 var wg sync.WaitGroup
	 wg.Add(3)

	 go func (){
		   println("Loop processs sending message ")
		   defer wg.Done()
			 clientPool.Loop()
	 }()

	 go func (){
		 	  println("Loop webservice ")
			  defer wg.Done()
				redisPool.Loop()
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
  v1 := router.Group("/api/v1")
  {
      v1.GET("/wallet/:method/:p1/:p2/:p3/:p4", processCall)
      v1.GET("/wallet/:method/:p1", processCall)
      v1.GET("/wallet/:method", processCall)
   }
   router.Run(":"+ cfg.Webserver.Port)
}
// createTodo add a new todo
func processCall(c *gin.Context){
  method := c.Param("method")
  switch method {
      case "transfer":
           fmt.Println("call transfer")
           transfer(c)
           return
       case "balance":
           fmt.Println("call balance")
           balance(c)
           return
       case "report":
           fmt.Println("call report")
           report(c)
           return
       case "accounts":
           fmt.Println("call accounts")
           accounts(c)
           return
       case "key":
           fmt.Println("call key")
           getKey(c)
           return
       case "test":
             fmt.Println("call test")
             c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "test"})
             return
   }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "not find"})
}

// call transfer token
func transfer(c *gin.Context){

    from := c.Param("p1")
    to := c.Param("p2")
    amount := c.Param("p3")
    append := c.Param("p4")

    if from == "" {
      c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": "Please add from address "})
      return
    }
    if to == "" {
      c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": "Please add to address "})
      return
    }
    from = strings.TrimPrefix(from,"0x")
    to = strings.TrimPrefix(to,"0x")

    //fmt.Println("Transfer: ", current," from ",from," to ",to, " amount: ",amount, " note:",append)

    // go func() {
    //     result, err := client.TransferToken(from,to,amount,append)
    //     if err != nil {
    //         fmt.Println("Error to transfer token: ", err)
    //         return
    //     }
    //     fmt.Println("Transaction: ", result)
    //   }()


    result, err := utils.TransferToken(from,to,amount,append)
    if err != nil {
          fmt.Println("Error to transfer token: ", err)
          c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
          return
    }

    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "transaction": result})
    //c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "transaction": "pending"})
}

// call transfer token
func balance(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")

    bal, err := utils.BalaneOf(account)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error ": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "balance": bal})
}

// call transfer token
func report(c *gin.Context){
    fmt.Println("Start report")
    report := utils.Report()

    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": report})
}
func accounts(c *gin.Context){
    accounts, err := utils.GetAccountList()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "accounts": accounts})
}
func getKey(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")

    val, err := utils.GetRedisAccountKey(account)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "key": val})
}
