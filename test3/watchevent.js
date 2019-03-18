const Web3 = require('web3');

// var web3 = new Web3("http://localhost:8502");

const web3 = new Web3("ws://localhost:8546");

// web3.eth.getAccounts().then(console.log);
// web3.eth.getGasPrice().then(console.log);
// web3.eth.getCoinbase().then(console.log);


var CoinSpawnABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"createCoin\",\"outputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"CreateCoinEvt\",\"type\":\"event\"}]"

var abi = JSON.parse(CoinSpawnABI);
var myContractInstance = new web3.eth.Contract(abi,'0x47562495779f98c048460514becf5e1a2d217c9f',{
  gasPrice: '20000000000'
});


myContractInstance.once('CreateCoinEvt', {
    fromBlock: 0
}, (error, event) => { console.log(event); });

//
// myContractInstance.getPastEvents('CreateCoinEvt', {
//     fromBlock: 0,
//     toBlock: 'latest'
// }, (error, events) => { console.log(events); })
// .then((events) => {
//     console.log(events) // same results as the optional callback above
// });
