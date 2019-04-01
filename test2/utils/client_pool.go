package utils

import (
  "time"
  "github.com/ethereum/go-ethereum/core/types"
  "fmt"
)


type TxTransaction struct {
  Data *types.Transaction
  Nonce uint64
}

type ClientPool struct {
  Clients []*EthClient
  Current int
  TxCh chan *TxTransaction
}


var clientPool *ClientPool


func NewClientPool()  *ClientPool{
    var clients []*EthClient
    max_connection := cfg.Webserver.MaxRpcConnection
    for i:=0 ; i<max_connection; i++ {
         for _,host := range cfg.Networks {
              ethclient, err := NewEthClient(host.Http)
              if err != nil {
                continue
              }
              clients = append(clients,ethclient)
            }
     }
     txCh := make(chan *TxTransaction,cfg.Channel.TransferQueue)
     clientPool =  &ClientPool{
        Clients: clients,
        Current: 0,
        TxCh: txCh,
     }
     return clientPool
}

func (cp *ClientPool) GetClient() (*EthClient) {
    if cp.Current >=  len(cp.Clients) {
        cp.Current = cp.Current % len(cp.Clients)
    }
    client := cp.Clients[cp.Current]
    cp.Current = cp.Current + 1
    return client
}

func (cp *ClientPool) TransferToken(signedTx *types.Transaction, nonce uint64){
//  fmt.Println("Send Transaction to channel")
  tx := &TxTransaction{
    Data: signedTx,
    Nonce: nonce,
  }
  cp.TxCh <-tx
}

func (cp *ClientPool) Process(){
  for {
      select {
            case  tx:= <- cp.TxCh:
               go func() {
                  fmt.Println("Get Transaction from channel")
                   start := time.Now().UnixNano()
                   client := cp.GetClient()
                  // fmt.Println("Submit Transaction to geth")
                    _, err :=  client.TransferToken(tx.Data, tx.Nonce)
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
