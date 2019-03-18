pragma solidity ^0.5.4;

contract Counter {

  uint public count = 0;

  event CounterDeployEvt(address owner);
  event CounterIncreasedEvt(address owner, uint count);

  function increase() public {
      count = count + 1;
      emit  CounterIncreasedEvt(msg.sender, count);
  }
}
