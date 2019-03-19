abigen -sol inbox.sol -pkg contracts -out inbox.go

0xbb7dcbfb547171d1d619ba4f83fde03da3cbc5bc

abigen --abi PathToCivilTCRContract.abi --pkg main --type CivilTCR --out CivilTCRContract.go

cd /Users/nguyenthanhbinh/Blockchain/test/ethereum/my-eth-chain-poa
geth --datadir datadir account list

 geth --datadir datadir account new

Passphrase: 123456



git clone https://github.com/ethereum/cpp-ethereum.git
mkdir cpp-ethereum/build
cd cpp-ethereum/build
cmake -DJSONRPC=OFF -DMINER=OFF -DETHKEY=OFF -DSERPENT=OFF -DGUI=OFF -DTESTS=OFF -DJSCONSOLE=OFF ..
make -j4
make install
which solc

brew install solidity
which solc
admin.setSolc("/usr/local/bin/solc")


node1: 30302 8502
node2: 30304 8504
node3: 30306 8503

clique.getSigners()

miner.setEtherbase(eth.accounts[0])
personal.unlockAccount(eth.coinbase)
miner.setEtherbase(eth.coinbase)
miner.start(30)

eth.accounts
eth.coinbase

net.listening
net.peerCount

admin.addPeer
admin.nodeInfo
admin.nodeInfo.enode


admin.addPeer("
enode://27620b1ef2208de5509c30df0aba4e087e3508071acf5df9143f69d24c8db4cf1715d6518e4f45cd452b3f692592b1af8622e14f4a2efa08dbcf1563be28482d@127.0.0.1:30302

admin.peers

eth.accounts

clique.propose(<NEWSEALER>, true)
clique.getSigners()
clique.proposals
clique.propose("0x500b47e4262b65565d7fccdb7c9e20f1e721407b", false)
clique.propose("0x2554d20e9f437301151874d8fff6e439f10139f3", true)

0x500b47e4262b65565d7fccdb7c9e20f1e721407b

personal.unlockAccount("0xeb80964e1567064ba810b45300fd2ce3193d1684", "123456")

geth attach --datadir datadir  ipc://Users/nguyenthanhbinh/Blockchain/test/ethereum/my-eth-chain-poa/datadir/geth.ipc --rpccorsdomain "http://localhost:8000"

bootnode -genkey boot.key
bootnode -nodekey boot.key -verbosity 9 -addr :30310

 bootnode -nodekey boot.key  -addr :30310 -v5  -writeaddress

 bootnode -nodekey boot.key
 
miner.stop()


//Deploy vndwallet

0xa1d531a20dd8ae58b030324ab071a0f04557fe3e
0x4ba018d94eb3e9ba3cc34a0c8e4296db974a680a

brew update
brew tap ethereum/ethereum
brew install solidity

solc --abi Store.sol -o build
abigen --abi=./build/Store.abi --pkg=store --out=Store.go




admin.addPeer("enode://7939738f7643190ab4f7e6595644f178cf241581edf49f9a7a8c33f73cd1111e3553aa73bd782af06eb0d8faa85d7ccf3faab4403d4fbe66443e365be71843ab@172.101.0.18:30302")
