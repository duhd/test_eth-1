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
