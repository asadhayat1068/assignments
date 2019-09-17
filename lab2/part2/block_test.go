package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Fixed block timestamp
const BlockTime int64 = 1563897484

func newMockBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{BlockTime, transactions, prevBlockHash, []byte{}}
	block.SetHash() //FIXME: remove this dependency
	return block
}

func TestGenesisBlock(t *testing.T) {
	// Genesis block
	gb := NewGenesisBlock(testTransactions["tx0"])
	assert.NotNil(t, gb)

	assert.Equal(t, []byte{}, gb.PrevBlockHash, "Genesis block shouldn't has PrevBlockHash")

	assert.Equal(t, testTransactions["tx0"].ID, gb.Transactions[0].ID, "Genesis block should contains the coinbase transaction")
}

func TestBlockHashTransactions(t *testing.T) {
	genesisBlock := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})

	// Merkle root of txs
	merkleRootTxsHash := Hex2Bytes("153b097cb029051834d285f69901237812f13d6dc166d684778096422f331677")

	b := &Block{BlockTime, []*Transaction{testTransactions["tx1"]}, genesisBlock.Hash, []byte{}}

	assert.Equalf(t, merkleRootTxsHash, b.HashTransactions(), "The block hash %x isn't equal to %x", b.HashTransactions(), merkleRootTxsHash)
}

func TestSetHash(t *testing.T) {
	genesisBlock := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})

	blockHeaderHash := Hex2Bytes("c6dcb622873e3c37238cabc5b13120f517083910ad2e344771343dd520faa567")

	b := &Block{BlockTime, []*Transaction{testTransactions["tx1"]}, genesisBlock.Hash, []byte{}}
	b.SetHash()

	assert.Equalf(t, blockHeaderHash, b.Hash, "The block hash %x isn't equal to %x", b.Hash, blockHeaderHash)
}

func TestNewBlock(t *testing.T) {
	genesisBlock := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})

	b := NewBlock([]*Transaction{testTransactions["tx1"]}, genesisBlock.Hash)

	expected := &Block{b.Timestamp, []*Transaction{testTransactions["tx1"]}, genesisBlock.Hash, []byte{}}
	expected.SetHash() //FIXME: remove dependency

	assert.NotEqual(t, []byte{}, b.Hash, "The block hash should have a valid value")

	assert.Equalf(t, expected.Hash, b.Hash, "The block hash %x isn't equal to %x", b.Hash, expected.Hash)

	assert.Equal(t, genesisBlock.Hash, b.PrevBlockHash, "Previous block of the current should be the genesis block")
}
