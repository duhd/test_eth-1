package utils

import (
  "fmt"
  // "encoding/json"
  // "time"
  	"context"
  	"github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/core/types"
  // "github.com/go-redis/redis"
  "sync"
  	"log"
)

type BlockSubscriber struct {
    Blocks map[string]string
    mutex sync.RWMutex
}

func NewBlockSubscriber() *BlockSubscriber{
  blocks := make(map[string]string)
  return &BlockSubscriber{
    Blocks: blocks,
  }
}
func (bl *BlockSubscriber) Get(number string) (string, bool){
  bl.mutex.Lock()
  defer bl.mutex.Unlock()
  v,k := bl.Blocks[number]
  return string(v), k
}
func (bl *BlockSubscriber) Set(number string,value string){
  bl.mutex.Lock()
  defer bl.mutex.Unlock()
  bl.Blocks[number] = value
}

type Subscriber struct {
   Name string
   SocketUrl string
   Workers *WorkerPool
   Blocks *BlockSubscriber
}

func NewSubscriber(name string, httpUrl string,socketUrl string,blocks *BlockSubscriber)  *Subscriber {
     workerpool := NewWorkerPool(httpUrl)
     //Create transaction
     subscriber :=  &Subscriber{
       Name: name,
       SocketUrl: socketUrl,
       Workers: workerpool,
       Blocks:blocks,
     }
     return subscriber
}

func (sb *Subscriber) CheckHeader(header *types.Header){
    //Query redis
    blNumber := header.Number.String()

    fmt.Println("Subscriber:",sb.Name,"Check block: ",blNumber)
    if value, ok := sb.Blocks.Get(blNumber); ok {
         if value != sb.Name {
             fmt.Println("Call worker to get transaction from block:",blNumber)
             sb.Workers.QueryTransaction(header)
         }else{
           fmt.Println("Same subscriber received same block :",blNumber)
         }
    } else {
        fmt.Println("Not find blockNumber:",blNumber)
        sb.Blocks.Set(blNumber,sb.Name)
    }

}

func (sb *Subscriber) ListenBlockEvent(){
		fmt.Println("Subscriber:", sb.Name ,"Listening from: ", sb.SocketUrl)
		websocket, err := ethclient.Dial("ws://" + sb.SocketUrl)
		if err != nil {
				log.Fatal("Cannot connect to websocket", err)
				return
		}
		headers := make(chan *types.Header)
		sub, err := websocket.SubscribeNewHead(context.Background(), headers)
		if err != nil {
		    fmt.Println("Cannot SubscribeNewHead to host: ", sb.SocketUrl ," Error: ",err)
				return
		}
	  fmt.Println("Start listening: ",sb.SocketUrl,"  ")
		for {
					select {
								case err := <-sub.Err():
										fmt.Println("Error from: ",sb.SocketUrl," Error: ",err)
										log.Fatal(err)
								case header := <-headers:
                   fmt.Println("Block Number: ", header.Number.String()," Subscriber: ", sb.Name, " call CheckHeader")
                    //Process header
                    go func(){
                        sb.CheckHeader(header)
                    }()
						}
		}
}

func (sb *Subscriber) Start(){
    go func (){
        fmt.Println("Loop Subscriber waiting event ")
        sb.ListenBlockEvent()
    }()
    go func (){
        fmt.Println("Loop Subscriber query transactions ")
        sb.Workers.LoopQueryTransaction()
    }()
}
