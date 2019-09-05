package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockchain(t *testing.T) {
	const from = "rodrigo"
	const to = "leander"

	// NewBlockchain
	bc := NewBlockchain(from)
	assert.NotNil(t, bc)
	utxos := bc.FindUTXOSet()
	assert.NotNil(t, utxos)

	// GetGenesisBlock
	gb := bc.GetGenesisBlock()
	assert.NotNil(t, gb)

	// AddBlock
	tx, err := NewUTXOTransaction(from, to, 4, utxos)
	assert.Nil(t, err)
	assert.NotNil(t, tx)

	b1 := bc.AddBlock([]*Transaction{tx})
	assert.NotNil(t, b1)

	if !bytes.Equal(gb.Hash, b1.PrevBlockHash) {
		t.Errorf("Genesis block Hash %x isn't equal to current PrevBlockHash %x", gb.Hash, b1.PrevBlockHash)
	}

	// Update the utxos with the new block
	// to get the new balance of "to"
	utxos.Update(b1.Transactions)

	tx, err = NewUTXOTransaction(to, from, 2, utxos)
	assert.Nil(t, err)
	assert.NotNil(t, tx)

	b2 := bc.AddBlock([]*Transaction{tx})
	assert.NotNil(t, b2)

	utxos.Update(b2.Transactions)
	if !bytes.Equal(b1.Hash, b2.PrevBlockHash) {
		t.Errorf("Previous block Hash %x isn't equal to current PrevBlockHash %x", b1.Hash, b2.PrevBlockHash)
	}

	// CurrentBlock
	b3 := bc.CurrentBlock()
	assert.NotNil(t, b3)

	if !bytes.Equal(b3.Hash, b2.Hash) {
		t.Errorf("Current block Hash %x isn't the expected %x", b3.Hash, b2.Hash)
	}

	// GetBlock
	b4, err := bc.GetBlock(b2.Hash)
	assert.Nil(t, err)
	assert.NotNil(t, b4)

	if !bytes.Equal(b4.Hash, b2.Hash) {
		t.Errorf("Block Hash %x isn't the expected %x", b4.Hash, b2.Hash)
	}
}

func TestMineBlockWithInvalidTxInput(t *testing.T) {
	bc := NewBlockchain("rodrigo")

	// Ignore transaction that refer to non-existent transaction input
	invalidTx := &Transaction{
		ID: Hex2Bytes("bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa"),
		Vin: []TXInput{
			TXInput{
				Txid:      Hex2Bytes("non-existentID"),
				OutIdx:    0,
				ScriptSig: "rodrigo",
			},
		},
		Vout: []TXOutput{
			TXOutput{Value: 5, ScriptPubKey: "leander"},
			TXOutput{Value: 5, ScriptPubKey: "rodrigo"},
		},
	}

	b, err := bc.MineBlock([]*Transaction{invalidTx})
	assert.Error(t, err, "there are no valid transactions to be mined")
	assert.Nil(t, b)
}

func TestMineBlock(t *testing.T) {
	bc := NewBlockchain("rodrigo")
	tx := &Transaction{
		ID: Hex2Bytes("bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa"),
		Vin: []TXInput{
			TXInput{
				Txid:      Hex2Bytes("e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"),
				OutIdx:    0,
				ScriptSig: "rodrigo",
			},
		},
		Vout: []TXOutput{
			TXOutput{Value: 5, ScriptPubKey: "leander"},
			TXOutput{Value: 5, ScriptPubKey: "rodrigo"},
		},
	}

	b, err := bc.MineBlock([]*Transaction{tx})
	assert.Nil(t, err)

	minedBlock, err := bc.GetBlock(b.Hash)
	assert.Equal(t, b, minedBlock)

	txMinedBlock, err := b.FindTransaction(tx.ID)
	assert.Nil(t, err)
	assert.NotNil(t, txMinedBlock)
}

func TestVerifyTransaction(t *testing.T) {
	bc := NewBlockchain("rodrigo")

	invalidTx := &Transaction{
		ID: Hex2Bytes("bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa"),
		Vin: []TXInput{
			TXInput{
				Txid:      Hex2Bytes("non-existentID"),
				OutIdx:    0,
				ScriptSig: "rodrigo",
			},
		},
		Vout: []TXOutput{
			TXOutput{Value: 5, ScriptPubKey: "leander"},
			TXOutput{Value: 5, ScriptPubKey: "rodrigo"},
		},
	}
	assert.False(t, bc.VerifyTransaction(invalidTx))

	genesisTx := bc.GetGenesisBlock().Transactions[0]
	assert.True(t, bc.VerifyTransaction(genesisTx))
}

func TestFindTransaction(t *testing.T) {
	bc := NewBlockchain("rodrigo")

	// Find coinbase genesis transaction
	tx, err := bc.FindTransaction(Hex2Bytes("e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"))
	assert.Nil(t, err)
	assert.NotNil(t, tx)

	notFoundTx, err := bc.FindTransaction(Hex2Bytes("non-existentID"))
	assert.Error(t, err, "Transaction not found")
	assert.Nil(t, notFoundTx)
}

func TestFindUTXOSet(t *testing.T) {

	bc := NewBlockchain("rodrigo")
	initialUTXO := bc.FindUTXOSet()
	assert.Equal(t, GetTestGenesisUTXOSet(), initialUTXO)

	utxos := bc.FindUTXOSet()
	assert.Equal(t, initialUTXO, utxos)

	b, err := bc.MineBlock(GetTestTransactions("block1"))
	assert.Nil(t, err)
	initialUTXO.Update(b.Transactions)

	utxos = bc.FindUTXOSet()
	assert.Equal(t, initialUTXO, utxos)
}
