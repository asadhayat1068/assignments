# Wallet contract
Build a simple wallet contract in Solidity that stores the owner's funds.
Your implementation need to fit the following rules:

## Rules
* There can only be one contract owner.
* Anyone should be able to make a deposit for the owner.
* Should be possible to anyone to retrieve the current balance stored in the contract.
* The contract should emit events of all performed operations except get the balance.
* Should be possible to the owner to withdraw any amount from the current contract's balance. But only the owner should be authorized to do so.
* The owner should be able to transfer some amount of the contract's balance to a given address.
* Should revert with the error message: `insufficient funds` if the owner attempt to withdraw or transfer more than the current contract's balance.
* You may add as many auxiliary functions as you want.
