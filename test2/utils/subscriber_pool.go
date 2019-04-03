package utils

import (
  "fmt"
  // "encoding/json"
  // "time"
  	// "context"
  // 	"github.com/ethereum/go-ethereum/ethclient"
  // "github.com/ethereum/go-ethereum/core/types"
  // "github.com/go-redis/redis"
  // "sync"
  	// "log"
)



type SubscriberPool struct {
    Subscribers []*Subscriber
    Blocks *BlockSubscriber
}

var subpool *SubscriberPool

func NewSubscriberPool() *SubscriberPool{
    blockSubscriber := NewBlockSubscriber()

    var subscribers []*Subscriber
    for _,host := range cfg.Networks {
        sb := NewSubscriber(host.Name,host.Http,host.WebSocket,blockSubscriber)
        subscribers = append(subscribers,sb)
    }

    subpool := &SubscriberPool{
      Subscribers:subscribers,
      Blocks: blockSubscriber,
    }
    return subpool
}

func (sp *SubscriberPool) Start(){
    for _,sub := range sp.Subscribers {
      fmt.Println("Start subscriber: ",sub.Name)
      sub.Start()
    }
}
