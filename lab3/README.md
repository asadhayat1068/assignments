![UiS](https://www.uis.no/getfile.php/13391907/Biblioteket/Logo%20og%20veiledninger/UiS_liggende_logo_liten.png)

# Lab 3: Proof-Of-Work

| Lab 3:           | PoW                          |
| ---------------- | ---------------------------- |
| Subject:         | DAT650 Blockchain Technology |
| Deadline:        | 03. OCT                      |
| Expected effort: | 1 week                       |
| Grading:         | Pass/fail                    |

## Table of Contents
- [Lab 3: Proof-Of-Work](#lab-3-proof-of-work)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Proof-Of-Work](#proof-of-work)
  - [Demo Application](#demo-application)
  - [Lab Approval](#lab-approval)

## Introduction

In our current blockchain implementation adding new blocks is easy and fast, but in real blockchain adding new blocks requires some work: one has to perform some heavy computations before getting permission to add block (this mechanism is called Proof-Of-Work).

## Proof-Of-Work

One of the key innovations of Bitcoin was to use a Proof-Of-Work algorithm to conduct a global "election" every 10 minutes (adjusted by the target difficulty), allowing the decentralized network to achieve consensus about the state of transactions.
The "elected" peer, i.e., the one that solves PoW for a block and publishes it to a majority of the peers in the network faster than the others, has granted the permission to write data on the blockchain, adding his published block to it.
And thus, consistently replicating the state in a decentralized network.

The miner that accomplish this task following the protocol receives a reward for his hard work to create a block and solve the PoW puzzle.
This is how new coins are introduced in the ecosystem.
The reward is the coinbase transaction that is added by the miner as the first transaction in the block, and there is only one coinbase transaction per block.
When a peer receives a new block, it will validate the block by checking, among other things, the block hash resulted from the PoW, the block timestamp, the block size, if and only if the first transaction in the block is a coinbase transaction, etc.

The miner will only "collect" its reward if his published block ends up on the longest chain because just like every other transaction, the coinbase transaction will only be accepted by other peers and included in the longest chain if a majority of miners decide to do so.
That's the key idea behind the Bitcoin incentive mechanism.
If most of the network is following the longest valid chain rule, all other peers are incentivized to continue to follow that rule.

In this lab, we will implement the Proof-Of-Work algorithm similar to the one used in Bitcoin.
Bitcoin uses [Hashcash](https://en.wikipedia.org/wiki/Hashcash), a Proof-of-Work algorithm that was initially developed to prevent email spam.
The goal of such work is to find a hash for a block, that should be computationally hard to produce.
And it's this hash that serves as a proof which should be easy for others to verify its validity.
Thus, finding a proof is the actual work.

In order to create a block, the peer that proposes that block is required to find a number, or _​nonce​_, such that when concatenating the header fields of the block with the _nonce_ and producing a hash of it, this hash output should be a number that falls into a target space that is quite small in relation to the much larger output space of that hash function.
We can define such a target space as any value falling below a certain target value, this is the _difficulty of mining_ a block.

Thus, this is a brute force algorithm: you change the _nonce_, calculate a new hash, check if it is smaller then the _target_, if not, increment the _nonce_, calculate the new hash, etc.
That's why the PoW is computationally expensive.

In the original Hashcash implementation, the target difficulty sounds like "find the hash in which the first 20 bits are zeros".
In Bitcoin, the "target bits" is added to the block header to inform the difficulty at which the block was mined, and this difficulty is adjusted from time to time, because, by design, a block must be generated every 10 minutes, despite computation power possibly increasing over the time and the number of miners in the network.

As we will not implement an adjustable mining difficulty now, we will just define the difficulty as a global constant `TARGETBITS` in the file `proof-of-work.go`.

We will set the default value of `TARGETBITS` to 8 to make the block creation fast on the tests, our goal is to have a _target_ that takes less than 256 bits in memory.
So our target will be calculated using the following formula: `2^(256-TARGETBITS)`.
The bigger the `TARGETBITS` the smaller will be the `target` number, and consequently, it's more difficult to find a proper hash.
Think of a _target_ as the upper boundary of a range: if a number (a hash) is lower than the boundary, it's valid, and vice versa.
Lowering the boundary will result in fewer valid numbers, and thus, more difficult will be the work required to find a valid one.

With the `TARGETBITS` equals to 8, you will be required to find a block hash in which the number representation is less than 2^248. Or in other words, a hash in which the first 8 bits are zeros.

For example, suppose that in the first iteration, after hash a block using the nonce value of 1, you obtained the hash `73d40a0510b6327d0fbcd4a2baf6e7a70f2de174ad2c84538a7b09320e9db3f2`, which converted to big integer representation is equals to `52390618318831801638175855856716822591931229920359547228571203746793472766962`.
As mentioned before, the default target difficulty is `2^248 == 452312848583266388373324160190187140051835877600158453279131187530910662656`.
Thus, the hash above isn't a valid PoW solution, since it is bigger than the target, i.e., `52390618318831801638175855856716822591931229920359547228571203746793472766962 > 452312848583266388373324160190187140051835877600158453279131187530910662656`. 

However, if you continue hashing the same header data just incrementing the nonce, let's say until the nonce value 59, you can eventually find a hash that is smaller than the target, like the hash `00d4eeaee903dce5468d4c6975376dfbc4c45ea1bc6c5bbbfd8e13b26aaf6e3b`, which can be represented by the big integer number `376218908933626769012171312496768664868580826658885427967934344392923377211` and is a valid solution.

Ok, so let's do the work!
Your task is to implement all functions marked with `TODO(student)` on the lab code templates.
A small description of what each function should do is given below:

- `ProofOfWork.NewProofOfWork`: Creates a new proof of work containing the given block and calculates the target difficulty based on the `TARGETBITS`. The _target_ should be a [big number](https://golang.org/pkg/math/big/).
- `ProofOfWork.setupHeader`: Prepares the header of the block by concatenating the `block.PrevBlockHash`, Merkle root hash of the `block.Transactions`, `block.Timestamp` and the `TARGETBITS` in this order.
- `ProofOfWork.addNonce`: Adds a nonce to the prepared header.
- `ProofOfWork.Run`: Performs the Proof-Of-Work algorithm. You should make a brute-force loop incrementing the nonce and hashing it with the prepared header using the [SHA256](https://golang.org/pkg/crypto/sha256/) until you find a valid hash or you reach the defined `maxNonce`. Remember that a valid PoW hash is the one that is smaller than the _target_, so you need to be able to compare them, converting the hash bytes to a big number.
- `ProofOfWork.Validate`: Validates a block's Proof-Of-Work. Until now, this function just validates if the block header hash is smaller than the target, and it ignores validation errors like if the block timestamp is in the future.
  
- `Block.Mine`: Replace the function `SetHash()` for a new one called Mine that create a ProofOfWork and sets the block hash and nonce based on the result of the `Run()`.
- `Block.NewBlock`: Modify to use `Mine()` instead of `SetHash()` method.
  
- `Blockchain.ValidateBlock`: Validates a block after mining or before adding it to the blockchain (in case of receiving it from another peer). Currently, this function should perform the following validations:
    1. Check if, and only if, the first transaction in the block is coinbase.
    2. Check if the Proof-Of-Work for that block is valid.

## Demo Application

Create two addresses and start a transaction sending some coins between them like in the previous lab.
Store the transaction in a buffer that will be used to create a block.

Extend your client application to have a command to start mining and display the hash of each mining attempt until finding a solution. Show the hash in the hexadecimal format.
You need to be able to mine the block containing the transaction that you made.

When you start mining, your client application must add a new coinbase transaction giving the miner address (e.g., can be the first address that you create) and create a block with this transaction as the first transaction in the block. Also include the transaction that you created initially, in the buffer, and use both transactions to mine the new block. This will confirm your transactions and transfer the coins.

You will be requested to display the block and transactions information.

## Lab Approval

To have your lab assignment approved, you must come to the lab during lab hours and present your solution. This lets you present the thought process behind your solution, and allows us to provide feedback on your solution then and there.
When you are ready to show your solution, reach out to a member of the teaching staff. It is expected that you can explain your code and show how it works. You may show your solution on a lab workstation or your own computer.

You should for this lab present a working demo of the application described in the previous section making a transaction between two addresses. 
You should demonstrate that your implementation fulfills the previously listed specification of each assignments part.
The task will be verified by a member of the teaching staff during lab hours.