pragma solidity >=0.5.0 <0.6.0;

import "./SafeMath.sol";

contract MyWallet {
    using SafeMath for uint;

    address public owner;
    uint internal balance; // in wei

    constructor() public {
        owner = msg.sender;
    }

    event Deposited(address indexed _from, uint _weiAmount);
    event Withdrawn(address indexed _to, uint _weiAmount);
    event Transferred(address indexed _from, address indexed _to, uint _weiAmount);

    modifier onlyOwner {
        require(owner == msg.sender, "sender is not an owner");
        _;
    }

    /**
    * @dev Return the current balance.
    */
    function getBalance() public view returns (uint) {
        return balance;
    }

    /**
    * @dev Stores the sent amount in the wallet as credit to be withdrawn.
    * msg.value contains the amount in Wei to be stored.
    */
    function deposit() public payable {
        balance = balance.add(msg.value);
        emit Deposited(msg.sender, msg.value);
    }

    /**
    * @dev Withdraw to the owner the requested amount from the wallet.
    * @param weiAmount contains the amount in Wei to be withdrawn.
    */
    function withdraw(uint weiAmount) public onlyOwner {
        require(weiAmount <= balance, "insufficient funds");
        balance = balance.sub(weiAmount);
        emit Withdrawn(msg.sender, weiAmount);
        msg.sender.transfer(weiAmount);
    }

    /**
    * @dev Transfer the requested amount to another address.
    * @param recipient the address of the recipient.
    * @param weiAmount contains the amount in Wei to be transferred.
    */
    function transfer(address payable recipient, uint weiAmount) public onlyOwner {
        require(weiAmount <= balance, "insufficient funds");
        balance = balance.sub(weiAmount);
        emit Transferred(msg.sender, recipient, weiAmount);
        recipient.transfer(weiAmount);
    }
}
