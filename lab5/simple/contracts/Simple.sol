pragma solidity ^0.5.11;

contract Simple {
    address public owner;
    string public value;
    uint public balance;

    constructor() public {
        owner = msg.sender;
        value = "initial value";
    }

    function get() public view returns (string memory) {
        return value;
    }

    function set(string memory newValue) public payable {
        require(msg.value >= balance);
        balance += msg.value;
        value = newValue;
    }

    function stop() public{
        require(msg.sender == owner);
        selfdestruct(msg.sender);
    }
} 