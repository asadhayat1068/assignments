pragma solidity >=0.5.0 <0.6.0; 

import "../MyWallet.sol";

contract MyWalletMock is MyWallet {

    function setBalance() payable public {
        balance += msg.value;
    }
}