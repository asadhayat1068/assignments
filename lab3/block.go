package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Block keeps block information
type Block struct {
	Timestamp     int64          // the block creation timestamp
	Transactions  []*Transaction // The block transactions
	PrevBlockHash []byte         // the hash of the previous block
	Hash          []byte         // the hash of the block
	Nonce         int            // the nonce of the block
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}, 0}
	// TODO(student) -- Replace SetHash by Mine
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// Mine calculates and sets the block hash and nonce.
func (b *Block) Mine() {
	// TODO(student)
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return []byte{}
}

// FindTransaction finds a transaction by its ID
func (b *Block) FindTransaction(ID []byte) (*Transaction, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil, errors.New("Transaction not found")
}

func (b *Block) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("============ Block %x ============", b.Hash))
	lines = append(lines, fmt.Sprintf("Prev. hash: %x", b.PrevBlockHash))
	lines = append(lines, fmt.Sprintf("Timestamp: %v\n", time.Unix(b.Timestamp, 0)))
	for _, tx := range b.Transactions {
		lines = append(lines, fmt.Sprintf("%v\n", tx))
	}
	return strings.Join(lines, "\n")
}
