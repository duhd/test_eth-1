package utils

import (
  "os"
  "fmt"
  "io/ioutil"
  "strings"
  "path/filepath"
    "context"
  // "encoding/json"
  // "time"
  "strconv"
  "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/accounts/keystore"
  "crypto/ecdsa"
   "sync/atomic"
  "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
)

type WalletAccount struct {
    Address string
    Nonce uint64
    PrivateKey *ecdsa.PrivateKey
}


func (w *WalletAccount) GetNonce() uint64 {
    nonce := atomic.AddUint64(&w.Nonce, 1)
    fmt.Println("Get Nonce:",nonce)
    return nonce
}

func (w *WalletAccount) UpdateNonce(nonce uint64)  {
    fmt.Println("Update: ",w.Address," Nonce:",nonce)
    atomic.StoreUint64(&w.Nonce, nonce-1)
}


func SyncNonce(backend   *ethclient.Client){
  for _,wallet := range Wallets {
    keyAddr := common.HexToAddress(wallet.Address)
    nonce, err := backend.PendingNonceAt(context.Background(), keyAddr)
    if err != nil {
      fmt.Errorf("failed to retrieve account nonce: %v", err)
      nonce = 0
    }
    fmt.Println("Nonce from eth: ",nonce)
    wallet.UpdateNonce(nonce)
  }
}

func GetWallet(addr string) *WalletAccount {
    for _, wallet := range Wallets {
       if wallet.Address == addr {
         return wallet
       }
    }
    return nil
}
// var cfg *Config
var Wallets []*WalletAccount

func DeleteData(pattern string){
  client := Rclients.getClient()
  keys, err  := client.Keys(pattern).Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get keys ")
  }
  if len(keys) >0 {
    res := client.Del(keys...)
    fmt.Println("Redis delete: ", res)
  }
}

func LoadKeyStores(root string){
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
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }


              //Store account private key
              accountKey, err := keystore.DecryptKey( []byte(keyjson), cfg.Keys.Password)
              if err != nil {
                  fmt.Println("Cannot decrypt key file: ", err)
                  return
              }
              privateKey := accountKey.PrivateKey


              //Add to array list
              wallet := WalletAccount{
                  Address: list[2],
                  Nonce: uint64(0),
                  PrivateKey: privateKey,
              }
              Wallets = append(Wallets,&wallet)

              client := Rclients.getClient()
               //Store full account key
              account := "account:" + list[2]
              //Set key in redis
             err = client.Set(account,string(keyjson), 0).Err()
             if err != nil {
               panic(err)
             }

            private := "private:" + list[2]
            //Set key in redis
            err = client.Set(private,string(crypto.FromECDSA(privateKey)), 0).Err()
            if err != nil {
               panic(err)
            }
         }
    }
}
func GetAccountList() ([]string, error){
  client := Rclients.getClient()
  keys, err  := client.Keys("account*").Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get keys ")
    return nil, err
  }
  accounts := []string{}
  for _, element := range keys {
     account := strings.TrimPrefix(element,"account:")
     accounts = append(accounts,account)
  }
  return accounts, nil
}
func GetRedisAccountKey(account string) (string, error) {
    client := Rclients.getClient()
    key, err := client.Get("private:"+account).Result()
    if err != nil {
      // handle error
      fmt.Println(" Cannot get keys ")
      return "", err
    }
    return key, err
}

func getNonce(backend  *ethclient.Client, account string) uint64 {
     nonce := GetNonce(account)
     if nonce == 0 {
        nonce, _ = UpdateNonceFromEth(backend, account)
        CommitNonce(account,nonce)
     }
     NoneIncr(account)
     return nonce
}

func UpdateNonceFromEth(backend  *ethclient.Client, account string) (uint64,error) {
      client := Rclients.getClient()
      keyjson, err := client.Get("account:"+account).Result()
      if err != nil {
          return 0, err
      }

      opts, err := bind.NewTransactor(strings.NewReader(keyjson),cfg.Keys.Password)
      if err != nil {
            fmt.Println("Failed to create authorized transactor: %v", err)
            return 0, err
      }
      var nonce uint64
      if opts.Nonce == nil {
        nonce, err = backend.PendingNonceAt(context.Background(), opts.From)
        if err != nil {
          return 0, fmt.Errorf("failed to retrieve account nonce: %v", err)
        }
      } else {
        nonce = opts.Nonce.Uint64()
      }
      if CommitNonce(account,nonce) {
        fmt.Println("Failed to create authorized transactor: %v", err)
      }
      return nonce,nil
}

func  GetNonce(account string) uint64 {
  client := Rclients.getClient()
  val, err := client.Get("nonce:" + account).Result()
  if err != nil {
      fmt.Println("Cannot find nonce of account: ", account)
      return uint64(0)
  }
  value , err := strconv.ParseUint(val, 10, 64)
  if err != nil {
      fmt.Println("Cannot parce nonce of ", val)
      return uint64(0)
  }
  return value
}
func CommitNonce(account string, nonce uint64) bool {
  client := Rclients.getClient()
  err := client.Set("nonce:" + account,uint64(nonce), 0).Err()
  if err != nil {
       fmt.Println("Cannot set nonce  ", err)
       return false
  }
  return true
}
func  NoneIncr(account string) bool {
  client := Rclients.getClient()
  _, err := client.Incr("nonce:" + account).Result()
	if err != nil {
    fmt.Println("Cannot increase nonce  ", err)
    return false
	}
  return true
}
