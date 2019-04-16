package main

import (
		// "os"
		"strings"
	  "context"
		"github.com/ethereum/go-ethereum/ethclient"
		"github.com/ethereum/go-ethereum/accounts/abi/bind"
		"fmt"
		"time"
		"github.com/ethereum/go-ethereum/common"
	  "test_eth/test2/utils"
		"sync"
	  "test_eth/contracts/f5coin"
	  // "github.com/ethereum/go-ethereum/accounts/abi"
)
func str2addr(str string) []common.Address {
	 list := strings.Split(str,",")
	 ret :=  make([]common.Address,0)
	 for _, element := range list {
		 fmt.Println("Addr: ",element)
		  addr := common.HexToAddress(element)
      ret = append(ret,addr)
	 }
	 return ret
}

var cfg *utils.Config
var instance *f5coin.Business
var auth *bind.TransactOpts

func init(){
	 config_file := "config.yaml"
	 cfg = utils.LoadConfig(config_file)


	node := cfg.Networks[0]
	Name := node.Name
	SocketUrl := node.WebSocket
	contractAddr := cfg.F5Contract.Address
	fmt.Println("Subscriber:", Name ,"Listening from: ", SocketUrl)
	websocket, err := ethclient.Dial("ws://" + SocketUrl)
	if err != nil {
			fmt.Println("Cannot connect to websocket", err)
			return
	}
	fmt.Println("Listen event from address:", contractAddr)
	contractAddress := common.HexToAddress(contractAddr)
	instance, err = f5coin.NewBusiness(contractAddress, websocket)
	if err != nil {
			fmt.Println("Unable to bind to deployed instance of contract")
			return
	}
	fmt.Println("Start listening")

}

func WatchEventRegisterAccETH() {
	  fmt.Println("Watch Event RegisterAccETH ")
		eventCh := make(chan *f5coin.BusinessEventRegisterAccETH)
		sub,err := instance.WatchEventRegisterAccETH(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
									 fmt.Println("time:",time.Now(),", Lenght of Addresses: ", len( event.ListAcc) )
						}
		 }
}
func WatchEventCreateStash() {
	  fmt.Println("Watch Event CreateStash")
		eventCh := make(chan *f5coin.BusinessEventCreateStash)
		sub,err := instance.WatchEventCreateStash(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh,nil)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
									fmt.Println("time:",time.Now(),", WalletAddress: ", event.WalletAddress.Hex(),", Code: ", string(event.WalletCode[:]) )
						}
		 }
}
func WatchEventSetState() {
	  fmt.Println("Watch Event CreateStash")
		eventCh := make(chan *f5coin.BusinessEventSetState)
		sub,err := instance.WatchEventSetState(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh,nil)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
									fmt.Println("time:",time.Now(),", Set State:  WalletCode: ", string(event.WalletCode[:]),", StashState: ", event.StashState)
						}
		 }
}

func WatchEventCredit() {
	  fmt.Println("Watch Business Event Credit ")
		eventCh := make(chan *f5coin.BusinessEventCredit)
		sub,err := instance.WatchEventCredit(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh,nil,nil)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
									 fmt.Println("time:",time.Now(),", Credit TxRef: ",  string(event.TxRef[:]), ", WalletCode:",  string(event.WalletCode[:]), ", Amount: ",event.Amount )
						}
		 }
}
func WatchEventDebit() {
	  fmt.Println("Watch Business Event Debit ")
		eventCh := make(chan *f5coin.BusinessEventDebit)
		sub,err := instance.WatchEventDebit(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh,nil,nil)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
								 fmt.Println("time:",time.Now(),", Debit TxRef: ",  string(event.TxRef[:]), ", WalletCode:", string(event.WalletCode[:]), ", Amount: ",event.Amount )
						}
		 }
}
func WatchEventTransfer() {
	  fmt.Println("Watch Business Event Debit ")
		eventCh := make(chan *f5coin.BusinessEventTransfer)
		sub,err := instance.WatchEventTransfer(&bind.WatchOpts{Start: nil,  Context: context.Background()},eventCh,nil,nil,nil)
		if err != nil {
				fmt.Println("Cannot create watch instance: ", err )
				return
		}
		defer sub.Unsubscribe()
		 for {
				 select {
							 case event := <-eventCh:
								 fmt.Println("time:",time.Now(),", Transfer TxRef: ",  string(event.TxRef[:]), ", Sender:", string(event.Sender[:]) , ", Receiver:",string(event.Receiver[:]), ", Amount: ",event.Amount , ", Note: ",event.Note , ", TxType: ",event.TxType )
						}
		 }
}
func main() {
	  var wg sync.WaitGroup
	  wg.Add(6)

		go func() {
			println("Loop Watch Even tRegisterAccETH ")
			defer wg.Done()
			 WatchEventRegisterAccETH()
		 }()

	 	go func() {
			println("Loop Watch Event CreateStash ")
			defer wg.Done()
			WatchEventCreateStash()
		}()
		go func() {
			println("Loop Watch Event SetState ")
			defer wg.Done()
			WatchEventSetState()
		}()
		go func() {
			println("Loop Watch Event Credit ")
			defer wg.Done()
			WatchEventCredit()
		}()
		go func() {
			println("Loop Watch Event Debit ")
			defer wg.Done()
			WatchEventDebit()
		}()
		go func() {
			println("Loop Watch Event Transfer ")
			defer wg.Done()
			WatchEventTransfer()
		}()
		 wg.Wait()
		 fmt.Println("Finished Event Listening")
}


