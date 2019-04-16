package utils

import (
  // "time"
  "math/big"
  // "strings"
  "fmt"
  // "encoding/json"
  "errors"
  "strings"
  "github.com/ethereum/go-ethereum/crypto"
  	"github.com/ethereum/go-ethereum/common"
    _ "github.com/jinzhu/gorm/dialects/mysql"
      "github.com/ethereum/go-ethereum/core/types"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/jinzhu/gorm"
  "encoding/hex"
  "test_eth/contracts/f5coin"
  "sync"
  "math"
)

type F5WalletHandler struct {
    Client *RpcRouting
    Wallets []*WalletAccount
    ContractAddress common.Address
    Current int
    Mutex sync.Mutex
}
func stringTo32Byte(data string) [32]byte {
  //hexstring := hex.EncodeToString([]byte(data))
  var arr [32]byte
	copy(arr[:], data)
  return arr
}

func NewF5WalletHandler(contract_address string, client *RpcRouting)  *F5WalletHandler{
      contractAddress := common.HexToAddress(contract_address)
      wallHandler :=  &F5WalletHandler{
        Client: client,
        ContractAddress: contractAddress,
        Current: 0,
      }
      // wallHandler.LoadAccountEth()
      // wallHandler.AutoFillGas()
      // wallHandler.RegisterBatchEthToContract()
      return wallHandler
}
func (fw *F5WalletHandler) RegisterBatchEthToContract() []string {
    ret := []string{}
    list := fw.GetAccountList()
    j := 0
    sublist :=  []common.Address{}
    for _,item := range list {
      if j == 0 {
         sublist = []common.Address{}
      }
      sublist = append(sublist,item)

      tx,err := fw.RegisterAccETH(sublist)
      if err != nil {
        ret = append(ret, err.Error())
      } 	else {
         ret = append(ret, tx.Hash().Hex())
      }
      j = ( j + 1 ) % 5
    }
    return ret
}
func (fw *F5WalletHandler) NewAccountEth() (string, error) {
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
        Active: true,
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
        Routing: fw.Client,
        Address: account,
        PrivateKey: privateKey,
        Nonce: 0,
        Account: new_account,
        Active: false,
      }
      fw.Wallets = append(fw.Wallets,&wallet)
      return account, nil
}

func (fw *F5WalletHandler) GetAccountEthAddress(addr string) *WalletAccount {
    for _, wallet := range fw.Wallets {
       if wallet.Address == addr {
         return wallet
       }
    }
    return nil
}

func (fw *F5WalletHandler) GetAccountEth() *WalletAccount{
    fw.Mutex.Lock()
    defer fw.Mutex.Unlock()
    len := len(fw.Wallets)
    if len == 0 {
      return nil
    }
    if fw.Current >= len {
         fw.Current = fw.Current % len
    }
    wallet := fw.Wallets[fw.Current]
    fw.Current = fw.Current + 1
    return wallet
}

func (fw *F5WalletHandler) LoadAccountEth(){
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
        Routing: fw.Client,
        Address: account.Address,
        PrivateKey: privateKey,
        Active: true,
        Account: &account,
        Nonce: 0,
      }

      if cfg.Webserver.NonceMode == 2 {
          fmt.Println("Start sync nonce of ",account.Address)
          wallet.SyncNonce()
      }
      wallets = append(wallets,&wallet)
  }
  fmt.Println("End load accounts from db: ", len(wallets))
  fw.Mutex.Lock()
  defer fw.Mutex.Unlock()
  fw.Wallets = wallets
}

