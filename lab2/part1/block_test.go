package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Fixed block timestamp
const BlockTime int64 = 1563897484

func TestGenesisBlock(t *testing.T) {
	// Genesis block
	tx := NewCoinbaseTX("rodrigo", GenesisCoinbaseData)
	assert.NotNilf(t, tx, "Coinbase transaction shouldn't be nil")

	gb := NewGenesisBlock(tx)

	assert.Equal(t, []byte{}, gb.PrevBlockHash, "Genesis block shouldn't has PrevBlockHash")

	assert.Equal(t, tx.Hash(), gb.Transactions[0].ID, "Genesis block should contains the coinbase transaction")

	assert.Equal(t, GetTestGenesisHash(), gb.Transactions[0].ID, "The coinbase transaction in genesis block isn't the expected")
}

func TestBlockHashTransactions(t *testing.T) {
	// Merkle root of txs
	merkleRootTxsHash := Hex2Bytes("a71240865cfd49552de1c40a2582065d7c0edfe27d906c90ff203dc8a8664a37")

	tx, err := NewUTXOTransaction("rodrigo", "leander", 5, GetTestGenesisUTXOSet())
	if err != nil {
		t.Error(err)
	}
	b := &Block{BlockTime, []*Transaction{tx}, GetTestGenesisHash(), []byte{}}

	assert.Equalf(t, merkleRootTxsHash, b.HashTransactions(), "The block hash %x isn't equal to %x", b.HashTransactions(), merkleRootTxsHash)
}

func TestSetHash(t *testing.T) {
	// SetHash
	blockHeaderHash := Hex2Bytes("a2967beb708b6a23603c02bff979fd40327cd475a08a6bf434c7d8a75f406b00")

	tx, err := NewUTXOTransaction("rodrigo", "leander", 5, GetTestGenesisUTXOSet())
	if err != nil {
		t.Error(err)
	}
	b1 := &Block{BlockTime, []*Transaction{tx}, GetTestGenesisHash(), []byte{}}
	b1.SetHash()

	assert.Equalf(t, blockHeaderHash, b1.Hash, "The block hash %x isn't equal to %x", b1.Hash, blockHeaderHash)
}

func TestNewBlock(t *testing.T) {
	// NewBlock
	tx, err := NewUTXOTransaction("rodrigo", "leander", 5, GetTestGenesisUTXOSet())
	if err != nil {
		t.Error(err)
	}
	b2 := NewBlock([]*Transaction{tx}, GetTestGenesisHash())
	expected := &Block{b2.Timestamp, []*Transaction{tx}, GetTestGenesisHash(), []byte{}}
	expected.SetHash()

	assert.NotEqual(t, []byte{}, b2.Hash, "The block hash should have a valid value")

	assert.Equalf(t, expected.Hash, b2.Hash, "The block hash %x isn't equal to %x", b2.Hash, expected.Hash)
}
