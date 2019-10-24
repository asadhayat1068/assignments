package main

//go:generate abigen --sol ../../contracts/MyWallet.sol --pkg contract --out ./go-bindings/mywallet-contract/mywallet.go

import (
	"./go-bindings/mywallet-contract"
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"os"
)

var (
	defaultAddress  = "0x90f8bf6a479f320ead074411a4b0e7944ea8c9c1"
	defaultPKHex    = "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
	receiverAddress = "0xffcf8fdee72ac11b5c542428b35eef5769c409f0"
)

var commands = []string{
	"Deploy", "GetBalance", "Deposit", "Withdraw", "Transfer", "Quit",
}

type Client struct {
	scanner        *bufio.Scanner
	backend        *ethclient.Client
	walletAddress  common.Address
	walletInstance *contract.MyWallet
}

func connect() *ethclient.Client {
	backend, err := ethclient.Dial("ws://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	return backend
}

func (c Client) getAuth(privateKey *ecdsa.PrivateKey) *bind.TransactOpts {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.backend.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := c.backend.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(6721975) // in units
	auth.GasPrice = gasPrice
	return auth
}

func (c Client) deploy(auth *bind.TransactOpts) (common.Address, *types.Transaction, *contract.MyWallet, error) {
	return contract.DeployMyWallet(auth, c.backend)
}

func (c Client) getBalance(address common.Address) *big.Int {
	balance, err := c.backend.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func main() {
	client := &Client{}
	client.backend = connect()
	defer client.backend.Close()

	// Get first account as the deployer/sender account
	privateKey, err := crypto.HexToECDSA(defaultPKHex)
	if err != nil {
		log.Fatal(err)
	}

	client.scanner = bufio.NewScanner(os.Stdin)
	var cmd int
	for {
		defaultAccount := common.HexToAddress(defaultAddress)
		balance := client.getBalance(defaultAccount)
		fmt.Printf("Balance of account %s: %v\n", defaultAccount.Hex(), balance)

		receiverAccount := common.HexToAddress(receiverAddress)
		balance = client.getBalance(receiverAccount)
		fmt.Printf("Balance of account %s: %v\n", receiverAccount.Hex(), balance)
		fmt.Println("------------------\nChoose a command:")
		for i, c := range commands {
			fmt.Printf("(%v) %s\n", i+1, c)
		}
		fmt.Scanln(&cmd)
		fmt.Println("------------------")
		switch cmd {
		case 1:
			auth := client.getAuth(privateKey)
			address, _, instance, err := client.deploy(auth)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			client.walletAddress = address
			client.walletInstance = instance
			fmt.Println("Contract deployed at:", address.Hex())
		case 2:
			balance, err := client.walletInstance.GetBalance(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Println("Current wallet balance is:", balance)
		case 3:
			fmt.Println("Enter the amount to deposit (in wei):")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			auth := client.getAuth(privateKey)
			tx, err := client.walletInstance.Deposit(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
				Value:  amount,
			})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		case 4:
			fmt.Println("Enter the amount to withdraw (in wei):")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			auth := client.getAuth(privateKey)
			tx, err := client.walletInstance.Withdraw(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
			}, amount)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		case 5:
			fmt.Println("Enter the recipient address")
			client.scanner.Scan()
			recipient := common.HexToAddress(client.scanner.Text())
			fmt.Println("Enter the amount to transfer (in wei):")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			auth := client.getAuth(privateKey)
			tx, err := client.walletInstance.Transfer(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
			}, recipient, amount)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		case 6:
			return
		}
	}
}
