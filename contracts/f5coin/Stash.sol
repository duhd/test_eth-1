/*
    @author: DuHD
    @version: 1.0
    @date: 09/04/2019

    To be deploy by Business, so that no human has owner access
*/
pragma solidity ^0.5.7;

import "./Owned.sol";


contract Stash is Owned {
    bytes32 walletCode;
    int balance;
    int8 stateStash; /*  0: Inactive, 1:ActiveA, 2:ActiveB, 3:Closed, 4:Blocked  */
    int8 typeState; /* 0:A, 1:V, 2:D, 3:E, 4:F */

    constructor(bytes32 _walletCode, int8 _typeState) public{
        walletCode = _walletCode;
        typeState = _typeState;
    }

    function credit(int _crAmt) onlyOwner public {
        balance += _crAmt;
    }

    function debit(int _dAmt) onlyOwner public {
        balance -= _dAmt;
    }

    function safe_debit(int _dAmt) onlyOwner public {
        require(_dAmt < balance);
        balance -= _dAmt;
    }

    function getState() view public returns (int8){
        return stateStash;
    }

    function setState(int8 _stateStash) onlyOwner public {
        stateStash = _stateStash;
    }

    function getType() view public returns (int8){
        return typeState;
    }

    function setType(int8 _typeState) onlyOwner public {
        typeState = _typeState;
    }

    function getBalance() view public returns (int){
        return balance;
    }

}