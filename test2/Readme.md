**********************************************************************************************
# Deploy F5 Contract

## 0. Config
File config thay doi. Please update new

## 1. Deploy contract & Test with manual
go run deploy_f5wallet.go
 go run f5wallet_manual.go

## 2. Run webserver, BockEvebt listen, ContractEvent Listen
go run fast_web_server.go
go run block_subscribe.go
go run event_log.go

## 3. List of Api
#### **************  ETH Account ***************
##### 1. Create new ETH account
http://http://localhost:8080/api/v2/wallet/new_account

##### 2. View list Eth Account in system
http://localhost:8080/api/v2/wallet/accounts

##### 3. View Balance of Eth Account
http://localhost:8080/api/v2/wallet/eth_balance/59f5545079762e130efaf873e7d28b2756fbc4b6

##### 2. View list Eth Account auto fill gas
http://localhost:8080/api/v2/wallet/autofill 

#### **************  Contract API ***************
##### 1. Register EthAccount
http://localhost:8080/api/v2/wallet/register

##### 2. Create Wallet
http://localhost:8080/api/v2/wallet/create/vi03/1

##### 3. Set Wallet state
http://localhost:8080/api/v2/wallet/set_state/vi03/1

##### 4. View wallet balance
http://localhost:8080/api/v2/wallet/balance/vi03

##### 5. Credit
http://localhost:8080/api/v2/wallet/deposit/vi03/vi03/1000

##### 6. Debit
http://localhost:8080/api/v2/wallet/withdraw/vi03/vi03/1000

##### 7. Transfer Money
http://localhost:8080/api/v2/wallet/transfer/tx01/VI01/vi03/100/Test/1

##### 8. Summary
http://localhost:8080/api/v2/wallet/summary

Running
**********************************************************************************************
### 1. Install go
sudo apt-get update
sudo apt-get -y upgrade
cd /tmp
wget https://dl.google.com/go/go1.11.linux-amd64.tar.gz
sudo tar -xvf go1.11.linux-amd64.tar.gz
sudo mv go /usr/local

export GOROOT=/usr/local/go
mkdir -p $HOME/go/src
export GOPATH=$HOME/go

go get -d https://github.com/binhdt101/test_eth.git
go get -d https://github.com/ethereum/go-ethereum

cd $HOME/go/src/test_eth/test2
### Preparing

+ go run add_peers.go
+ go run deploy_wallet.go
### Running webserver & listening server

+ go run web_server.go
+ go run block_subscribe.go
### Running client to test

### API:
1. List accounts:
http://localhost:8080/api/v1/wallet/accounts
2. View Blance of account
http://localhost:8080/api/v1/wallet/balance/ffbcd481c1330e180879b4d2b9b50642eea43c02
3. Transfer Token from ffbcd481c1330e180879b4d2b9b50642eea43c02  to a17a7a153c8d873a1df803c74e0664c13726f5e8 with mount of 2 and note of "Test"
http://localhost:8080/api/v1/wallet/transfer/ffbcd481c1330e180879b4d2b9b50642eea43c02/a17a7a153c8d873a1df803c74e0664c13726f5e8/2/Test
4. View report
http://localhost:8080/api/v1/wallet/report