func (fw *F5WalletHandler) CreateStash(stashName string, typeStash int8) (*types.Transaction, error)  {
    retry := 0
    for retry <10 {
        account := fw.GetAccountEth()
        if account.IsAvailable() {
          conn := fw.Client.GetConnection()
          session, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          auth := account.NewTransactor()
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          return session.CreateStash(auth,stringTo32Byte(stashName), typeStash)
        }
        retry = retry + 1
    }
    return nil, errors.New("Cannot find wallet in pool to create transaction")
}
func (fw *F5WalletHandler) GetBalance(stashName string) (*big.Int, error)  {
    fmt.Println("F5WalletHandler.GetBalance: Start get balance ")
    conn := fw.Client.GetConnection()
    conn.Mux.Lock()
    defer  conn.Mux.Unlock()
    session,err  := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
    if err != nil {
        fmt.Println("Cannot find F5 contract")
        return nil,err
    }
    fmt.Println("F5WalletHandler.GetBalance: call  GetBalance")
    return session.GetBalance(&bind.CallOpts{},stringTo32Byte(stashName))
}
// func (fw *F5WalletHandler) GetStateHistoryLength() (*big.Int, error)  {
//     conn := fw.Client.GetConnection()
//     conn.Mux.Lock()
//     defer  conn.Mux.Unlock()
//     session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
//     if err != nil {
//         fmt.Println("Cannot find F5 contract")
//         return nil,err
//     }
//     return session.GetStateHistoryLength(&bind.CallOpts{})
// }
func (fw *F5WalletHandler) SetState(txRef string, stashName string, stashState int8 ) (*types.Transaction, error)  {
  retry := 0
  for retry <10 {
      account := fw.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          conn := fw.Client.GetConnection()
          session, err  := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          return session.SetState(auth, stringTo32Byte(stashName),stashState)
      }
        retry = retry + 1
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}

func (fw *F5WalletHandler) GetState(stashName string) (int8, error)  {
  conn := fw.Client.GetConnection()
  conn.Mux.Lock()
  defer  conn.Mux.Unlock()

  session, err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
  if err != nil {
      fmt.Println("Cannot find F5 contract")
      return 0,err
  }
  return session.GetState(&bind.CallOpts{},stringTo32Byte(stashName))

}
// func (fw *F5WalletHandler) GetRedeemHistoryLength() (*big.Int, error)  {
//     conn := fw.Client.GetConnection()
//     conn.Mux.Lock()
//     defer  conn.Mux.Unlock()
//
//     session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
//     if err != nil {
//         fmt.Println("Cannot find F5 contract")
//         return nil,err
//     }
//     return session.GetRedeemHistoryLength(&bind.CallOpts{})
// }
func (fw *F5WalletHandler) Withdraw(txRef string, stashName string, amount *big.Int) (*types.Transaction, error) {
    retry := 0
    for retry <10 {
        account := fw.GetAccountEth()
        if account.IsAvailable() {
            auth := account.NewTransactor()
            conn := fw.Client.GetConnection()
            session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
            if err != nil {
                fmt.Println("Cannot find F5 contract")
                return nil,err
            }
            conn.Mux.Lock()
            defer  conn.Mux.Unlock()
            return session.Debit(auth, stringTo32Byte(txRef),stringTo32Byte(stashName),amount)
        }
          retry = retry + 1
    }
    return nil, errors.New("Cannot find wallet in pool to create transaction")
}
// func (fw *F5WalletHandler) GetPledgeHistoryLength() (*big.Int, error)  {
//     conn := fw.Client.GetConnection()
//     conn.Mux.Lock()
//     defer  conn.Mux.Unlock()
//     session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
//     if err != nil {
//         fmt.Println("Cannot find F5 contract")
//         return nil,err
//     }
//     return session.GetPledgeHistoryLength(&bind.CallOpts{})
// }
func (fw *F5WalletHandler) Deposit(txRef string, stashName string, amount *big.Int) (*types.Transaction, error) {
  retry := 0
  for retry <10 {
      account := fw.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          conn := fw.Client.GetConnection()
          session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          return session.Credit(auth, stringTo32Byte(txRef),stringTo32Byte(stashName),amount)
      }
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}
func (fw *F5WalletHandler) GetTransferHistoryLength() (*big.Int, error)  {
  conn := fw.Client.GetConnection()
  conn.Mux.Lock()
  defer  conn.Mux.Unlock()
  session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
  if err != nil {
      fmt.Println("Cannot find F5 contract")
      return nil,err
  }
  return session.GetTransferHistoryLength(&bind.CallOpts{})

}
func (fw *F5WalletHandler) Transfer(txRef string, sender string, receiver string, amount *big.Int, note string, txType int8) (*types.Transaction, error) {
  retry := 0
  for retry <10 {
      account := fw.GetAccountEth()
      if account.IsAvailable() {
          auth := account.NewTransactor()
          conn := fw.Client.GetConnection()
          session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          return session.Transfer(auth, stringTo32Byte(txRef),stringTo32Byte(sender),stringTo32Byte(receiver),amount,note,txType)
      }
        retry = retry + 1
  }
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}
func (fw *F5WalletHandler) RegisterAccETH(listAcc []common.Address) (*types.Transaction, error) {
  fmt.Println("Start RegisterAccETH")
  retry := 0
  for retry <10 {
      account := fw.GetAccountEth()
      if account == nil {
         fmt.Println("Cannot find active account")
         return nil, errors.New("Cannot find bugdet account")
      }
      if account.IsAvailable() {
          auth := account.NewTransactor()
          auth.GasLimit = 9000000
          conn := fw.Client.GetConnection()
          session,err := f5coin.NewBusiness(fw.ContractAddress,conn.Client)
          if err != nil {
              fmt.Println("Cannot find F5 contract")
              return nil,err
          }
          conn.Mux.Lock()
          defer  conn.Mux.Unlock()
          return session.RegisterAccETH(auth,listAcc)
      } else {
          fmt.Println("Account: ",account.Address," is unavailable ")
      }
      retry = retry + 1
  }
  fmt.Println("End RegisterAccETH: retry failed ")
  return nil, errors.New("Cannot find wallet in pool to create transaction")
}
func (fw *F5WalletHandler) GetAccountList() ([]common.Address) {
   fmt.Println("F5WalletHandler.GetAccountList: start read wallets")
   fw.Mutex.Lock()
   defer fw.Mutex.Unlock()
   accounts := []common.Address{}
   for _,wallet := range fw.Wallets {
       if wallet.Active {
         address := common.HexToAddress("0x"+wallet.Address)
         accounts = append(accounts,address)
       }
   }
   fmt.Println("F5WalletHandler.GetAccountList: end read wallets")
   return accounts
}

