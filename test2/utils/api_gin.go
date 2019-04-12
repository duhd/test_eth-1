package utils
import (
        "strings"
       "github.com/gin-gonic/gin"
        "net/http"
        "fmt"
)

type ApiGin struct {
	 config * Config
	 walletHandler *WalletHandler
	 redisHandler *RedisHandler
}

func NewApiGin(cfg *Config, client *RpcRouting, rcache *RedisPool) *ApiGin{
			whandler := NewWalletHandler(cfg.Contract.Address, client)
			rhandler := NewRedisHandler(rcache)

      return &ApiGin{
        walletHandler:whandler,
				redisHandler:rhandler,
				config: cfg,
      }
}


// createTodo add a new todo
func (api *ApiGin) ProcessCall(c *gin.Context){
  method := c.Param("method")
  switch method {
      case "transfer":
           fmt.Println("call transfer")
           api.transfer(c)
           return
       case "balance":
           fmt.Println("call balance")
           api.balance(c)
           return
       case "report":
           fmt.Println("call report")
           api.report(c)
           return
       case "accounts":
           fmt.Println("call accounts")
           api.accounts(c)
           return
       case "key":
           fmt.Println("call key")
           api.getKey(c)
           return
       case "test":
             fmt.Println("call test")
             c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "test"})
             return
   }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "not find"})
}

// call transfer token
func (api *ApiGin) transfer(c *gin.Context){

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


    result, err := api.walletHandler.TransferToken(from,to,amount,append)
    if err != nil {
          fmt.Println("Error to transfer token: ", err)
          c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
          return
    }

    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "transaction": result})
    //c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "transaction": "pending"})
}

// call transfer token
func (api *ApiGin)  balance(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")

    bal, err := api.walletHandler.BalaneOf(account)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error ": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "balance": bal})
}

// call transfer token
func (api *ApiGin) report(c *gin.Context){
    fmt.Println("Start report")
    report := api.redisHandler.Report()

    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": report})
}
func (api *ApiGin) accounts(c *gin.Context){
    accounts, err := api.walletHandler.GetAccountList()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "accounts": accounts})
}
func (api *ApiGin) getKey(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")

    val, err := api.walletHandler.GetAccountKey(account)
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "key": val})
}
