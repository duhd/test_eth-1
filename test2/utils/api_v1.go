package utils

import (
	"fmt"
	"github.com/qiangxue/fasthttp-routing"
		"github.com/valyala/fasthttp"
  "strings"
  // "github.com/go-redis/redis"
  "encoding/json"
)

type ApiResponse struct  {
	  Rescode int  `json:"rescode"`
		Resdecr string  `json:"resdecr"`
		Resdata interface{}  `json:"resdata"`
}

func (res *ApiResponse) toJson() string {
	  b, err := json.Marshal(res)
    if err != nil {
        //fmt.Println("Json parse error: : ",err)
        return ""
    }
		//fmt.Println("Json parse ok: : ",string(b))
    return string(b)
}

func Response(c  *routing.Context, code int, desr string, data interface{} ) error {
		res := &ApiResponse{
				Rescode: code,
				Resdecr: desr,
				Resdata: data,
		}
		fmt.Fprintf(c,res.toJson())
		return nil
}


////////////// Login function ///////////
func Api_login_get(c  *routing.Context) error {
		fmt.Println("GET: API Access Login")
		//fmt.Fprintf(c,"Please POST: username and password ")
		return Response(c,99,"Please use POST method",nil)
}

// JSON Authentication
func Api_login(c  *routing.Context) error {
	fmt.Println("API Access Login. ")

	username := string(c.FormValue("username"))
	password := string(c.FormValue("password"))

	fmt.Println("Username:  " + username  + ", " + password)

	if cfg.JwtVerifyUsername(username,password) {
			fmt.Println("Start JwtVerifyUsername")
			qUser := []byte(username)
			qPasswd := []byte(password)

			fasthttpJwtCookie := c.Request.Header.Cookie("fasthttp_jwt")

			// for example, server receive token string in request header.
			if len(fasthttpJwtCookie) == 0 {
				fmt.Println("Start creating token ")

				tokenString, expireAt := CreateToken(qUser, qPasswd)
				//fmt.Println("Get cookied ")
				// Set cookie for domain
				cookie := fasthttp.AcquireCookie()
				//fmt.Println("End get cookie ")

				cookie.SetKey("fasthttp_jwt")

				cookie.SetValue(tokenString)
				cookie.SetExpire(expireAt)
				c.Response.Header.SetCookie(cookie)

				fmt.Println("End creating token ")
				return Response(c,0,"New Jwt token successfully",nil)
			}
			return Response(c,98,"Existed JWT token, no need recreate",nil)
	}
	return Response(c,99,"Username or password failed",nil)
}


/***************** Cash credit function  ************/
 func Api_cash_credit(c *routing.Context) error {
     address := c.Param("address")
     amount := c.Param("amount")
     traceid := c.Param("traceid")

     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if amount == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }
     if traceid == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     address = strings.TrimPrefix(address,"0x")
     //
     //
  	 // result, err := utils.TransferToken(from,to,amount,append)
     // if err != nil {
     //       fmt.Fprintf(c,"Error to transfer token: ", err)
     //       return
     // }
		 // fmt.Fprintf(c,"transaction: ", result)
     // fmt.Fprintf(c,"transaction: penđing")
		 return nil
 }

 // call transfer token
 func Api_cash_debit(c *routing.Context) error {
     address := c.Param("address")
     amount := c.Param("amount")
     traceid := c.Param("traceid")

     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if amount == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }
     if traceid == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     address = strings.TrimPrefix(address,"0x")
     // fmt.Fprintf(c,"transaction: penđing")
		  return nil
 }

 // call transfer token
 func Api_cash_transfer(c *routing.Context) error {
     from := c.Param("from")
     to := c.Param("to")
     amount := c.Param("amount")
     note := c.Param("note")
     traceid := c.Param("traceid")

     if from == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     if to == "" {
       fmt.Fprintf(c,"error: Please add to address ")
       return nil
     }

     from = strings.TrimPrefix(from,"0x")
     to = strings.TrimPrefix(to,"0x")
     //
  	 // result, err := utils.TransferToken(from,to,amount,append)
     // if err != nil {
     //       fmt.Fprintf(c,"Error to transfer token: ", err)
     //       return
     // }
		 // fmt.Fprintf(c,"transaction: ", result)
     fmt.Fprintf(c,"Api_cash_transfer: %v,%v,%v,%v,%v ",from,to,amount,note,traceid)
		 return nil
 }



