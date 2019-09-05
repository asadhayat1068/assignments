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
}

// NewBlock creates and returns Block
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions, prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

// NewGenesisBlock creates and returns genesis Block
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

// SetHash calculates and sets the block hash
func (b *Block) SetHash() {
	var headers []byte
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	//  - b.HashTransactions() need to be used here when combining the block header data.
	//  - You should set the block hash to be the hash of the header, so the line below should be changed.
	b.Hash = headers[:]
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTransactions() []byte {
	var merkleRoot [32]byte
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// You should return the merkle root hash
	return merkleRoot[:]
}

// FindTransaction finds a transaction by its ID
func (b *Block) FindTransaction(ID []byte) (*Transaction, error) {
	// TODO(student)
	// search for a transaction using the tx.ID
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
