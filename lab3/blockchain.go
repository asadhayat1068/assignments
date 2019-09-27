package main

import (
	"crypto/ecdsa"
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
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
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

// ValidateBlock validates the a block after mining or
// before adding it to the blockchain
func (bc *Blockchain) ValidateBlock(block *Block) bool {
	// TODO(student)
	// check if and only if the first tx is coinbase
	// validates block's Proof-Of-Work
	return true
}

// MineBlock mines a new block with the provided transactions
func (bc *Blockchain) MineBlock(transactions []*Transaction) (*Block, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil, errors.New("there are no valid transactions to be mined")
}

// FindTransaction finds a transaction by its ID in the whole blockchain
func (bc Blockchain) FindTransaction(ID []byte) (*Transaction, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil, errors.New("Transaction not found in any block")
}

// FindUTXOSet finds and returns all unspent transaction outputs
func (bc Blockchain) FindUTXOSet() UTXOSet {
	UTXO := make(UTXOSet)
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return UTXO
}

// GetInputTXsOf returns a map index by the ID,
// of all transactions used as inputs in the given transaction
func (bc *Blockchain) GetInputTXsOf(tx *Transaction) (map[string]*Transaction, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// Use bc.FindTransaction to search over all transactions
	// in the blockchain and if the referred input into tx exists,
	// if so, get the transaction of this input and add it
	// to a map, where the key is the id of the transaction found
	// and the value is the pointer to transaction itself.
	// To use the id as key in the map, convert it to string
	// using the function: hex.EncodeToString
	// https://golang.org/pkg/encoding/hex/#EncodeToString
	return nil, nil
}

// SignTransaction signs inputs of a Transaction
func (bc *Blockchain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// Get the previous transactions referred in the input of tx
	// and call Sign for tx.
}

// VerifyTransaction verifies transaction input signatures
func (bc *Blockchain) VerifyTransaction(tx *Transaction) bool {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// Modify the function to get the inputs referred in tx
	// and return false in case of some error (i.e. not found the input).
	// Then call Verify for tx passing those inputs as parameter and return the result.
	// Remember that coinbase transaction doesn't have input.
	return true
}

func (bc Blockchain) String() string {
	var lines []string
	for _, block := range bc.blocks {
		lines = append(lines, fmt.Sprintf("%v", block))
	}
	return strings.Join(lines, "\n")
}
