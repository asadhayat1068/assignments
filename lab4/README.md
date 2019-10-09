![UiS](https://www.uis.no/getfile.php/13391907/Biblioteket/Logo%20og%20veiledninger/UiS_liggende_logo_liten.png)

# Exploring the blockchain

## Course Schedule

| Lab                                                              | Topic                            | Duration  |
| ---------------------------------------------------------------- | -------------------------------- | --------- |
| [1](https://github.com/dat650-2019/assignments/tree/master/lab1) | Data structures                  | 2 weeks   |
| [2](https://github.com/dat650-2019/assignments/tree/master/lab2) | UTXO                             | 2 weeks   |
| [3](https://github.com/dat650-2019/assignments/tree/master/lab3) | PoW                              | 2 weeks   |
| [4](https://github.com/dat650-2019/assignments/tree/master/lab4) | Intro to Smart Contracts         | 1 week    |
| [5](https://github.com/dat650-2019/assignments/tree/master/lab5) | Distributed Application          | 1 week    |
| Project                                                          | [Projects](#final-project-ideas) | 3-4 weeks |


## Environment setup
- CLI usage
- Docker?

## LAB 1
### Data structures
- Defining a block
- Defining a simple blockchain
- Cryptographic hashing
- Merkle Tree
- Basic command client
- Benchmark merkle tree

## LAB 2
### UTXO model
- Implementing transactions (utxo, send/receive, inspect)
- Cryptographic keys and Digital signatures
- Address generation (encoding)

## LAB 3
### Proving to commit
- Basic PoW
- Benchmark PoW
- Mining reward and Coinbase

## LAB 4
### Distributing applications
- Intro Smart Contracts (solidity)
- Basic running + testing (remix IDE, truffle + ganache, geth)

## LAB 5
- Building a DApp (contracts + client)

### Alternatives topics
- Pay to Public Key Hash ([P2PKH](https://en.bitcoin.it/wiki/Script)) - simple script language parsing ScriptPubKey string
- Bloom filters
- Smart contracts vulnerabilities [wargame](https://ethernaut.zeppelin.solutions/)

## Final Project Ideas
- Dapp implementation: Launching an ICO (Crowdsale + Token contract)
- Implement Paxos-variant Consensus - node sync (version and chain height)
- Implement a simple [proof-of-stake](https://blog.ethereum.org/2014/11/25/proof-stake-learned-love-weak-subjectivity/)
- Target adjusting algorithm (adjust puzzle difficult according with the number of miners)
- Full p2p and node discovery (using consensus or gossip) - nodes start mining and broadcast blocks -- attacks?
  - Adding persistence using [bbolt](https://github.com/etcd-io/bbolt)
    - Define two data types (as a key-value store) to persist data: blocks and chainstate like in [Bitcoin](https://en.bitcoin.it/wiki/Bitcoin_Core_0.11_(ch_2):_Data_Storage#Block_index_.28leveldb.29)
      1. blocks stores metadata describing all blocks in the chain
      2. chainstate stores the state of the chain, which is currently the unspent transactions output and some metadata.
    - Store the whole db in a single file
  - P2P Network ?(or a centralized option where the chainstate is broadcast to all nodes from a single server)
- Implement transaction fees

# References
- https://jeiwan.cc/
- https://learnblockcha.in/
- https://smartcontractsecurity.github.io/SWC-registry/
- https://github.com/ethereum/wiki/wiki/Problems
