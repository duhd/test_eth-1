package main

import (
	"os"
	"bytes"
	"log"
	"test_eth/contracts/f5coin"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	  	"github.com/ethereum/go-ethereum/common"
	"math/big"
	// "io/ioutil"
	"fmt"
	"time"
	"bufio"
	"test_eth/test2/utils"
	  "strconv"
	// "gopkg.in/yaml.v2"
)
func stringTo32Byte(data string) [32]byte {
  var arr [32]byte
	copy(arr[:], data)
  return arr
}

var cfg * utils.Config
var instance *f5coin.Business
var auth *bind.TransactOpts

func init(){
	config_file := "config.yaml"
	cfg = utils.LoadConfigFromFile(config_file)
	// fmt.Println("Config", cfg)

	node := cfg.Networks[0]

	backend, err  := ethclient.Dial("http://"+node.Http)

	if err != nil {
		log.Fatalf("Unable to connect to network:%s with %v\n",node.Http, err)
	}
	keyjson := utils.LoadKey(cfg.Keys.KeyStore,cfg.F5Contract.Owner)

	auth, err = bind.NewTransactor(bytes.NewReader(keyjson), cfg.Keys.Password)
	if err != nil {
		 log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	gasPrice := new(big.Int)
	gasPrice.SetString(cfg.F5Contract.GasPrice,10)
	auth.GasPrice = gasPrice
	auth.GasLimit = cfg.F5Contract.GasLimitDefault

	contractAddress := common.HexToAddress(cfg.F5Contract.Address)

	instance, err = f5coin.NewBusiness(contractAddress,backend)

	if err != nil {
		log.Fatalf("Failed to create instance of contract: %v", err)
	}
}
func RegisterEth(){

		address := common.HexToAddress("0x" + cfg.F5Contract.Owner)
		list := []common.Address{address,}
		tx, err := instance.RegisterAccETH(auth,list)
		if err != nil {
			fmt.Println("Cannot RegisterAccETH: ", cfg.F5Contract.Owner, ", error: ",err)
			return
		}
		fmt.Println("Register transaction: ",tx.Hash().Hex())
		owner, err := instance.GetOwner(&bind.CallOpts{From: address})
		if err != nil {
			fmt.Println("Cannot Get Owner: ", cfg.F5Contract.Owner, ", error: ",err)
			return
		}
		fmt.Println("Owner is: ",owner.Hex())
		fmt.Println("Wait to get list of register accounts:.")
		time.Sleep(2*time.Second)
		n,err := instance.GetRegistedAccEthLength(&bind.CallOpts{From: address})
		if err != nil {
			fmt.Println("Cannot Get RegistedAccEthLength: ", cfg.F5Contract.Owner, ", error: ",err)
			return
		}
		fmt.Println("RegistedAccEthLength: ",n)


}
func CreateWallet(stashName string) {
		address := common.HexToAddress("0x" + cfg.F5Contract.Owner)
		tx, err := instance.CreateStash(auth,stringTo32Byte(stashName),int8(1))
		if err != nil {
			fmt.Println("Cannot create stash: ", stashName, ", error: ",err)
			return
		}
			fmt.Println("Register transaction: ",tx.Hash().Hex())
			fmt.Println("Wait to get list of Wallets.")
			time.Sleep(2*time.Second)
   		n,err := instance.GetStashNamesLenght(&bind.CallOpts{From: address})
			if err != nil {
				fmt.Println("Cannot Get Wallet Length: ", cfg.F5Contract.Owner, ", error: ",err)
				return
			}
			fmt.Println("Number of wallet: ",n)
}
func ActiveWallet(stashName string) {
	 address := common.HexToAddress("0x" + cfg.F5Contract.Owner)

		tx, err := instance.SetState(auth, stringTo32Byte(stashName),int8(1))
		if err != nil {
			fmt.Println("Cannot set state of stash: ", stashName, ", error: ",err)
			return
		}
		fmt.Println("Set State transaction: ",tx.Hash().Hex())
		fmt.Println("Sleep in 2 second")
		time.Sleep(2*time.Second)

		state, err := instance.GetState(&bind.CallOpts{From: address},stringTo32Byte(stashName))
		if err != nil {
			fmt.Println("Cannot get state of stash: ", stashName, ", error: ",err)
			return
		}
		fmt.Println("State of stash: ",stashName,", State: ", state)
}

func Balance(stashName string) {
	address := common.HexToAddress("0x" + cfg.F5Contract.Owner)

	bal, err := instance.GetBalance(&bind.CallOpts{From: address},stringTo32Byte(stashName))
	if err != nil {
		fmt.Println("Get balance of ", stashName, ", error: ",err)
		return
	}
	fmt.Println("Balance of ",stashName, ": ",bal)
}

func Deposit(stashName string, value int64) {

	txRef := "tx01"
	amount := big.NewInt(value)

	tx, err := instance.Credit(auth,stringTo32Byte(txRef), stringTo32Byte(stashName),amount)
	if err != nil {
		fmt.Println("Cannot deposit of stash: ", stashName, ", error: ",err)
		return
	}
	fmt.Println("Deposit transaction: ",tx.Hash().Hex())
}


func Withdraw(stashName string, value int64) {
	txRef := "tx01"
	amount := big.NewInt(value)

	tx, err := instance.Debit(auth,stringTo32Byte(txRef), stringTo32Byte(stashName),amount)
	if err != nil {
		fmt.Println("Cannot withdraw of stash: ", stashName, ", error: ",err)
		return
	}
	fmt.Println("Withdraw transaction: ",tx.Hash().Hex())
}
func Transfer(sender string, receiver string, value int64) {
	txRef := "tx01"
	amount := big.NewInt(value)
	note := "Test tranfer"
	txType := int8(1)

	tx, err := instance.Transfer(auth,stringTo32Byte(txRef),stringTo32Byte(sender),stringTo32Byte(receiver), amount, note, txType)
	if err != nil {
		fmt.Println("Cannot tranfer from: ", sender, " to: ",receiver," amount: ", value)
		return
	}
	fmt.Println("Transfer transaction: ",tx.Hash().Hex())
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	var text string
	for text != "q" {  // break the loop if text == "q"
			fmt.Println("1. Register EthUser ")
			fmt.Println("2. Create a wallet ")
			fmt.Println("3. Activate wallet ")
			fmt.Println("4. Balance")
			fmt.Println("5. Deposit ")
			fmt.Println("6. Withdraw ")
			fmt.Println("7. Transfer ")
			fmt.Print("Enter your choice: ")
			scanner.Scan()
			text = scanner.Text()
			switch text {
			case "q" :
				break;
			case "1":
				fmt.Println("1. Register EthUser : ", cfg.F5Contract.Owner)
				RegisterEth()
			case "2":
				fmt.Print("2. Create wallet. Please enter wallet name:")
				scanner.Scan()
				name := scanner.Text()
				CreateWallet(name)
				fmt.Println("")
			case "3":
				fmt.Print("3. Active wallet. Please enter wallet name:")
				scanner.Scan()
				name := scanner.Text()
				ActiveWallet(name)
				fmt.Println("")
			case "4":
				fmt.Print("4. View balance. Please enter wallet name:")
				scanner.Scan()
				name := scanner.Text()
				Balance(name)
				fmt.Println("")
			case "5":
				fmt.Print("5. Deposit. Please enter wallet name: ")
				scanner.Scan()
				name := scanner.Text()
				fmt.Println("")
				fmt.Print("Please enter amount: ")
				scanner.Scan()
				amount := scanner.Text()
				am, err := strconv.ParseInt(amount,10,64)
				if err != nil {
						// handle error
						fmt.Println(err)
						os.Exit(2)
				}
				Deposit(name,am)
				fmt.Println("")
			case "6":
				fmt.Print("6. Withdraw. Please enter wallet name: ")
				scanner.Scan()
				name := scanner.Text()
				fmt.Println("")
				fmt.Print("Please enter amount: ")
				scanner.Scan()
				amount := scanner.Text()
				am, err := strconv.ParseInt(amount,10,64)
				if err != nil {
						// handle error
						fmt.Println(err)
						os.Exit(2)
				}
				Withdraw(name,am)
				fmt.Println("")
			case "7":
				fmt.Print("7. Transfer. Please enter sender: ")
				scanner.Scan()
				sender := scanner.Text()
				fmt.Println("")
				fmt.Print("Receiver: ")
				scanner.Scan()
				receiver := scanner.Text()
	    	fmt.Println("")
				fmt.Print("Amount: ")
				scanner.Scan()
				amount := scanner.Text()
				am, err := strconv.ParseInt(amount,10,64)
				if err != nil {
						// handle error
						fmt.Println(err)
						os.Exit(2)
				}
				Transfer(sender,receiver,am)
				fmt.Println("")
			default:
				fmt.Println("Please type: 1,2,3,4,5,6,7")
			}
	}
}
