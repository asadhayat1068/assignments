pragma solidity >=0.5.0 <0.6.0;

import "truffle/Assert.sol";
import "truffle/DeployedAddresses.sol";
import "../contracts/MyWallet.sol";

contract TestMyWallet {
  MyWallet wallet;
  // Truffle will send the TestWallet one Ether after deploying the contract.
  uint256 public initialBalance = 10 ether;

  function () external payable {}

  function beforeEach() public {
    wallet = new MyWallet();
  }

  function testInitialBalance() public {
    Assert.equal(address(this).balance, 10 ether, "The contract TestWallet should have a starting balance of 10 ether");
    Assert.equal(wallet.getBalance(), 0, "A new Wallet contract has a zero balance");
  }

  function testSettingAnOwnerDuringCreation() public {
    // The TestMyWallet contract is the deployer
    wallet = new MyWallet(); //create a new contract
    Assert.equal(wallet.owner(), address(this), "The owner shouldn't be different than the deployer");
  }

  function testSettingAnOwnerUsingDeployedContract() public {
    // The msg.sender(first address in ganache) is the new contract creator
    wallet = MyWallet(DeployedAddresses.MyWallet());
    Assert.equal(wallet.owner(), msg.sender, "The owner shouldn't be different than the deployer");
  }

  function testValidDeposit() public {
    // Send 'deposit' 1000 wei (it's being sent by 'TestWallet').
    // Wei is the default denomination in Solidity if you don't 
    // specify another unit 
    uint initBalance = address(this).balance;

    wallet.deposit.value(1000)();

    Assert.equal(address(this).balance, initBalance - 1000 wei, "Current balance of the sender doesn't correspond to the amount deposited");
    Assert.equal(wallet.getBalance(), 1000 wei, "The owner balance is different than the deposited amount");
  }

  function testAccumulatedDeposits() public {
    wallet.deposit.value(3 finney)();
    wallet.deposit.value(15 finney)();
    Assert.equal(wallet.getBalance(), 18 finney, "Owner balance is different than sum of the deposits");
  }

  function testWithdrawalAttemptWithNoBalance() public {
    Assert.equal(wallet.getBalance(), 0, "Balance should be 0 initially");

    (bool r, ) = address(wallet).call(abi.encodeWithSelector(wallet.withdraw.selector, 1 ether));

    Assert.isFalse(r, "Should revert due to a withdrawal attempt without funds");
  }

  function testWithdrawalByAnOwner() public {
    uint initBalance = address(this).balance;
  
    wallet.deposit.value(100 finney)();
    Assert.equal(address(this).balance, initBalance - 100 finney, "Balance of the sender before the withdrawal isn't correct");

    Assert.equal(wallet.getBalance(), 100 finney, "Contract balance is different than the value deposited");

    (bool r, ) = address(wallet).call(abi.encodeWithSelector(wallet.withdraw.selector, 100 finney));

    Assert.isTrue(r, "Should successfully withdrawal ether from the wallet contract");
    Assert.equal(address(this).balance, initBalance, "Balance after withdrawal should be equal to the initial balance");
  }

  function testTransferByAnOwner() public {
    uint initBalance = address(this).balance;

    address payable _to = 0xdCad3a6d3569DF655070DEd06cb7A1b2Ccd1D3AF;
    uint toInitBalance = _to.balance;

    wallet.deposit.value(100 finney)();
    Assert.equal(address(this).balance, initBalance - 100 finney, "Balance of the sender before the transfer isn't correct");

    Assert.equal(wallet.getBalance(), 100 finney, "Contract balance is different than the value deposited");

    (bool r, ) = address(wallet).call(abi.encodeWithSelector(wallet.transfer.selector, _to, 100 finney));

    Assert.isTrue(r, "Should successfully transfer ether from the wallet contract to the given account");

    Assert.equal(wallet.getBalance(), 0, "Balance at contract after the transfer should be zero");

    Assert.equal(_to.balance, toInitBalance + 100 finney, "The transfered value isn't the expected");
  }
}