///// Balance function ///////////
 func Api_balance(c *routing.Context) error {
   address := c.Param("address")

   if address == "" {
     fmt.Fprintf(c,"error: Please add from address ")
     return nil
   }
   address = strings.TrimPrefix(address,"0x")

   if address == "" {
     fmt.Fprintf(c,"error: Please add from address ")
     return nil
   }

    fmt.Fprintf(c,"Api_balance: %v ", address)
		return nil
 }
 // call transfer token
 func Api_balance_all(c *routing.Context) error {
		 fmt.Fprintf(c,"Api_balance_all ")
     // fmt.Fprintf(c,"transaction: penđing")
		 return nil
 }

 //////// Accunt functions //////////
 func Api_account_new(c *routing.Context) error {
   fmt.Println("Api_account_new: start")
	 res := &ApiResponse{
			 Rescode: 0,
			 Resdecr: "successfully create new account",
			 Resdata: nil,
	 }
	 fmt.Fprintf(c,res.toJson())
	 return nil
 }
 func Api_account_total(c *routing.Context) error {
   fmt.Fprintf(c,"Api_account_total: ")
	 return nil
 }

 func Api_account_list_active(c *routing.Context) error  {
    fmt.Fprintf(c,"Api_account_list_active: ")
		return nil
 }

 func Api_account_list_inactive(c *routing.Context) error  {
   fmt.Fprintf(c,"Api_account_list_inactive: ")
	 return nil
 }

 func Api_account_lock(c *routing.Context) error {
    fmt.Fprintf(c,"Api_account_list_inactive: ")
    address := c.Param("address")
    traceid := c.Param("traceid")

    if address == "" {
      fmt.Fprintf(c,"error: Please add from address ")
      return nil
    }
    address = strings.TrimPrefix(address,"0x")
    fmt.Fprintf(c,"Api_account_lock: %v %s %v",address,", traceid: ",traceid)
		return nil
 }
 func Api_account_status(c *routing.Context) error {
     fmt.Fprintf(c,"Api_account_status: ")
     address := c.Param("address")
     if address == "" {
       fmt.Fprintf(c,"error: Please add from address ")
       return nil
     }
     fmt.Fprintf(c,"Api_account_status: ")
		 return nil
 }


 ///// Transaction functions
 func Api_transaction(c *routing.Context) error {
   txhash := c.Param("txhash")

   if txhash == "" {
     fmt.Fprintf(c,"error: Please add txhash ")
     return nil
   }
   txhash = strings.TrimPrefix(txhash,"0x")
   fmt.Fprintf(c,"Api_transaction: %v",txhash)
	 return nil
 }

 func Api_transaction_list(c *routing.Context) error  {
    account := c.Param("account")
    fromdate := c.Param("fromdate")
    todate := c.Param("todate")

   if account == "" {
     fmt.Fprintf(c,"error: Please add account ")
     return nil
   }
   account = strings.TrimPrefix(account,"0x")
   fmt.Fprintf(c,"Api_transaction_list: %v,%v,%v",account,fromdate,todate)
return nil
 }
 func Api_transaction_lock(c *routing.Context) error {
   account := c.Param("account")
   fromdate := c.Param("fromdate")
   todate := c.Param("todate")

  if account == "" {
    fmt.Fprintf(c,"error: Please add account ")
    return nil
  }
  account = strings.TrimPrefix(account,"0x")
  fmt.Fprintf(c,"Api_transaction_lock: %v,%v,%v",account,fromdate,todate)
	return nil
 }
