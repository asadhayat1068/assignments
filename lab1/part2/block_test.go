package main

import (
	"bytes"
	"testing"
)

// Fixed block timestamp
const BlockTime int64 = 1563897484

// Genesis hash with timestamp BlockTime and data: "Genesis data info"
var GenesisHash = Hex2Bytes("556b087be95a9918ca21c2d25e8fcbe2a484299bc621dff9402d94088d1070e3")

// Set of test transactions
var txs = []*Transaction{
	&Transaction{[]byte("tx 1")},
	&Transaction{[]byte("tx 2")},
}

func TestGenesisBlock(t *testing.T) {
	// Genesis block
	tx := Transaction{[]byte("Genesis data info")}
	gb := NewGenesisBlock(&tx)
	if !bytes.Equal([]byte{}, gb.PrevBlockHash) {
		t.Error("Genesis block shouldn't has PrevBlockHash")
	}

	if !bytes.Equal(gb.Transactions[0].Data, []byte("Genesis data info")) {
		t.Error("Genesis data should be stored in the block")
	}
}

func TestBlockHashTransactions(t *testing.T) {
	// Merkle root of txs
	merkleRootTxsHash := Hex2Bytes("857b55c130bc3e5d48c0e5810995a6a1db42d0e241fb9b0559333a456d3bc36e")

	b := &Block{BlockTime, txs, GenesisHash, []byte{}}

	if !bytes.Equal(merkleRootTxsHash, b.HashTransactions()) {
		t.Errorf("The block hash %x isn't equal to %x", b.HashTransactions(), merkleRootTxsHash)
	}
}

func TestSetHash(t *testing.T) {
	// SetHash
	// Hash of headers containing: {GenesisHash, merkleRootTxsHash, BlockTime}
	blockHeaderHash := Hex2Bytes("715738e65cde945ad07d8b16f9e32adc4b3d2e9984c2d712384c1b202bee1337")

	b1 := &Block{BlockTime, txs, GenesisHash, []byte{}}
	b1.SetHash()

	if !bytes.Equal(blockHeaderHash, b1.Hash) {
		t.Errorf("The block hash %x isn't equal to %x", b1.Hash, blockHeaderHash)
	}
}

func TestNewBlock(t *testing.T) {
	// NewBlock
	b2 := NewBlock(txs, GenesisHash)
	b3 := &Block{b2.Timestamp, txs, GenesisHash, []byte{}}
	b3.SetHash()

	if bytes.Equal(b2.Hash, []byte{}) {
		t.Error("The block hash should have a valid value")
	}

	if !bytes.Equal(b2.Hash, b3.Hash) {
		t.Errorf("The block hash %x isn't equal to %x", b3.Hash, b2.Hash)
	}
}
