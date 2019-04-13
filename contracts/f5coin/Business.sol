/*
    @author: DuHD
    @version: 1.0
    @date: 09/04/2019
*/

pragma solidity ^0.5.7;

import "./Owned.sol";
import "./Stash.sol";


contract Business is Owned {

    constructor() public {
        owner = msg.sender;
    }

    // Bang luu giu address cua account, mapping giữa mã số ví và address contract  - accCode => contract address
    mapping(bytes32 => address) public stashRegistry;
    bytes32[] public stashNames;

    /////////////////////////////////////////////
    event event_createStash(bytes32 indexed wallet_code, address wallet_address, int8 wallet_type);
    /* Tao mot contract kieu Stash de luu data cua mot Account ví và map qua mã số ví  */
    function createStash(bytes32 _stashName, int8 _typeState) onlyMember public
    {
        require(!(stashRegistry[_stashName] > address(0x0)), 'Tai khoan da ton tai');
        address stash = address(new Stash(_stashName, _typeState));
        stashRegistry[_stashName] = stash;
        stashNames.push(_stashName);
        emit event_createStash(_stashName, stash, _typeState);
    }
    /////////////////////////////////////////////
    event event_reCreateStash(bytes32 indexed wallet_code, address old_wallet_address, address new_wallet_address, int8 wallet_type);
    /* Tao lai mot contract kieu Stash de luu data cua mot Account ví và map qua mã số ví  */
    function reCreateStash(bytes32 _stashName, int8 _typeState) onlyMember public
    {
        require(stashRegistry[_stashName] > address(0x0), 'Tai khoan chua ton tai');
        Stash oldStash = Stash(stashRegistry[_stashName]);
        require(oldStash.getState() == 3, 'Tai khoan muon tao lai phai o trang thai Closed');
        address stash = address(new Stash(_stashName, _typeState));
        stashRegistry[_stashName] = stash;
        stashNames.push(_stashName);
        emit event_reCreateStash(_stashName, address(oldStash), stash, _typeState);
    }
    /////////////////////////////////////////////
    /*getBalance*/
    function getBalance(bytes32 _stashName) onlyMember view public returns (int)
    {
        require(stashRegistry[_stashName] > address(0x0), 'Tai khoan khong ton tai');
        int bal = 0;
        Stash stash = Stash(stashRegistry[_stashName]);
        bal = stash.getBalance();
        return bal;
    }
    /////////////////////////////////////////////
    event event_setState(bytes32 indexed wallet_code, int8 stashState, int oldState);
    /* Stash state */
    function setState(bytes32 _stashName, int8 _stashState) onlyMember public
    isRegister(_stashName)
    {
        Stash stash = Stash(stashRegistry[_stashName]);
        int8 oldState = stash.getState();
        stash.setState(_stashState);
        emit event_setState(_stashName, _stashState, oldState);
    }
    /////////////////////////////////////////////
    function getState(bytes32 _stashName) onlyMember view public returns (int8)
    {
        int8 state = 0;
        require(stashRegistry[_stashName] > address(0x0), 'Tai khoan khong ton tai');
        Stash stash = Stash(stashRegistry[_stashName]);
        state = stash.getState();
        return state;

    }
    /////////////////////////////////////////////
    function getType(bytes32 _stashName) onlyMember view public returns (int8)
    {
        int8 _type = 0;
        require(stashRegistry[_stashName] > address(0x0), 'Tai khoan khong ton tai');
        Stash stash = Stash(stashRegistry[_stashName]);
        _type = stash.getType();
        return _type;

    }
    /////////////////////////////////////////////
    event event_debit(bytes32 indexed txRef, bytes32 indexed wallet_code, int amount, uint timestamp);
    /* debit */
    struct Debit {
        bytes32 txRef;
        bytes32 stashName;
        int amount;
        uint timestamp;
    }

    bytes32[] public debitIdx;
    mapping(bytes32 => Debit) public debits;

    function getDebitHistoryLength() view public returns (uint) {
        return debitIdx.length;
    }

    function debit(bytes32 _txRef, bytes32 _stashName, int _amount) onlyMember public
    isRegister(_stashName)
    {
        Stash stash = Stash(stashRegistry[_stashName]);
        require(stash.getState() == 1 || stash.getState() == 2, 'Yeu cau tai khoan o trang thai Active');
        stash.safe_debit(_amount);

        debitIdx.push(_txRef);
        debits[_txRef].txRef = _txRef;
        debits[_txRef].stashName = _stashName;
        debits[_txRef].amount = _amount;
        debits[_txRef].timestamp = now;
        emit event_debit(_txRef, _stashName, _amount, debits[_txRef].timestamp);
    }

    /////////////////////////////////////////////
    event event_credit(bytes32 indexed txRef, bytes32 indexed wallet_code, int amount, uint timestamp);
    /* credit */
    struct Credit {
        bytes32 txRef;
        bytes32 stashName;
        int amount;
        uint timestamp;
    }

    bytes32[] public creditIdx;
    mapping(bytes32 => Credit) public credits;

    function getCreditHistoryLength() view public returns (uint) {
        return creditIdx.length;
    }

    function credit(bytes32 _txRef, bytes32 _stashName, int _amount) onlyMember public
    isRegister(_stashName)
    {
        Stash stash = Stash(stashRegistry[_stashName]);
        require(stash.getState() == 1 || stash.getState() == 2, 'Yeu cau tai khoan o trang thai Active');
        stash.credit(_amount);

        creditIdx.push(_txRef);
        credits[_txRef].txRef = _txRef;
        credits[_txRef].stashName = _stashName;
        credits[_txRef].amount = _amount;
        credits[_txRef].timestamp = now;
        emit event_credit(_txRef, _stashName, _amount, credits[_txRef].timestamp);
    }

    /////////////////////////////////////////////
    event event_transfer(bytes32 indexed txRef, bytes32 indexed sender, bytes32 indexed receiver, int amount, string note, int8 txType, uint timestamp);
    /*transfer*/
    struct Transfer {
        bytes32 txRef;
        bytes32 sender;
        bytes32 receiver;
        int amount; // > 0
        string note;
        int8 txType;
        uint timestamp; // added to include sorting in API layer - Laks
    }

    bytes32[] public transferIdx;                  // @private (list of all-trans)
    mapping(bytes32 => Transfer) public transfers; // @private

    function getTransferHistoryLength() view public returns (uint) {
        return transferIdx.length;
    }

    function transfer(bytes32 _txRef, bytes32 _sender, bytes32 _receiver, int _amount, string memory _note, int8 _txType) onlyMember public returns (int sender_bal, int receiver_bal)
    {

        require(_amount >= 0, 'Yeu cau amount > 0');
        require(stashRegistry[_sender] > address(0x0), 'Tai khoan CHUYEN khong ton tai');
        require(stashRegistry[_receiver] > address(0x0), 'Tai khoan NHAN khong ton tai');

        Stash sender = Stash(stashRegistry[_sender]);
        require(sender.getState() == 1 || sender.getState() == 2, 'Yeu cau tai khoan CHUYEN o trang thai Active');
        require(sender.getBalance() >= _amount, 'So du khong du chuyen');
        require(sender.getType() == 1, 'Yeu cau tai khoan phai la loai V (1)');

        Stash receiver = Stash(stashRegistry[_receiver]);
        require(receiver.getState() == 1 || receiver.getState() == 2, 'Yeu cau tai khoan NHAN o trang thai Active');

        sender.safe_debit(_amount);
        receiver.credit(_amount);

        transferIdx.push(_txRef);
        transfers[_txRef].txRef = _txRef;
        transfers[_txRef].sender = _sender;
        transfers[_txRef].receiver = _receiver;
        transfers[_txRef].amount = _amount;
        transfers[_txRef].note = _note;
        transfers[_txRef].txType = _txType;
        transfers[_txRef].timestamp = now;

        emit event_transfer(_txRef, _sender, _receiver, _amount, _note, _txType, now);
        return (sender.getBalance(), receiver.getBalance());

    }

    /////////////////////////////////////////////
    event event_registerAccETH(address[] listAcc);

    function registerAccETH(address[] memory _listAcc) onlyOwner public {
        for (uint i = 0; i < _listAcc.length; i++) {
            registerMemberApi(_listAcc[i]);
        }
        emit event_registerAccETH(_listAcc);

    }
    /////////////////////////////////////////////
    function getRegistedAccEthLength() onlyOwner view public returns (int16) {
        return getMemberApiIdxLenght();
    }

    function getStashNamesLenght() view public returns (int) {
        return int(stashNames.length);

    }

    /////////////////////////////////////////////
    function changeOwnerAllStash(address _newOwner) onlyOwner public
    {
        for (uint i = 0; i < stashNames.length; i++) {
            Stash stash = Stash(stashRegistry[stashNames[i]]);
            stash.changeOwner(_newOwner);
        }
    }
    /////////////////////////////////////////////
    function loadStashRegistry(bytes32 _stashName, address _stash) onlyOwner public
    {
        stashRegistry[_stashName] = _stash;
        stashNames.push(_stashName);
    }
    /////////////////////////////////////////////


    modifier isPositive(int _amount) {require(_amount >= 0, 'Yeu cau amount > 0');
        _;}
    modifier isActive(Stash _stash) {require(_stash.getState() == 1 || _stash.getState() == 2, 'Yeu cau tai khoan o trang thai Active');
        _;}
    modifier isRegister(bytes32 _stashName) {require(stashRegistry[_stashName] > address(0x0), 'Tai khoan khong ton tai');
        _;}

    modifier isTypeV(Stash _stash) {require(_stash.getType() == 1, 'Yeu cau tai khoan phai la loai V (1)');
        _;}


    function string_tobytes(string memory s) pure public returns (bytes memory){
        bytes memory b3 = bytes(s);
        return b3;
    }

}