pragma solidity >=0.5.0 <0.6.0;

contract MyWallet {
    address public owner;
    uint256 internal balance; // in wei

    constructor() public {/* TODO (students) */}

    /* TODO (students): Take a look at the test to figure out which fields your events need to have. */
    event Deposited();
    event Withdrawn();
    event Transferred();

    modifier onlyOwner {
        /* TODO (students) */
        _;
    }

    /**
    * @dev Return the current balance.
    */
    function getBalance() public view returns (uint256) {
        /* TODO (students) */
    }

    /**
    * @dev Stores the sent amount in the wallet as credit to be withdrawn.
    * msg.value contains the amount in Wei to be stored.
    */
    function deposit() public payable {/* TODO (students) */}

    /**
    * @dev Withdraw to the owner the requested amount from the wallet.
    * @param weiAmount contains the amount in Wei to be withdrawn.
    */
    function withdraw(uint256 weiAmount) public {/* TODO (students) */}

    /**
    * @dev Transfer the requested amount to another address.
    * @param recipient the address of the recipient.
    * @param weiAmount contains the amount in Wei to be transferred.
    */
    function transfer(address payable recipient, uint256 weiAmount) public {/* TODO (students) */}
}
