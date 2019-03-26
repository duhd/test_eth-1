Running

#1. Install go
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
