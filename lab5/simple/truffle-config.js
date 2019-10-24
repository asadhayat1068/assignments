// https://github.com/chaijs/chai/pull/868
// Using Should style globally
require('chai/register-should');
const path = require('path');

module.exports = {
  // contracts_build_directory: path.join(__dirname, "client/src/contracts"),
  // Uncommenting the defaults below 
  // provides for an easier quick-start with Ganache.
  // You can also follow this format for other networks;
  // see <http://truffleframework.com/docs/advanced/configuration>
  // for more details on how to specify configuration options!
  networks: {
    development: { // local test net
      host: "127.0.0.1",
      port: 8545,
      network_id: "*"
    },
    ganache: { // ganache-cli
      host: "127.0.0.1",
      port: 7555,
      network_id: "5777"
    },
  },
  compilers: {
    solc: {
      version: '0.5.11'
      // settings: {
      //  optimizer: {
      //    enabled: false,
      //    runs: 200
      //  },
      //  evmVersion: "byzantium"
      // }
    }
  }
};