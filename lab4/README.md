![UiS](https://www.uis.no/getfile.php/13391907/Biblioteket/Logo%20og%20veiledninger/UiS_liggende_logo_liten.png)

# Lab 4: Introduction to Smart Contracts

| Lab 4:           | Smart Contracts              |
| ---------------- | ---------------------------- |
| Subject:         | DAT650 Blockchain Technology |
| Deadline:        | 10. OCT                      |
| Expected effort: | 1 weeks                      |
| Grading:         | Pass/fail                    |

## Table of Contents
- [Lab 4: Introduction to Smart Contracts](#lab-4-introduction-to-smart-contracts)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Tools](#tools)
    - [Install and Running](#install-and-running)
  - [Lab Approval](#lab-approval)

## Introduction

In this lab you will require to write two smart contracts.
The description of each part of the assignment can be found under the respective directories, [Wallet contract](wallet/README.md) and [Betting contract](betting/README.md).
All tests should pass, but you can add as many functions you need.

## Tools

For this lab you will require to have [Truffle Suite](https://www.trufflesuite.com/docs/truffle/overview) installed in your machine.
It is also desirable that you have [Ganache](https://www.trufflesuite.com/docs/ganache/overview) installed or any other develop blockchain configured, with at least two accounts, to perform correctly the tests.

Both assignments have a `package.json` file with the dependencies and scripts for easy development, including truffle and ganache-cli, that can be installed locally using the `npm`.

* Note that the commands shown below and specified in the scripts section of the `package.json` file are optional. If you have truffle installed globally in your system you can use it instead, by running direct the commands specified in the scripts.

### Install and Running

To install the necessary dependencies to run and test each assignment, enter in the correspondent directory and run the `npm install` command. Like the example below for the wallet project:

```
$ cd wallet
$ npm install
```

After the installation you can compile and run the tests as following:
```
$ npm run compile
$ npm run migrate:ganache
$ npm run test:ganache
```

If you get the following error:

```
Could not connect to your Ethereum client with the following parameters:
    - host       > 127.0.0.1
    - port       > 8545
    - network_id > *
Please check that your Ethereum client:
    - is running
    - is accepting RPC connections (i.e., "--rpc" option is used in geth)
    - is accessible over the network
    - is properly configured in your Truffle configuration file (truffle-config.js)
```

It means that you need to have running a blockchain instance in another terminal.
There are many options for perform that, and you can find more information [here](https://www.trufflesuite.com/docs/truffle/reference/choosing-an-ethereum-client).
For the purpose of this lab, we will be using the `ganache` GUI or the `ganache-cli` command, both with same setup.
More information about the ganache configuration can be found [here](https://www.trufflesuite.com/docs/ganache/truffle-projects/linking-a-truffle-project)

```
$ ganache-cli --deterministic --networkId 5777 --host 127.0.0.1 --port 7545
```

## Lab Approval

To have your lab assignment approved, you must come to the lab during lab hours and present your solution. This lets you present the thought process behind your solution, and allows us to provide feedback on your solution then and there.
When you are ready to show your solution, reach out to a member of the teaching staff. It is expected that you can explain your code and show how it works. You may show your solution on a lab workstation or your own computer.

You should demonstrate that your implementation fulfills the previously listed specification of each assignments part.
The task will be verified by a member of the teaching staff during lab hours.