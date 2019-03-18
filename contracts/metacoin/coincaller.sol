pragma solidity >=0.4.21 <0.6.0;

import './coinspawn.sol';

contract coinCaller{
  struct transfer{
  		metaCoin coinContract;
  		uint amount;
  		address recipient;
      bool successful;
      uint balance;
	}
	mapping(uint => transfer) transfers;
	uint numTransfers;

  event sendCoinEvt(address from, address to, uint amount, bool txstatus, uint balance);

	function sendCoin(address coinContractAddress, address receiver, uint amount) public {
      transfer storage t = transfers[numTransfers]; //Creates a reference t
  		t.coinContract = metaCoin(coinContractAddress);
  		t.amount = amount;
  		t.recipient = receiver;
  		t.successful = t.coinContract.sendToken(receiver, amount);
  		t.balance = t.coinContract.balances(tx.origin);
  		numTransfers++;
      emit sendCoinEvt(tx.origin,receiver,amount,t.successful,t.balance);
	}
}
