package utils

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
  "strings"
	"math/big"
	// "github.com/ethereum/go-ethereum/common"
  // "github.com/go-redis/redis"
  // "encoding/json"
	  "strconv"
)

type ApiFastV2 struct {
	 config * Config
	 walletHandler *F5WalletHandler
	 redisHandler *RedisHandler
}

func NewApiFastV2(cfg *Config, client *RpcRouting, rcache *RedisPool) *ApiFastV2{
			whandler := NewF5WalletHandler(cfg.F5Contract.Address, client)
			rhandler := NewRedisHandler(rcache)
      return &ApiFastV2{
					config: cfg,
	        walletHandler:whandler,
					redisHandler:rhandler,
      }
}

// createTodo add a new todo
func (api *ApiFastV2) ProcessCall(c *routing.Context) error {
  method := c.Param("method")
  switch method {
			case "create":
           fmt.Println("call create")
           api.create(c)
           return  nil
       case "balance":
           fmt.Println("call balance")
           api.balance(c)
           return nil
			 case "state":
           fmt.Println("call state")
           api.state(c)
           return nil
       case "set_state":
           fmt.Println("call set_state")
           api.set_state(c)
           return nil
			 case "withdraw":
           fmt.Println("call withdraw")
           api.withdraw(c)
           return nil
			 case "deposit":
           fmt.Println("call deposit")
           api.deposit(c)
           return nil
			 case "transfer":
 					fmt.Println("call transfer")
 					api.transfer(c)
 					return nil
			case "new_account":
           fmt.Println("call new_account")
           api.new_account(c)
           return nil

			 case "register":
		 			 fmt.Println("call register")
		 			 api.registerAccounts(c)
		 			 return nil
       case "accounts":
           fmt.Println("call accounts")
           api.accounts(c)
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
func (api *ApiFastV2)  registerAccounts(c *routing.Context){
		 api.walletHandler.LoadAccountEth()
		 api.walletHandler.AutoFillGas()
		 list := api.walletHandler.RegisterBatchEthToContract()
		 list_string := strings.Join(list,",")
		 fmt.Fprintf(c,list_string)
}
 // call create wallet
 func (api *ApiFastV2)  create(c *routing.Context){
     account := c.Param("p1")
     account = strings.TrimPrefix(account,"0x")
		 typeStash := c.Param("p2")

		 typewallet, err :=  strconv.Atoi(typeStash)
		 if err != nil {
         fmt.Fprintf(c,"error: %v",err)
         return
     }

     tx, err := api.walletHandler.CreateStash(account,int8(typewallet))
     if err != nil {
         fmt.Fprintf(c,"error: %v",err)
         return
     }
     fmt.Fprintf(c,"transaction hash: ",tx.Hash().Hex())
 }

 // call balance of wallet
 func (api *ApiFastV2)  balance(c *routing.Context){
 		account := c.Param("p1")
 		account = strings.TrimPrefix(account,"0x")

 		bal, err := api.walletHandler.GetBalance(account)
 		if err != nil {
				fmt.Println("Error in call GetBalance:", err)
 				fmt.Fprintf(c,"error: ")
 				return
 		}
		fmt.Println("Return value:", bal)
 		fmt.Fprintf(c,"balance: %d",bal)
 }
 // call get wallet state
 func (api *ApiFastV2) state(c *routing.Context){
		account := c.Param("p1")
		account = strings.TrimPrefix(account,"0x")

		state, err := api.walletHandler.GetState(account)
		if err != nil {
			  fmt.Println("Error in state: ",err)
				fmt.Fprintf(c,"error:")
				return
		}
		fmt.Fprintf(c,"transaction hash: ",state)
 }
 // call set wallet state
 func (api *ApiFastV2) set_state(c *routing.Context){
		 account := c.Param("p1")
		 account = strings.TrimPrefix(account,"0x")
		 state := c.Param("p2")

		 stashState, err := strconv.Atoi(state)
		 if err != nil {
			fmt.Fprintf(c,"error: Please txType as integer ")
			return
		}
		 tx, err := api.walletHandler.SetState(account,int8(stashState))
		 if err != nil {
				 fmt.Fprintf(c,"error: %v",err)
				 return
		 }
		 fmt.Fprintf(c,"transaction hash: ",tx.Hash().Hex())
 }

 // call get wallet state
 func (api *ApiFastV2) withdraw(c *routing.Context){
		txRef := c.Param("p1")
		account := c.Param("p2")
		account = strings.TrimPrefix(account,"0x")
		value := c.Param("p3")

		amount := new(big.Int)
		amount.SetString(value,10)

		tx, err := api.walletHandler.Withdraw(txRef,account,amount)
		if err != nil {
				fmt.Fprintf(c,"error: %v",err)
				return
		}
		fmt.Fprintf(c,"transaction hash: ",tx.Hash().Hex())
 }
 // call get wallet state
 func (api *ApiFastV2) deposit(c *routing.Context){
	 txRef := c.Param("p1")
	 account := c.Param("p2")
	 account = strings.TrimPrefix(account,"0x")
	 value := c.Param("p3")

	 amount := new(big.Int)
	 amount.SetString(value,10)

	 tx, err := api.walletHandler.Deposit(txRef,account,amount)
	 if err != nil {
			 fmt.Fprintf(c,"error: %v",err)
			 return
	 }
	 fmt.Fprintf(c,"transaction hash: ",tx.Hash().Hex())
 }
 // call transfer token
 func (api *ApiFastV2) transfer(c *routing.Context){
	   txRef := c.Param("p1")
     sender := c.Param("p2")
     receiver := c.Param("p3")
     value := c.Param("p4")
     note := c.Param("p5")
		 txtyp := c.Param("p6")

     if sender == "" {
       fmt.Fprintf(c,"error: Please add sender address ")
       return
     }
     if receiver == "" {
       fmt.Fprintf(c,"error: Please add receiver address ")
       return
     }
		 amount := new(big.Int)
		 amount.SetString(value,10)

		 txType, err :=  strconv.Atoi(txtyp)
		 if err != nil {
			 fmt.Fprintf(c,"error: Please txType as integer ")
			 return
		 }

  	 result, err := api.walletHandler.Transfer(txRef,sender,receiver,amount,note,int8(txType))
     if err != nil {
           fmt.Fprintf(c,"Error to transfer token: %v", err)
           return
     }
		 fmt.Fprintf(c,"transaction: %v ", result.Hash().Hex())
     // fmt.Fprintf(c,"transaction: penđing")
 }
 // call transfer token
 func (api *ApiFastV2) report(c *routing.Context){
     fmt.Println("Start report")
     report := api.redisHandler.Report()

     fmt.Fprintf(c,"data:" + report)
 }
 func (api *ApiFastV2) new_account(c *routing.Context){
     account, err := api.walletHandler.NewAccountEth()
     if err != nil {
       // handle error
			 fmt.Fprintf(c,"error: %v",err )
			 return
     }
    fmt.Fprintf(c,"account: %v",account )
 }

 func (api *ApiFastV2) accounts(c *routing.Context){
     accounts := api.walletHandler.GetAccountList()
  	 list := []string{}
		 for _,account := range accounts {
			   addr := account.Hex()
				 addr = strings.ToLower(strings.TrimPrefix(addr,"0x"))
				 list = append(list,addr)
		 }
    fmt.Fprintf(c,"accounts: %v",list )
 }
  // call transfer eth
  func (api *ApiFastV2)  eth_transfer(c *routing.Context){
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
  func (api *ApiFastV2) eth_balance(c *routing.Context){
      account := c.Param("p1")
      account = strings.TrimPrefix(account,"0x")

      bal, err := api.walletHandler.EthBalaneOf(account)
      if err != nil {
          fmt.Fprintf(c,"error: %v",err)
          return
      }
      fmt.Fprintf(c,"balance: %d",bal)
  }
