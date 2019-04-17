package utils

import (
  "fmt"
  // "encoding/json"
  "time"
  "context"
  "strings"
  // "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/core/types"
  // "github.com/go-redis/redis"
  "sync"
  	"log"
)


type WorkerPool struct {
   TxCh chan *types.Header
   HttpUrl string
   Clients []*RpcConnection
   Current int
   mutex sync.Mutex
}

func NewWorkerPool(httpUrl string)  *WorkerPool {

    //Create RPC connections
    var clients  []*RpcConnection
    max_client := cfg.Webserver.MaxListenRpcConnection
    for i:=0; i< max_client; i++ {
       ethclient, err := NewRpcConnection("",httpUrl)
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

func (wp *WorkerPool) getClient() *RpcConnection {
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

func (wp *WorkerPool) UpdateReceipt(header *types.Header ){
      conn := wp.getClient()

      conn.Mux.Lock()
      defer  conn.Mux.Unlock()

      block, err := conn.Client.BlockByHash(context.Background(), header.Hash())
      if err != nil {
        fmt.Println("Error block by hash: ",err)
        return
        //log.Fatal(err)
      }
      t := time.Now()
      fmt.Println(t.Format(time.RFC822),"Block Number: ", header.Number.String(),"number of transactions:", len(block.Transactions()), " header hash: " , header.Hash().Hex())
      coinbase := block.Coinbase()
      for _, transaction := range block.Transactions(){
           nonce := transaction.Nonce()
           key := strings.TrimPrefix(transaction.Hash().Hex(),"0x")
           redisCache.LogEnd(key,nonce,coinbase.Hex())
      }
}

func (wp *WorkerPool) LoopQueryTransaction(){
    for {
          select {
                case header := <-wp.TxCh:
                      fmt.Println("Query transaction",header)
                      //Query transaction
                     wp.UpdateReceipt(header)

            }
    }
}
func (wp *WorkerPool) QueryTransaction(header *types.Header){
    wp.TxCh <- header
}
