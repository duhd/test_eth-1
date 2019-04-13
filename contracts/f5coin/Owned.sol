/*
    @author: DuHD
    @version: 1.0
    @date: 09/04/2019
*/

pragma solidity ^0.5.7;

contract Owned {
    address owner;
    mapping(address => bool)  private MemberApi;
    address[] private memberApiIdx;

    constructor() public{
        owner = msg.sender;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, 'Chi owner contract moi duoc goi ham');
        _;
    }

    function getOwner() view public returns (address) {
        return owner;
    }

    function changeOwner(address _newOwner) onlyOwner public {
        owner = _newOwner;
    }

    function registerMemberApi(address _newMember) onlyOwner public {
        MemberApi[_newMember] = true;
        memberApiIdx.push(_newMember);
    }

    function getMemberApiIdxLenght() view public returns (int16)
    {
        return int16(memberApiIdx.length);
    }

    modifier onlyMember() {
        require(MemberApi[msg.sender], 'Chi cac acc ETH da dang ky moi goi duoc ham');
        _;
    }
}