func (fw *F5WalletHandler) EthBalaneOf(account string) (*big.Float,error) {
  wallet := fw.GetAccountEthAddress(account)
  if wallet != nil {
      return wallet.EthBalaneOf()
  }
  return nil, errors.New("Cannot find account in system")
}
func (fw *F5WalletHandler) EthTransfer(from string,to string,amount string) (string,error) {
   wallet := fw.GetAccountEthAddress(from)

   fromAddress := common.HexToAddress("0x" + wallet.Address)
   nonce, err := wallet.Routing.PendingNonceAt(fromAddress)
   if err != nil {
     fmt.Println("Error in getting nonce ")
     return "", err
   }

   gLimit := cfg.Contract.GasLimit
   gPrice := cfg.Contract.GasPrice

   gasLimit := uint64(gLimit)
   gasPrice := new(big.Int)
   gasPrice, _ = gasPrice.SetString(gPrice, 10)

   toAddress := common.HexToAddress("0x" + to)

   eth_unit := big.NewFloat(math.Pow10(18))
   amount_value := new(big.Float)
   value, ok := amount_value.SetString(amount)

   if !ok {
        fmt.Println("SetString: error")
        return "", errors.New("convert amount error")
   }
   value = value.Mul(value,eth_unit)

   value_transfer := new(big.Int)
   value.Int(value_transfer)

   var data []byte
   rawTx := types.NewTransaction(nonce, toAddress, value_transfer, gasLimit, gasPrice, data)

   signer := types.FrontierSigner{}
   signature, err := crypto.Sign(signer.Hash(rawTx).Bytes(), wallet.PrivateKey)
   if err != nil {
     fmt.Println(" Cannot sign contract: ", err)
     return "",err
   }

   signedTx, err := rawTx.WithSignature(signer, signature)

   txhash := strings.TrimPrefix(signedTx.Hash().Hex(),"0x")
   err = wallet.Routing.SubmitTransaction(signedTx,nonce)

   return txhash, err
}

func (fw *F5WalletHandler) AutoFillGas() bool {
    fw.Mutex.Lock()
    defer fw.Mutex.Unlock()

    for _, wallet := range fw.Wallets {

      bal, err := wallet.EthBalaneOf()
      if err != nil {
         fmt.Println("Cannot get wallet balance. Deactive wallet")
         wallet.Active = false
         continue
      }
      ba,_ := bal.Float64()
      if ba < 1000 {
         fmt.Println("Create transaction to fillGass from budget")
         txhash, err := fw.EthTransfer(cfg.F5Contract.EthBudget, wallet.Address,"1000")
         if err != nil {
           fmt.Println("Cannot fill more gas. Deactive wallet ")
           wallet.Active = false
           continue
         }
         fmt.Println("Fill Eth to account: ", wallet.Address, " transaction: ", txhash)
      } else {
         fmt.Println("Account: ", wallet.Address, " balance: ", ba)
      }
    }
    return true
}
