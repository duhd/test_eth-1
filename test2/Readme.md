Running

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
