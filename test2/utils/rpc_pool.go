package utils

import (
  // "time"
  // "github.com/ethereum/go-ethereum/core/types"
  "fmt"
  "sync"
  "context"
)

type EthNode struct {
  Name string
  Connections []*RpcConnection
  Current int
  Active bool
  MaxConn int
  HttpUrl string
  mutex sync.Mutex
}

func NewEthNode(name string, maxConn int, httpurl string) *EthNode{
      var rpc_conns []*RpcConnection
      for i:=0 ; i<maxConn; i++ {
          ethclient, err := NewRpcConnection(name, httpurl)
          if err != nil {
            fmt.Println("Failed to connect to: " + httpurl)
            continue
          }
          rpc_conns = append(rpc_conns,ethclient)
      }
     ethNode :=  &EthNode{
         Name: name,
         Connections: rpc_conns,
         MaxConn: maxConn,
         HttpUrl: httpurl,
         Current: 0,
         Active: false,
     }
     ethNode.UpdateHealth()
     return ethNode
}

func (n *EthNode) GetConnection() (*RpcConnection) {
    n.mutex.Lock()
    defer n.mutex.Unlock()

    len := len(n.Connections)
    if n.Current >=  len {
        n.Current = n.Current % len
    }
    conn := n.Connections[n.Current]
    n.Current = n.Current + 1
    return conn
}

func (n *EthNode) UpdateHealth() {
    fmt.Println("Node: ",n.Name, " retry active")
    conn := n.GetConnection()

    conn.Mux.Lock()
    defer conn.Mux.Unlock()
    if _, err := conn.Client.NetworkID(context.Background()); err != nil {
      fmt.Println("Try connect to host: ",n.Name, ", Error: ", err)
      return
    }
    //Active ok
    n.mutex.Lock()
    defer n.mutex.Unlock()
    n.Active = true
    fmt.Println("Node: ",n.Name, " succesfully active")

}
