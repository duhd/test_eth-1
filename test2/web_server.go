package main

import (
        "os"
        "strings"
       "github.com/gin-gonic/gin"
       "github.com/go-redis/redis"
        "net/http"
        "encoding/json"
        // "strconv"
        // "math/rand"
        "fmt"
  			"time"
        // "crypto/sha1"
        // "encoding/base64"
        // "hash"
        "path/filepath"
      	"io/ioutil"
        "gopkg.in/yaml.v2"
        "test_eth/contracts"
      	"math/big"
        "github.com/ethereum/go-ethereum/ethclient"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/common"
)
type Transaction struct {
        Id                string  `json:"Id"`
        RequestTime       int64   `json:"RequestTime"`
        TxReceiveTime     int64   `json:"TxReceiveTime"`
        TxConfirmedTime    []int64 `json:"TxConfiredTime"`
   }


type Config struct {
			Keys  struct {
				  Keystore string `yaml:"keystore"`
					Password string `yaml:"password"`
			} `yaml:"keys"`
			Networks []struct {
					Name string `yaml:"name"`
					Http string `yaml:"http"`
					WebSocket string `yaml:"websocket"`
					LocalAddr string `yaml:"local"`
			} `yaml:"networks"`
			Redis struct {
				  Host string `yaml:"host"`
				  Password string `yaml:"password"`
				  Db int `yaml:"db"`
			} `yaml:"redis"`
			Contract struct {
					Owner string `yaml:"owner"`
					InitialToken int64 `yaml:"initialToken"`
					MasterKey1 string `yaml:"masterkey1"`
					MasterKey2 string `yaml:"masterkey2"`
					Address string `yaml:"address"`
			} `yaml:"contract"`
	}

var cfg *Config
var redis_client *redis.Client
// var sha hash.Hash
var clients []*ethclient.Client
var current int = 0

func init() {
   //open a db connection
   println("init function")
   cfg = loadConfig("config.yaml")

   //Creat redis connection
   redis_client = redis.NewClient(&redis.Options{
     Addr:     cfg.Redis.Host,
     Password: cfg.Redis.Password, // no password set
     DB:       cfg.Redis.Db,  // use default DB
   })

   // sha = sha1.New()
   loadKeyStores(cfg.Keys.Keystore)


    //Load all wallets in hosts
   for _,host := range cfg.Networks {
         fmt.Println("Connect to host: ", host.Http)
         client, err  := ethclient.Dial("http://" + host.Http)
         if err != nil {
            fmt.Println("Unable to connect to network:%v\n", err)
         }
        clients = append(clients,client)
   }
}
func loadConfig(file string) *Config {
     cfg := &Config{}

     yamlFile, err := ioutil.ReadFile(file)
     if err != nil {
         fmt.Println("yamlFile.Get err   #%v ", err)
     }

     err = yaml.Unmarshal(yamlFile, cfg)
     if err != nil {
         fmt.Println("Unmarshal: %v", err)
     }
     return cfg
}

func loadKeyStores(root string){
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
               files = append(files, path)
               return nil
           })
    if err != nil {
         panic(err)
    }
    for _, file := range files {
         fmt.Println("File:", file)
         list := strings.Split(file,"--")
         if len(list) == 3 {
             account := "account:" + list[2]
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }
             //Set key in redis
              err = redis_client.Set(account,string(keyjson), 0).Err()
              if err != nil {
                panic(err)
              }
         }
    }
}


func logStart(key string,requesttime int64){
  trans :=  &Transaction{
              Id: key,
              RequestTime: requesttime,
              TxReceiveTime: time.Now().UnixNano()}
  value, err := json.Marshal(trans)
  if err != nil {
      fmt.Println(err)
      return
  }
  err = redis_client.Set("transaction:" + key,string(value), 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
  router := gin.Default()
  // Simple group: v1
  v1 := router.Group("/api/v1")
  {
      v1.GET("/wallet/:method/:p1/:p2/:p3/:p4", processCall)
      v1.GET("/wallet/:method/:p1", processCall)
      v1.GET("/wallet/:method", processCall)
   }
   router.Run()
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
   }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "not find"})
}

