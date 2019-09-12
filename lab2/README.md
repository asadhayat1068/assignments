![UiS](https://www.uis.no/getfile.php/13391907/Biblioteket/Logo%20og%20veiledninger/UiS_liggende_logo_liten.png)

# Lab 2: UTXO

| Lab 2:                | UTXO                         |
| --------------------  | ---------------------------- |
| Subject:              | DAT650 Blockchain Technology |
| Deadline:             | 12. SEP                      |
| Expected effort:      | 2 weeks                      |
| Grading:              | Pass/fail                    |

## Table of Contents

- [Lab 2: UTXO](#lab-2-utxo)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Part 1](#part-1)
    - [UTXO](#utxo)
    - [Transaction Outputs](#transaction-outputs)
    - [Transaction Inputs](#transaction-inputs)
    - [The egg](#the-egg)
    - [Unspent Transaction Output Set](#unspent-transaction-output-set)
  - [Demo Application](#demo-application)
  - [Lab Approval](#lab-approval)

## Introduction

We built a very simple blockchain prototype: it's just an array of blocks, with each block having a connection to the previous one.
The actual blockchain is much more complex though.
The purpose of the blockchain is to store transactions in a secure and reliable way in a distributed environment, so no one could modify them after they are created.
Until now our transaction implementation just store all the data in a byte array and don't contain any information that could help to identify the sender or the receiver of each transactions.

## Part 1

If you've ever developed an application that manage payments, you would likely to create these tables in a database: `accounts` and `transactions`.
An account would store information about a user, including their personal information and balance, and a transaction would store information about money transferring from one account to another.
In Bitcoin, payments are realized in completely different way, there are no accounts or balances fields, or even tables.

Since blockchain is a public and open database, we don't want to store sensitive information about wallet owners.
Coins are not collected in accounts.
Transactions do not transfer money from one address to another.
There's no field or attribute that holds account balance.
There are only transactions.

### UTXO

A transaction is a combination of inputs and outputs, in our implementation we will change the previous definition of a transaction in the `transaction.go` file to the struct below:

```go
type Transaction struct {
  ID   []byte
  Vin  []TXInput
  Vout []TXOutput
}
```

Inputs of a new transaction reference outputs of a previous transaction (there's an exception though, which we'll discuss later).
Outputs are where coins are actually stored.
An output has an implied index number based on its location in the `TXOutput` list (the index of the first output is zero).
The [Figure 3](#btctransactions) illustrate the interconnection between transactions inputs and outputs.

![Bitcoin transaction propagation][btctransactions]

_(image extracted from https://bitcoin.org/en/blockchain-guide#introduction)_

Is important to notice that:

1. There are outputs that are not linked to inputs.
2. In one transaction, inputs can reference outputs from multiple transactions.
3. An input must reference an output.

In Bitcoin transactions just lock values with a script, which can be unlocked only by the one who locked them.

### Transaction Outputs

Actually, it's outputs that store _coins_ (notice the Value field on the struct `TXOutput` below).
And storing means locking them with a "puzzle", which is stored in the `ScriptPubKey` field.
Internally, Bitcoin uses a scripting language called [Script](https://en.bitcoin.it/wiki/Script), that is used to define outputs locking and unlocking logic.
The language is quite primitive and not [Turing-Complete](https://en.wikipedia.org/wiki/Turing_completeness) (this is made intentionally, to avoid possible hacks and misuses), but we won't implement it in this course.

```go
type TXOutput struct {
	Value        int
	ScriptPubKey string
}
```

In Bitcoin, the `Value` field stores the number of [Satoshis](https://en.bitcoin.it/wiki/Satoshi_(unit)) which it pays to a conditional pubkey script. Anyone who can satisfy the conditions of that `ScriptPubKey` can spend up to the amount of satoshis paid to it.
A satoshi is a hundred millionth of a bitcoin (0.00000001 BTC), thus this is the smallest unit of currency in Bitcoin (like a cent).

Since we don't have addresses implemented yet, we'll avoid the whole scripting related logic for now.
`ScriptPubKey` will store an arbitrary string (e.g., user defined wallet address/identifier, like a user name).
Having such scripting language means that Bitcoin can be used as a [contract](https://en.bitcoin.it/wiki/Contract) platform as well, but with much more limited functionalities than in the [Ethereum platform](https://github.com/ethereumbook/ethereumbook/blob/develop/07smart-contracts-solidity.asciidoc#what-is-a-smart-contract).

One important thing about outputs is that they are _indivisible_, which means that you cannot reference a part of its value.
When an output is referenced in a new transaction, it's spent as a whole.
And if its value is greater than required, a change is generated and sent back to the sender.
This is similar to a real world situation when you pay, let's say, $10 for something that costs $8 and get a change of $2.

### Transaction Inputs

As mentioned earlier, an input references a previous output: `Txid` stores the ID of such transaction, and `OutIdx` stores an index to identify a particular output to be spent in that transaction.
`ScriptSig` is a script which provides data to be used in an output's `ScriptPubKey`.
If the data is correct, the output can be unlocked, and its value can be used to generate new outputs; if it's not correct, the output cannot be referenced in the input.
This is the mechanism that guarantees that users cannot spend coins belonging to other people.

```go
type TXInput struct {
  Txid      []byte
  OutIdx    int
  ScriptSig string
}
```

Again, since we don't have addresses implemented yet, `ScriptSig` will store just an arbitrary user defined wallet address (e.g., user name).
We'll implement public keys and signatures checking in the [part 2](#part-2) of this lab.

So, basically, _Outputs_ are where _coins_ are stored.
Each output comes with an unlocking **script**, which determines the logic of **unlocking** the output.
Every new transaction **must** have at least one input and output.
An input **references** an output from a previous transaction and provides data (the `ScriptSig` field) that is used in the output's unlocking script to unlock it and use its value to create new outputs.

But what came first: inputs or outputs?

### The egg

In Bitcoin, it's the egg that came before the chicken.
The _inputs-referencing-outputs_ logic is the classical "chicken or the egg" situation: inputs produce outputs and outputs make inputs possible.
And in Bitcoin, **outputs come before inputs**.

When a miner starts mining a block (you will see how mine works on the [Lab 3](../lab3/README.md)), it adds a _coinbase_ transaction to it.
A _coinbase_ transaction is a special type of transactions, which doesn't require previously existing outputs, and is the first transaction in each block.
It creates outputs (i.e., "coins") out of nowhere.
The egg without a chicken.
This is the **reward** miners get for mining new blocks.
 
As you know, there's the genesis block in the beginning of a blockchain.
It's this block that generates the very first output in the blockchain.
And no previous outputs are required since there are no previous transactions and no such outputs.

A _coinbase_ transaction has only one input. In our implementation its `Txid` should be empty and OutIdx equals to `-1`.
Also, a _coinbase_ transaction doesn’t store a script in ScriptSig. Instead, arbitrary data is stored there.
Take a look [here](https://www.blockchain.com/btc/tx/4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b?show_adv=true) to see the message stored on the first coinbase transaction of Bitcoin blockchain.

We define our default coinbase message for the genesis block in the file `config.go`.
In the config file we also define the constant `BlockReward`, that represents the amount of reward a miner will earn when adding a new block to the chain.
In Bitcoin, this number is not stored anywhere and is calculated based on the total number of blocks and any transaction fees paid by transactions included in this block.
Mining the Bitcoin genesis block produced 50 BTC in 2009, and every 210000 blocks the reward is halved (approximately every four years).
In our implementation, we’ll store the reward as a constant (at least for now).

### Unspent Transaction Output Set

An Unspent Transaction Output (UTXO) is and output that can be spent as an input in a new transaction.
Unspent means that these outputs weren’t referenced in any inputs.
A transaction consumes previously recorded unspent transaction outputs and creates new transaction outputs that can be consumed by a future transaction.
This way, chunks of bitcoin value move forward from owner to owner in a chain of transactions consuming and creating UTXO, as showing the [Figure 4](#utxo).

![UTXO model][utxo]

The ownership of some output is given by the ability of the user to provide an input in which the `ScriptSig` that can unlock the referred output's `ScriptPubKey`, and consequently an authorization to spend it.
The lock mechanism will be implemented by the methods `CanBeUnlockedWith` and `CanUnlockOutputWith` defined in the `transaction_output.go` and `transaction_input.go` respectively.
For now, these functions should just compare the script fields with `unlockingData`, which is a string.
These functions will be improved in the next part of this lab, after we implement addresses based on private keys.

But how we can find all the transactions containing unspent outputs?

The simple solution is just check every block in the blockchain.
Starting by searching and collecting for outputs that can be unlocked by the address we’re searching unspent transaction outputs for.
But before taking those outputs, we need to check if it was previously referenced in an input (e.g., was already spent).

If an output were referenced in some input (i.e., their values were moved to other outputs), thus we cannot count them and we should ignore it.
Thus, after checking the desired outputs inside the transaction, we should gather all inputs that could unlock outputs locked with the provided address, since this will be used to check if some output was already spent.
But remember that this doesn’t apply to coinbase transactions, since they don’t unlock outputs.
Then this process continues for every transaction in the block, for all blocks, resulting in a list of transactions containing only unspent outputs (i.e. UTXO).

But as you probably imagined, since transactions are stored in blocks, iterate over each block in the blockchain database and checks every transaction in it can be very inefficient and time-consuming.
Currently the whole blockchain database takes around 230 Gb of disk space, with more than 592 blocks.

The solution to the problem is to have an index that stores only unspent outputs, which is exactly what Bitcoin full nodes do, by tracking all available and spendable outputs.
The Bitcoin blockchain stores data in different LevelDB [databases](https://en.bitcoin.it/wiki/Bitcoin_Core_0.11_(ch_2):_Data_Storage#Overview). The blocks are stored in the `blocks` database, and instead of store all the transactions, it stores what is called the `UTXO set` in the `chainstate` database to easily retrieve unspent transactions outputs.

The `UTXO set` is the collection of all unspent transaction outputs (UTXO) and it is used as a cache that is built from all blockchain transactions, and later used to calculate balance and validate new transactions.
The UTXO set grows as new UTXO is created and shrinks when UTXO is consumed.
Every transaction represents a change (state transition) in the UTXO set.

In this lab, as we are not implementing persistence, we will represent the UTXO Set as a map, where the keys are the transaction ID and the content is a slice of transactions Outputs into that transaction.

```go
type UTXOSet map[string][]TXOutput
```

But to check the balance of some user, we don’t need all unspent transactions output (i.e., UTXO set), but only those that can be unlocked by the known key.
But we can use the UTXO set to search for those UTXO, which is faster than search on the whole blockchain.
When the user's wallet application "receive" bitcoin, this actually means that the wallet has detected an UTXO that can be spent with one of the keys controlled by that wallet (i.e., unlocked by the user).
The account balance is the sum of values of all UTXO locked by the account address (i.e., that the user's wallet can spend).

Your task is implement all functions marked with `TODO(student)` on the lab code templates.
A small description of what each function should do is given below:

- `Block.HashTransactions`: Modify this method to use the result of the `Transaction.Serialize` method instead of just the `Data` field used in the previous lab, since now your transaction is more than just a byte slice.
- `Transaction.NewCoinbaseTX`: creates a new _coinbase_ transaction.
- `Transaction.NewUTXOTransaction`: creates a new transaction by getting a list of spendable outputs to be used and the current balance of the sender. It also checks if the sender sufficient funds to perform the transaction.
- `Transaction.IsCoinbase`: test if a transaction is _coinbase_.
- `Transaction.Hash`: hash the serialized copy of a transaction ignoring the ID field, since the return of this function is the new ID of a transaction. **Before** serialize the transaction, make sure to ignore any existent data in the ID field (i.e., set ID in the copy to []byte{}).
- `Transaction.Serialize`: encodes the transaction struct using the [gob](https://golang.org/pkg/encoding/gob/) library. This method should not ignore the ID of the transaction, the whole transaction fields should be serialized.
- `Blockchain.FindUTXOSet`: finds all unspent outputs by iterating over all blocks. Returns an UTXO set.
- `Blockchain.FindTransaction`: finds a transaction in the blockchain by its ID. It iterates over all blocks until finds it.
- `Blockchain.VerifyTransaction`: verifies if the referred inputs of a given transaction exist in the blockchain, discard invalid transactions that make reference to unknown inputs.
- `Blockchain.MineBlock`: creates a new block on the blockchain. Besides the name, this function will not really "mine" a block yet, we will implement it on the next lab. But for now, this function should verify the existence of transactions inputs and add a new block if there is a list of valid transactions.

- `UTXOSet.FindSpendableOutputs`: this function is used when a new transaction is created, to finds the enough number of outputs holding the given amount to be spent. It uses the UTXO set instead of search in the blockchain.
- `UTXOSet.FindUTXO`: finds unspent outputs for a given address. Can be used to get the balance. It uses the UTXO set instead of search in the blockchain.
- `UTXOSet.Update`: updates the current UTXO set with a new set of transactions. Should ignore transactions that try to double-spend (i.e., use the same output twice).

[btctransactions]: btc-transactions.svg "Figure 3"
[utxo]: utxo.png "Figure 4"

## Demo Application

Extend your client application of lab 1 to create two addresses and make a demo transaction between the two generated addresses. Show the correspondent blocks where the transactions were added and the whole blockchain state in the end of the process. Create a command to get and show the current balance of both addresses based on the UTXO of each one.

The `getBalance` function should receive a address and the current UTXO set and return the balance.


## Lab Approval

To have your lab assignment approved, you must come to the lab during lab hours and present your solution. This lets you present the thought process behind your solution, and allows us to provide feedback on your solution then and there.
When you are ready to show your solution, reach out to a member of the teaching staff. It is expected that you can explain your code and show how it works. You may show your solution on a lab workstation or your own computer.

You should for this lab present a working demo of the application described in the previous section making a transaction between two addresses. 
You should demonstrate that your implementation fulfills the previously listed specification of each assignments part.
The task will be verified by a member of the teaching staff during lab hours.
