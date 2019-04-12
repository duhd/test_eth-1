package utils

import (
  "time"
  "math/big"
  // "strings"
  "fmt"
  // "encoding/json"
  "errors"
  "strings"
  "github.com/ethereum/go-ethereum/crypto"
    _ "github.com/jinzhu/gorm/dialects/mysql"
      "github.com/jinzhu/gorm"
    "encoding/hex"
)

type WalletHandler struct {
    Client *RpcRouting
    Wallets []*WalletAccount
    Address string
}

func NewWalletHandler( contract_address string, client *RpcRouting)  *WalletHandler{
      wallHandler :=  &WalletHandler{
        Client: client,
        Address: contract_address,
      }
      wallHandler.LoadWallets()
      return wallHandler
}

func (wh *WalletHandler) NewTokenAccount() (string, error){
   privateKey, err := crypto.GenerateKey()
   if err != nil {
     return "",err
   }

   address := crypto.PubkeyToAddress(privateKey.PublicKey)

   account := address.Hex()
   account = strings.TrimPrefix(account,"0x")
   account = strings.ToLower(account)

   priKey :=  hex.EncodeToString(crypto.FromECDSA(privateKey))

   new_account := &TokenAccount{
     Address: account,
     PrivateKey: priKey,
     Active: false,
   }

   fmt.Println("Update account to db ")
   db, err := gorm.Open("mysql", cfg.MysqlConnectionUrl())
   if cfg.Mysql.Debug {
      db.LogMode(true)
   }

   if err != nil {
     panic("failed to connect database: " + err.Error())
   }
   defer db.Close()
   //fmt.Println("Create new record")
   db.Create(new_account)

   fmt.Println("Update account to wallet ")
   wallet := WalletAccount{
     Routing: wh.Client,
     Address: account,
     PrivateKey: privateKey,
     Nonce: 0,
     Account: new_account,
   }
   wh.Wallets = append(wh.Wallets,&wallet)
   return account, nil
}

func (wh *WalletHandler) LoadWallets() {
  fmt.Println("Start load accounts from db to create wallets ")
  db, err := gorm.Open("mysql", cfg.MysqlConnectionUrl())
  if cfg.Mysql.Debug {
     db.LogMode(true)
  }

  if err != nil {
    panic("failed to connect database: " + err.Error())
  }
  defer db.Close()

  accounts := []TokenAccount{}

  if err := db.Where("active = ?", true).Find(&accounts).Error; err != nil {
    fmt.Println("Cannot get active Token Account in db: ",err)
    return
  }
  wallets := []*WalletAccount{}
  for _, account := range accounts {
      fmt.Println("Load wallet: ",account.Address)
      b, err := hex.DecodeString(account.PrivateKey)
    	if err != nil {
          fmt.Println("invalid hex string: " + account.PrivateKey)
    		  continue
    	}
      privateKey := crypto.ToECDSAUnsafe(b)
      wallet := WalletAccount{
        Routing: wh.Client,
        Address: account.Address,
        PrivateKey: privateKey,
        Account: &account,
      }
      fmt.Println("Start sync nonce of ",account.Address)
      wallet.SyncNonce()
      wallets = append(wallets,&wallet)
  }
  fmt.Println("End load accounts from db: ", len(wallets))
  wh.Wallets = wallets
}

func (wh *WalletHandler) GetWallet(addr string) *WalletAccount {
    for _, wallet := range wh.Wallets {
       if wallet.Address == addr {
         return wallet
       }
    }
    return nil
}

func (wh *WalletHandler) EthBalaneOf(account string) (*big.Float,error) {
  wallet := wh.GetWallet(account)
  if wallet != nil {
      return wallet.EthBalaneOf()
  }
  return nil, errors.New("Cannot find account in system")
}

func (wh *WalletHandler) BalaneOf(account string) (*big.Float,error) {
  wallet := wh.GetWallet(account)
  return wallet.BalaneOf()
}

func (wh *WalletHandler) EthTransfer(from string,to string,amount string) (string,error) {
   wallet := wh.GetWallet(from)
   txhash, _, err := wallet.EthTransfer(to,amount)
   return txhash, err
}
func (wh *WalletHandler) TransferToken(from string,to string,amount string,append string) (string,error) {
  requestTime := time.Now().UnixNano()
  //1, Query account and nonce
  wallet := wh.GetWallet(from)

  txhash, nonce, err := wallet.TransferToken(from,to,amount,append)

  submitTime := time.Now().UnixNano()

  // fmt.Println("Send message to log server (redis pool)")
  redisCache.LogStart(txhash, nonce, requestTime)

  logTime := time.Now().UnixNano()

  diff0 := (submitTime - requestTime)/1000
  diff1 := (logTime - submitTime)/1000
  fmt.Println("Transfer: ", nonce," from ",from," to ",to, " amount: ",amount, " note:",append)
  fmt.Println("SubmitTime, LogTime : ",diff0,diff1, " Transaction =",txhash)

  return txhash, err
}

func (wh *WalletHandler) GetAccountList() ([]string, error) {
   fmt.Println("Handler.GetAccountList: start read wallets")
   wh.LoadWallets()
   accounts := []string{}
   for _,wallet := range wh.Wallets {
       accounts = append(accounts,wallet.Address)
   }
   return accounts, nil
}

func (wh *WalletHandler) GetAccountKey(account string) (string, error) {
  for _,wallet := range wh.Wallets {
     if wallet.Address == account {
        return wallet.GetPrivateKey(), nil
     }
  }
  return "", errors.New("Not found account: " + account)
}
