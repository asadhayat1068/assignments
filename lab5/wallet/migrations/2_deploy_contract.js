const MyWallet = artifacts.require('MyWallet');
module.exports = function (deployer) {
  deployer.deploy(MyWallet);
};