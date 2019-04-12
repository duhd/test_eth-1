package utils

import (
   "context"
   "time"
   "fmt"
   "sync"
   "errors"
   "strings"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum"
    "math/big"
    "math"
)

type TxTransaction struct {
  Data *types.Transaction
  Nonce uint64
}


type RpcRouting struct {
  Nodes []*EthNode
  Current int
  TxCh chan *TxTransaction
  mutex sync.Mutex
  Mode int
}

var rpcrouting *RpcRouting

func NewRouting(cfg *Config)  *RpcRouting{
    max_conns := cfg.Webserver.MaxRpcConnection
    mode := cfg.Webserver.RoutingMode
    hosts := cfg.Networks

    var nodes []*EthNode
    for _,host := range hosts {
       node := NewEthNode(host.Name, max_conns, host.Http)
       nodes = append(nodes,node)
     }
     txCh := make(chan *TxTransaction,cfg.Channel.TransferQueue)

     rpcrouting =  &RpcRouting{
        Nodes: nodes,
        Current: -1,
        TxCh: txCh,
        Mode: mode,
     }
     return rpcrouting
}

func (r *RpcRouting) Process(){
      for {
          select {
          case <-time.After(5*time.Second):
               go func() {
                 r.UpdateHealth()
               }()
            case  tx:= <- r.TxCh:
                   go func() {
                      fmt.Println("Get Transaction Message from channel")
                       start := time.Now().UnixNano()
                        err :=  r.SendTransaction(tx.Data, tx.Nonce)
                        if err != nil {
                            fmt.Println("Error send transaction", tx.Nonce," error:", err)
                        }
                        end := time.Now().UnixNano()
                        diff:= (end-start)/1000
                        fmt.Println("End Submit transaction: ", tx.Nonce,", Time: ", diff)
                  }()
                }
        }
}
func (r *RpcRouting) UpdateHealth(){
    // r.mutex.Lock()
    // defer r.mutex.Unlock()
    fmt.Println("Process to check health of node. To automatically active node")
    for _, node := range r.Nodes {
      if !node.Active {
          node.UpdateHealth()
      }
   }
}
func (r *RpcRouting) DeactiveNode(name string)  {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    for _, node := range r.Nodes {
      if node.Name == name {
         node.Active  = false
      }
    }
}
func (r *RpcRouting) GetConnection() (*RpcConnection) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    len := len(r.Nodes)
    if len >0 {
      fmt.Println("Len : ", len, ", current: ", r.Current)
      if r.Current < 0 || r.Current >= len  {
        fmt.Println("Current < 0. Find first active node")
        for i, node := range r.Nodes {
          if node.Active {
             r.Current = i
             return node.GetConnection()
          }
        }
        return nil
      } else {

        node := r.Nodes[r.Current]
        if node.Active {
           return node.GetConnection()
        } else {
            fmt.Println("Current node not active. Find first active node")
            for i, node := range  r.Nodes {
              if node.Active {
                 r.Current = i
                 return node.GetConnection()
              }
            }
            return nil
        }
      }
    } else {
      fmt.Println("Not find node in list")
    }
    return nil
}

