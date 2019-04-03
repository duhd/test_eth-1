package utils

import (
  "fmt"
  // "encoding/json"
  // "time"
  // "context"
  // "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/core/types"
  // "github.com/go-redis/redis"
  "sync"
  	"log"
)


type WorkerPool struct {
   TxCh chan *types.Header
   HttpUrl string
   Clients []*EthClient
   Current int
   mutex sync.Mutex
}

func NewWorkerPool(httpUrl string)  *WorkerPool {

    //Create RPC connections
    var clients  []*EthClient
    max_client := cfg.Webserver.MaxListenRpcConnection
    for i:=0; i< max_client; i++ {
       ethclient, err := NewEthClient(httpUrl)
       if err != nil {
           log.Fatal("Cannot connect to: ",httpUrl," error:", err)
           continue
       }
       clients = append(clients,ethclient)
     }

     //Create channel to query transactions
     txCh := make(chan *types.Header)

     //Create transaction
     workerpool :=  &WorkerPool{
          TxCh: txCh,
          HttpUrl: httpUrl,
          Clients: clients,
          Current: 0,
     }
     return workerpool
}

func (wp *WorkerPool) getClient() *EthClient {
  wp.mutex.Lock()
  defer wp.mutex.Unlock()

  len := len(wp.Clients)
  if wp.Current >=  len {
      wp.Current = wp.Current % len
  }
  client := wp.Clients[wp.Current]
  wp.Current = wp.Current + 1
  return client
}
func (wp *WorkerPool) LoopQueryTransaction(){
    for {
          select {
                case header := <-wp.TxCh:
                      fmt.Println("Query transaction")
                      //Query transaction
                      client := wp.getClient()
                      client.UpdateReceipt(header)

            }
    }
}
func (wp *WorkerPool) QueryTransaction(header *types.Header){
    wp.TxCh <- header
}
