package main

import (
	"os"
	// "strings"
	"fmt"
	// "github.com/ethereum/go-ethereum/core/types"
  "test_eth/test2/utils"
	// "github.com/go-redis/redis"
	"sync"
	// "time"
)

var cfg *utils.Config


func main() {
	 var wg sync.WaitGroup

	 config_file := "config.yaml"
	 if len(os.Args) == 2 {
			config_file = os.Args[1]
	 }

	 	cfg = utils.LoadConfig(config_file)

		//Creat redis Poool
		redisPool := utils.NewRedisPool()
		wg.Add(1)
		//Loop waiting to process log
		go func (){
				fmt.Println("Loop Redis Pool ")
				defer wg.Done()
				redisPool.Process()
		}()

		subPool := utils.NewSubscriberPool()
		subPool.Start()
		
	  wg.Wait()
}
