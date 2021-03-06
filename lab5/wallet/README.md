![UiS](https://www.uis.no/getfile.php/13391907/Biblioteket/Logo%20og%20veiledninger/UiS_liggende_logo_liten.png)

# Lab 5: Developing a DApp

| Lab 5:           | Distributed Application      |
| ---------------- | ---------------------------- |
| Subject:         | DAT650 Blockchain Technology |
| Deadline:        | 17. OCT                      |
| Expected effort: | 1 week                       |
| Grading:         | Pass/fail                    |

## Table of Contents
- [Lab 5: Developing a DApp](#lab-5-developing-a-dapp)
  - [Table of Contents](#table-of-contents)
  - [Introduction](#introduction)
  - [Building the contracts (backend)](#building-the-contracts-backend)
  - [Client application](#client-application)
  - [Lab Approval](#lab-approval)

## Introduction

In this lab you will develop a client application for the [Betting contract](../lab4/betting/README.md) created in the [previous lab](../lab4/README.md).

## Building the contracts (backend)

In the lab5 folder, run the following commands:

1. Installing the necessary dependencies.
```
npm install
```

2. Running the development blockchain environment.
```
npm run start:ganache
```

3. Compiling and migrating the smart contracts.
```
npm run compile
npm run migrate:ganache
```

## Client application

A client implementation for the [Wallet contract](../lab4/wallet/README.md) is given as an example in different languages: in javascript under the directory [client/js](client/js/README.md) and in go under the directory [client/go](client/go/README.md) .

## Lab Approval

Your task is to implement all functions exposed by the Betting contract in your client application, allowing an user to interact with your contract using a web browser or a command line application. 

To have your lab assignment approved, you must come to the lab during lab hours and present your solution. This lets you present the thought process behind your solution, and allows us to provide feedback on your solution then and there.
When you are ready to show your solution, reach out to a member of the teaching staff. It is expected that you can explain your code and show how it works. You may show your solution on a lab workstation or your own computer.

You should demonstrate that your implementation fulfills the previously listed specification of each assignments part.
The task will be verified by a member of the teaching staff during lab hours.