func (r *RpcRouting) SendTransaction(signedTx *types.Transaction,  nonce uint64) (error){
      retry := 0
      for retry < 3 {
          conn := r.GetConnection()
          if conn == nil {
             return errors.New("Not find connection. Please check")
          }
          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          err := conn.Client.SendTransaction(context.Background(), signedTx);
          if  err == nil {
            fmt.Println("Successfully send transaction nonce:", nonce  , ", to host: ", conn.Name )
            return nil
          }
         err_msg := err.Error()
         fmt.Println("Error in sending transaction with nonce:", nonce  , ", to host: ", conn.Name ,", error: ", err_msg)
         if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             r.DeactiveNode(conn.Name)
             continue
         }
         if strings.Contains(err_msg, "insufficient funds") {
             fmt.Println("Account not enough Ether to run transaction")
             return err
         }
         retry = retry + 1
      }
      return nil
}
func (r *RpcRouting) CallContract(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
      retry := 0
      for retry < 3 {
          conn := r.GetConnection()
          if conn == nil {
             return nil, errors.New("Not find connection. Please check")
          }

          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          bs, err := conn.Client.CallContract(context.Background(),msg,blockNumber);
          if  err == nil {
            fmt.Println("Successfully send message to host: ", conn.Name )
            return bs, nil
          }
         err_msg := err.Error()
         if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             r.DeactiveNode(conn.Name)
             retry = retry + 1
             continue
         }
         return bs, err
      }
      return nil, nil
}
func (r *RpcRouting) CodeAt(account common.Address, blockNumber *big.Int) ([]byte, error) {
      retry := 0
      for retry < 3 {
          conn := r.GetConnection()
          if conn == nil {
             return  nil, errors.New("Not find connection. Please check")
          }
          //Lock rpc connection before send
          conn.Mux.Lock()
          defer 	conn.Mux.Unlock()
          bs, err := conn.Client.CodeAt(context.Background(),account,blockNumber);
          if  err == nil {
            fmt.Println("Successfully send message to host: ", conn.Name )
            return bs,err
          }
          err_msg := err.Error()
          if strings.Contains(err_msg, "connection refused") {
             fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
             r.DeactiveNode(conn.Name)
             retry = retry + 1
             continue
          }
          return bs, err
      }
      return nil, nil
}
func (r *RpcRouting) NonceAt(account common.Address) (uint64, error) {
    nonce := uint64(0)
    for _, node := range r.Nodes {
         conn := node.GetConnection()

         conn.Mux.Lock()
         defer 	conn.Mux.Unlock()
         n, err := conn.Client.NonceAt(context.Background(),account,nil)
         if err != nil {
           continue
         }
         if n > nonce {
           nonce = n
         }
    }
    if nonce == 0 {
      return 0, errors.New("Cannot find pending nonce")
    }
    return nonce, nil
}
func (r *RpcRouting) PendingNonceAt(account common.Address) (uint64, error) {
    retry := 0
    for retry < 3 {
        conn := r.GetConnection()
        if conn == nil {
           return  0, errors.New("Not find connection. Please check")
        }
        //Lock rpc connection before send
        conn.Mux.Lock()
        defer 	conn.Mux.Unlock()
        n, err := conn.Client.PendingNonceAt(context.Background(),account)
        if  err == nil {
          fmt.Println("Successfully get nonce from: ", conn.Name )
          return n, err 
        }
        err_msg := err.Error()
        if strings.Contains(err_msg, "connection refused") {
           fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
           r.DeactiveNode(conn.Name)
           retry = retry + 1
           continue
        }
        //Other error
        return 0, err
    }
    return 0, errors.New("Cannot find nonce from system")
}
func (r *RpcRouting) MaxPendingNonceAt(account common.Address) (uint64, error) {
    nonce := uint64(0)
    for _, node := range r.Nodes {
         conn := node.GetConnection()

         conn.Mux.Lock()
         defer 	conn.Mux.Unlock()
         n, err := conn.Client.PendingNonceAt(context.Background(),account)
         if err != nil {
           continue
         }
         if n > nonce {
           nonce = n
         }
    }
    if nonce == 0 {
      return 0, errors.New("Cannot find pending nonce")
    }
    return nonce, nil
}

func (r *RpcRouting) SubmitTransaction(signedTx *types.Transaction, nonce uint64) (error) {
   	if r.Mode >1 {
          fmt.Println("Send Transaction to channel via message channel")
          tx := &TxTransaction{
            Data: signedTx,
            Nonce: nonce,
          }
          r.TxCh <-tx
    } else {
          fmt.Println("Submit Transaction directly to miner ")
          start := time.Now().UnixNano()

          err :=  r.SendTransaction(signedTx, nonce)
          if err != nil {
              fmt.Println("Error send transaction", nonce," error:", err)
              return err
          }
          end := time.Now().UnixNano()
          diff:= (end-start)/1000
          fmt.Println("End Submit transaction: ", nonce,", Time: ", diff)
    }
    return nil
}
func (r *RpcRouting) BalanceAt(account common.Address) (*big.Float, error) {
  retry := 0
  for retry < 3 {
      conn := r.GetConnection()
      if conn == nil {
         return  nil, errors.New("Not find connection. Please check")
      }
      //Lock rpc connection before send
      conn.Mux.Lock()
      defer 	conn.Mux.Unlock()

      balance, err := conn.Client.BalanceAt(context.Background(), account, nil)
      if  err == nil {
        fmt.Println("Successfully send message to host: ", conn.Name )
        fbalance := new(big.Float)
      	fbalance.SetString(balance.String())
      	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
        return ethValue, err
      }

      err_msg := err.Error()
      if strings.Contains(err_msg, "connection refused") {
         fmt.Println("Cannot connect server:", conn.Name,", deactive server. Try with others")
         r.DeactiveNode(conn.Name)
         retry = retry + 1
         continue
      }
      return nil, err
  }
  return nil, nil
}