// func ReadEventCreateStash(contractAddr string, Name string, SocketUrl string) {
// 	fmt.Println("Subscriber:", Name ,"Listening from: ", SocketUrl)
// 	websocket, err := ethclient.Dial("ws://" + SocketUrl)
// 	if err != nil {
// 			fmt.Println("Cannot connect to websocket", err)
// 			return
// 	}
// 	fmt.Println("Listen event from address:", contractAddr)
// 	contractAddress := common.HexToAddress(contractAddr)
// 	instance, err := f5coin.NewBusiness(contractAddress, websocket)
// 	if err != nil {
// 			fmt.Println("Unable to bind to deployed instance of contract")
// 			return
// 	}
// 	fmt.Println("Start listening")
//
// 	evtIterator,err := instance.FilterEventCreateStash(&bind.FilterOpts{Start: 0, End: nil,  Context: context.Background()},nil)
// //	evtIterator,err := instance.FilterEventRegisterAccETH(&bind.FilterOpts{Start: 0, End: nil,  Context: context.Background()})
//   if err != nil {
//  		 fmt.Println("Failed to execute a filter query command", "err", err)
//  		 return
//   }
//   for  evtIterator.Next() {
//  			fmt.Println("time:",time.Now(),", WalletAddress: ", evtIterator.Event.WalletAddress.Hex(),", Code: ", string(evtIterator.Event.WalletCode[:]) )
// 			//	fmt.Println("time:",time.Now(),", Len: ",len(evtIterator.Event.ListAcc) )
//   }
//   fmt.Println("Finished")
// }
//
// func WatchLogs(contractAddr string, Name string, SocketUrl string){
// 	caller, err := ethclient.Dial("ws://" + SocketUrl)
// 	if err != nil {
// 			fmt.Println("Cannot connect to websocket", err)
// 			return
// 	}
// 	contractAddress := common.HexToAddress(contractAddr)
// 	abi, err := abi.JSON(strings.NewReader(f5coin.BusinessABI))
// 	if err != nil {
// 			fmt.Println("Error parse contract ABI", err)
// 		return
// 	}
// 	boundContract := bind.NewBoundContract(contractAddress, abi, caller, caller, caller)
// 	opts := &bind.WatchOpts{Start: nil,  Context: context.Background()}
//
// 	//event_name := "event_createStash"
// 	event_name := "event_registerAccETH"
// 	fmt.Println("boundContract WatchLogs: ")
//   logs, sub, err  := boundContract.WatchLogs(opts,event_name)
// 	if err != nil {
// 		fmt.Println("Error in watchLog: ",event_name)
// 		return
// 	}
//   defer sub.Unsubscribe()
//
// 	for {
// 		select {
// 				case log := <-logs:
// 					fmt.Println("Receive Log: ",log)
// 					// New log arrived, parse the event and forward to the user
// 					//event := new(f5coin.BusinessEventCreateStash)
// 					event := new(f5coin.BusinessEventRegisterAccETH)
// 					if err :=  boundContract.UnpackLog(event, event_name, log); err != nil {
// 						fmt.Println("Error in UnpackLog: ",event_name)
// 						return
// 					}
// 					event.Raw = log
// 					//fmt.Println("time:",time.Now(),", WalletAddress: ", event.WalletAddress.Hex(),", Code: ", string(event.WalletCode[:]) )
// 					fmt.Println("time:",time.Now(),", Len: ",len(event.ListAcc) )
// 				case err := <-sub.Err():
// 					fmt.Println("Subscription error: ",err)
// 					return
// 				// case <-quit:
// 				// 	return
// 		}
// 	}
// }
