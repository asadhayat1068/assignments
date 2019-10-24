package main

//go:generate abigen --sol ../../contracts/Simple.sol --pkg contract --out ./go-bindings/simple-contract/simple.go

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	contract "./go-bindings/simple-contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	defaultAddress = "0x269198657177722fCb88F371Ab5072d5570Ce017"
	defaultPKHex   = "47f7b766d34ef3c11959ede82c6de083cc016b6538d42a6ddfb9c7a4e43f4576"
)

var commands = []string{
	"Deploy", "GetBalance", "get", "set", "stop", "Quit",
}

type Client struct {
	scanner          *bufio.Scanner
	backend          *ethclient.Client
	contractAddress  common.Address
	contractInstance *contract.Simple
}

func connect() *ethclient.Client {
	backend, err := ethclient.Dial("ws://127.0.0.1:7555")
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

func (c Client) deploy(auth *bind.TransactOpts) (common.Address, *types.Transaction, *contract.Simple, error) {
	return contract.DeploySimple(auth, c.backend)
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

		fmt.Println("------------------\nChoose a command:")
		for i, c := range commands {
			fmt.Printf("(%v) %s\n", i+1, c)
		}
		fmt.Scanln(&cmd)
		fmt.Println("------------------")
		switch cmd {
		case 1:
			// Deploy
			auth := client.getAuth(privateKey)
			address, _, instance, err := client.deploy(auth)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			client.contractAddress = address
			client.contractInstance = instance
			fmt.Println("Contract deployed at:", address.Hex())
		case 2:
			// GetBalance
			balance := client.getBalance(client.contractAddress)
			// contractInstance.GetBalance(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Println("Current contract balance is:", balance)
		case 3:
			// get value
			balance, err := client.contractInstance.Get(&bind.CallOpts{Pending: true})
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Println("The value is:", balance)
		case 4:
			// set value
			fmt.Println("Enter the amount to send for set (in wei):")
			client.scanner.Scan()
			amount, _ := big.NewInt(0).SetString(client.scanner.Text(), 10)
			value := "Default"
			fmt.Println("Enter new value to set:")
			fmt.Scanln(&value)

			auth := client.getAuth(privateKey)
			tx, err := client.contractInstance.Set(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
				Value:  amount,
			}, value)
			if err != nil {
				fmt.Printf("An error occur: %v\n", err)
				continue
			}
			fmt.Printf("Transaction 0x%x successfully created\n", tx.Hash())
		case 5:
			//stop
			auth := client.getAuth(privateKey)
			tx, err := client.contractInstance.Stop(&bind.TransactOpts{
				From:   auth.From,
				Signer: auth.Signer,
			})
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
