pragma solidity >=0.4.21 <0.6.0;

contract Trigger {
  address owner;

  constructor() public {
      owner = msg.sender;
  }
  event TriggerEvt(address _sender, uint _trigger);

  function trigger(uint _trigger) public {
    emit  TriggerEvt(msg.sender, _trigger);
  }
  function getOwner() external returns  (address)  {
    return owner;
  }
}
