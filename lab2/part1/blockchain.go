package main

import (
	"errors"
	"fmt"
	"strings"
)

// Blockchain keeps a sequence of Blocks
type Blockchain struct {
	blocks []*Block
}

// CreateBlockchain creates a new blockchain with genesis Block
func CreateBlockchain(address string) *Blockchain {
	// TODO(student)
	// ADD the creation of the initial UTXO set
	return nil
}

// NewBlockchain creates a Blockchain
func NewBlockchain(address string) *Blockchain {
	return CreateBlockchain(address)
}

// AddBlock saves the block into the blockchain
func (bc *Blockchain) AddBlock(transactions []*Transaction) *Block {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil
}

// GetGenesisBlock returns the Genesis Block
func (bc Blockchain) GetGenesisBlock() *Block {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil
}

// CurrentBlock returns the last block
func (bc Blockchain) CurrentBlock() *Block {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil
}

// GetBlock returns the block of a given hash
func (bc Blockchain) GetBlock(hash []byte) (*Block, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil, nil
}

// MineBlock mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) (*Block, error) {
	// TODO(student)
	// 1) Verify the existence of transactions inputs and discard invalid transactions that make reference to unknown inputs
	// 2) Add a block if there is a list of valid transactions
	return nil, errors.New("there are no valid transactions to be mined")
}

// VerifyTransaction verifies if referred inputs exist
func (bc Blockchain) VerifyTransaction(tx *Transaction) bool {
	// TODO(student)
	// Check if all inputs of a given transaction refer to a existent transaction made previously
	// TIP: remember that Coinbase transaction doesn't have input
	return true
}

// FindTransaction finds a transaction by its ID in the whole blockchain
func (bc Blockchain) FindTransaction(ID []byte) (*Transaction, error) {
	// TODO(student)
	// TIP: the chain is made of what?
	return nil, errors.New("Transaction not found in any block")
}

// FindUTXOSet finds and returns all unspent transaction outputs
func (bc Blockchain) FindUTXOSet() UTXOSet {
	UTXO := make(UTXOSet)

	// TODO(student)
	// 1) Search in the blockchain for unspent transactions outputs
	// 2) Ignore an already spent output
	// TIP: what determines that an output was spent?

	return UTXO
}

func (bc Blockchain) String() string {
	var lines []string
	for _, block := range bc.blocks {
		lines = append(lines, fmt.Sprintf("%v", block))
	}
	return strings.Join(lines, "\n")
}
