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

	assert.Equal(t, []byte{}, gb.PrevBlockHash, "Genesis block shouldn't has PrevBlockHash")

	assert.Equal(t, testTransactions["tx0"].ID, gb.Transactions[0].ID, "Genesis block should contains the coinbase transaction")
}

func TestBlockHashTransactions(t *testing.T) {
	b := &Block{BlockTime, []*Transaction{testTransactions["tx1"]}, nil, []byte{}}

	merkleRootTxsHash := Hex2Bytes("a71240865cfd49552de1c40a2582065d7c0edfe27d906c90ff203dc8a8664a37")
	assert.Equalf(t, merkleRootTxsHash, b.HashTransactions(), "The block hash %x isn't equal to %x", b.HashTransactions(), merkleRootTxsHash)
}

func TestSetHash(t *testing.T) {
	genesisBlock := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})

	b1 := &Block{BlockTime, []*Transaction{testTransactions["tx1"]}, genesisBlock.Hash, []byte{}}
	b1.SetHash()

	expectedHeaderHash := Hex2Bytes("41172f5ac8c38746abbbc12e5b9c8c3c9e306833cce738d846c5281b1f731df8")
	assert.Equalf(t, expectedHeaderHash, b1.Hash, "The block hash %x isn't equal to %x", b1.Hash, expectedHeaderHash)
}

func TestNewBlock(t *testing.T) {
	genesisBlock := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})

	b := NewBlock([]*Transaction{testTransactions["tx1"]}, genesisBlock.Hash)
	expected := &Block{b.Timestamp, []*Transaction{testTransactions["tx1"]}, genesisBlock.Hash, []byte{}}
	expected.SetHash()

	assert.NotEqual(t, []byte{}, b.Hash, "The block hash should have a valid value")

	assert.Equalf(t, expected.Hash, b.Hash, "The block hash %x isn't equal to %x", b.Hash, expected.Hash)

	assert.Equal(t, genesisBlock.Hash, b.PrevBlockHash, "Previous block of the current should be the genesis block")
}
