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
  "github.com/go-redis/redis"
  // "strconv"
  // "github.com/ethereum/go-ethereum/crypto"
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

var Redis_client *redis.Client
var Wallets []*WalletAccount

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

func DeleteData(pattern string){
  keys, err  := Redis_client.Keys(pattern).Result()
  if err != nil {
    // handle error
    fmt.Println(" Cannot get keys ")
  }
  if len(keys) >0 {
    res := Redis_client.Del(keys...)
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
              //Store full account key
             //account := "account:" + list[2]
             keyjson, err := ioutil.ReadFile(file)
             if err != nil {
                  fmt.Println("Error in read file: ", file )
                  continue
             }
             // //Set key in redis
             //  err = Redis_client.Set(account,string(keyjson), 0).Err()
             //  if err != nil {
             //    panic(err)
             //  }
             //
              //Store account private key
              accountKey, err := keystore.DecryptKey( []byte(keyjson), cfg.Keys.Password)
              if err != nil {
                  fmt.Println("Cannot decrypt key file: ", err)
                  return
              }
              privateKey := accountKey.PrivateKey

             //  private := "private:" + list[2]
             //  //Set key in redis
             //  err = Redis_client.Set(private,string(crypto.FromECDSA(privateKey)), 0).Err()
             //  if err != nil {
             //     panic(err)
             //  }
              //Add to array list
              wallet := WalletAccount{
                  Address: list[2],
                  Nonce: uint64(0),
                  PrivateKey: privateKey,
              }
              Wallets = append(Wallets,&wallet)
         }
    }
}
