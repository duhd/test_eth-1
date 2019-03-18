pragma solidity >=0.4.21 <0.6.0;
import './metacoin.sol';

contract coinSpawn{
    mapping(uint => metaCoin) public deployedContracts;
  	uint numContracts = 0 ;

    event CreateCoinEvt(address addr, uint initialBalance);

    function createCoin(uint initialBalance) public returns(address a){
      metaCoin mC = new metaCoin(initialBalance);
  		deployedContracts[numContracts] = mC;
      emit CreateCoinEvt(address(mC),initialBalance);
  		numContracts++;
  		return address(mC);
  	}
}
