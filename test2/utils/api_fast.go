package utils

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
  "strings"
  // "github.com/go-redis/redis"
  // "encoding/json"
)

type ApiFast struct {
	 config * Config
	 walletHandler *WalletHandler
	 redisHandler *RedisHandler
}

func NewApiFast(cfg *Config, client *RpcRouting, rcache *RedisPool) *ApiFast{
			whandler := NewWalletHandler(cfg.Contract.Address, client)
			rhandler := NewRedisHandler(rcache)
      return &ApiFast{
					config: cfg,
	        walletHandler:whandler,
					redisHandler:rhandler,
      }
}

// createTodo add a new todo
func (api *ApiFast) ProcessCall(c *routing.Context) error {
  method := c.Param("method")
  switch method {

      case "transfer":
           fmt.Println("call transfer")
           api.transfer(c)
           return  nil
       case "balance":
           fmt.Println("call balance")
           api.balance(c)
           return nil
       case "report":
           fmt.Println("call report")
           api.report(c)
           return nil
			case "new_account":
           fmt.Println("call new_account")
           api.new_account(c)
           return nil
       case "accounts":
           fmt.Println("call accounts")
           api.accounts(c)
           return nil
       case "key":
           fmt.Println("call key")
           api.getKey(c)
           return nil
			 case "test":
           fmt.Println("call test")
           fmt.Fprintf(c, "data=test")
           return nil
			 case "eth_transfer":
            fmt.Println("call eth_transfer")
            api.eth_transfer(c)
            return  nil
      case "eth_balance":
            fmt.Println("call eth_balance")
            api.eth_balance(c)
            return nil
   }

   fmt.Fprintf(c, "URL not found ")
   return nil
 }


 // call transfer token
 func (api *ApiFast) transfer(c *routing.Context){
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

  	 result, err := api.walletHandler.TransferToken(from,to,amount,append)
     if err != nil {
           fmt.Fprintf(c,"Error to transfer token: %v", err)
           return
     }
		 fmt.Fprintf(c,"transaction: %v ", result)
     // fmt.Fprintf(c,"transaction: penđing")
 }
 // call transfer token
 func (api *ApiFast)  balance(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")

     bal, err := api.walletHandler.BalaneOf(account)
     if err != nil {
         fmt.Fprintf(c,"error: %v",err)
         return
     }
     fmt.Fprintf(c,"balance: %d",bal)
 }
 // call transfer token
 func (api *ApiFast) report(c *routing.Context){
     fmt.Println("Start report")
     report := api.redisHandler.Report()

     fmt.Fprintf(c,"data:" + report)
 }
 func (api *ApiFast) new_account(c *routing.Context){
     account, err := api.walletHandler.NewTokenAccount()
     if err != nil {
       // handle error
			 fmt.Fprintf(c,"error: %v",err )
			 return
     }
    fmt.Fprintf(c,"account: %v",account )
 }

 func (api *ApiFast) accounts(c *routing.Context){
     accounts, err := api.walletHandler.GetAccountList()
     if err != nil {
       // handle error
			 fmt.Fprintf(c,"error: %v",err )
			 return
     }
    fmt.Fprintf(c,"accounts: %v",accounts )
 }

 func (api *ApiFast) getKey(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")

		 val, err := api.walletHandler.GetAccountKey(account)
     if err != nil {
         fmt.Fprintf(c,"error: %v",err)
         return
     }
    fmt.Fprintf(c,"key: %v",val )
 }

  // call transfer eth
  func (api *ApiFast)  eth_transfer(c *routing.Context){
      from := c.Param("p1")
      to := c.Param("p2")
      amount := c.Param("p3")

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

   	 result, err := api.walletHandler.EthTransfer(from,to,amount)
      if err != nil {
            fmt.Fprintf(c,"Error to transfer token: %v", err)
            return
      }
 		 fmt.Fprintf(c,"transaction: %v ", result)
      // fmt.Fprintf(c,"transaction: penđing")
  }

  // call transfer token
  func (api *ApiFast) eth_balance(c *routing.Context){
      account := c.Param("p1")
      account = strings.TrimPrefix(account,"0x")

      bal, err := api.walletHandler.EthBalaneOf(account)
      if err != nil {
          fmt.Fprintf(c,"error: %v",err)
          return
      }
      fmt.Fprintf(c,"balance: %d",bal)
  }