// call transfer token
func transfer(c *gin.Context){
    requestTime := time.Now().UnixNano()

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

    fmt.Println("Transfer: ", current," from ",from," to ",to, " amount: ",amount, " note:",append)
    keyjson, err := redis_client.Get("account:"+from).Result()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }

    auth, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
  	if err != nil {
  		fmt.Println("Failed to create authorized transactor: %v", err)
  	}

    address := common.HexToAddress(to)
  	value := new(big.Int)
  	value, ok := value.SetString(amount, 10)
  	 if !ok {
  			 fmt.Println("SetString: error")
         c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": "Please add amount "})
  			 return
  	 }

  	note :=  fmt.Sprintf("Transaction:  %s", append)

    client := clients[current]
    fmt.Println("Add contract: ", cfg.Contract.Address)
    wallet, err1 := contracts.NewVNDWallet(common.HexToAddress(cfg.Contract.Address), client)
    if err1 != nil {
       fmt.Println("Unable to bind to deployed instance of contract:%v\n")
   }

  	tx, err := wallet.Transfer(auth, address, value, []byte(note))
  	if err != nil {
  			fmt.Println(" Transaction create error: ", err)
  	}
  	fmt.Println(" Transaction =",tx.Hash().Hex())


    // seed := rand.Intn(100)
    // sha.Write([]byte(strconv.Itoa(seed)))
    // key := "Transfer:" + base64.URLEncoding.EncodeToString(sha.Sum(nil))
    key := strings.TrimPrefix(tx.Hash().Hex(),"0x")
    logStart(key,requestTime)


    current = current + 1
    current = current % len(clients)
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "transaction hash": tx.Hash().Hex()})
}
// call transfer token
func balance(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")


    address := common.HexToAddress("0x" + account)

    client := clients[current]
    fmt.Println("Add contract: ", cfg.Contract.Address)
    wallet, err1 := contracts.NewVNDWallet(common.HexToAddress(cfg.Contract.Address), client)
    if err1 != nil {
       fmt.Println("Unable to bind to deployed instance of contract:%v\n")
   }

  	bal, err := wallet.BalanceOf(&bind.CallOpts{}, address)

  	if err != nil {
  		fmt.Println("Get balanceof: ", err)
  	}

  	fbal := new(big.Float)

  	fbal.SetString(bal.String())
  	fmt.Printf("balance: %f", bal) // "balance: 74605500.647409"
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "balance": bal})
}
// call transfer token
func report(c *gin.Context){
    keys, err  := redis_client.Keys("transaction:*").Result()
    if err != nil {
      // handle error
      fmt.Println(" Cannot get keys ")
    }
    vals, err1 := redis_client.MGet(keys...).Result()
    if err1 != nil {
      // handle error
      fmt.Println(" Cannot get values of  keys: ", keys)
    }

    diff_arr := []int64{}
    for _, element := range vals {
        data := &Transaction{}
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
        }else {
            max = time.Now().UnixNano()
        }

        diff := max  - data.TxReceiveTime
        diff_arr = append(diff_arr,diff)
    }
    var total int64 = 0
  	for _, value:= range diff_arr {
  		total += value
  	}
    len := int64(len(diff_arr))
  	avg := total/len
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Total": len, "Avg": avg})
}

func accounts(c *gin.Context){
    keys, err  := redis_client.Keys("account*").Result()
    if err != nil {
      // handle error
      fmt.Println(" Cannot get keys ")
    }
    accounts := []string{}
    for _, element := range keys {
       account := strings.TrimPrefix(element,"account:")
       accounts = append(accounts,account)
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "accounts": accounts})
}
func getKey(c *gin.Context){
    account := c.Param("p1")
    account = strings.TrimPrefix(account,"0x")
    val, err := redis_client.Get("account:"+account).Result()
    if err != nil {
        c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "error": err})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "key": val})
}
