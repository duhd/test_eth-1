package utils

import (
    "fmt"
    // "context"
    // "crypto/ecdsa"
    "sync/atomic"
    "github.com/ethereum/go-ethereum/common"
    // "github.com/ethereum/go-ethereum/ethclient"
)

func (w *WalletAccount) GetNonce() uint64 {
    nonce := atomic.AddUint64(&w.Nonce, 1)
    fmt.Println("WalletAccount.GetNonce: Get Nonce:",nonce)
    return nonce
}

func (w *WalletAccount) UpdateNonce(nonce uint64)  {
    fmt.Println("Update: ",w.Address," Nonce:",nonce)
    atomic.StoreUint64(&w.Nonce, nonce-1)
}


func (w *WalletAccount) SyncNonce(){
  keyAddr := common.HexToAddress(w.Address)
  nonce, err := w.Routing.PendingNonceAt(keyAddr)
  if err != nil {
    fmt.Errorf("failed to retrieve account nonce: %v", err)
    w.UpdateNonce(0)
    return
  }
  fmt.Println("Nonce from eth: ",nonce)
  w.UpdateNonce(nonce)
}

func (w *WalletAccount) IsAvailable() bool {
  w.Mutex.Lock()
  defer w.Mutex.Unlock()

  if w.Active {
      fmt.Println("Account: ",w.Address," is active")
      keyAddr := common.HexToAddress(w.Address)
      nonce, err := w.Routing.PendingNonceAt(keyAddr)
      fmt.Println("Account: ",w.Address," nonce: ", nonce," compare to ",w.Nonce)
      if err != nil {
        return false
      }
      if nonce == 0 {
          w.Nonce = 0
          return true
      }
      if nonce == w.Nonce {
          return false
      }
      if nonce > w.Nonce {
          w.Nonce = nonce
          return true
      }
      w.Nonce = nonce
  }
  return false
}
