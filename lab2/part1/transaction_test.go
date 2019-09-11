package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testTransactions = map[string][]*Transaction{
	"genesis": []*Transaction{
		&Transaction{
			ID: Hex2Bytes("e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"),
			Vin: []TXInput{
				TXInput{Txid: nil, OutIdx: -1, ScriptSig: GenesisCoinbaseData},
			},
			Vout: []TXOutput{
				TXOutput{Value: BlockReward, ScriptPubKey: "rodrigo"},
			},
		},
	},
	"block1": []*Transaction{
		&Transaction{
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
		},
	},
	"block2": []*Transaction{
		&Transaction{
			ID: Hex2Bytes("4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a"),
			Vin: []TXInput{
				TXInput{
					Txid:      Hex2Bytes("bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa"),
					OutIdx:    1,
					ScriptSig: "rodrigo",
				},
			},
			Vout: []TXOutput{
				TXOutput{Value: 1, ScriptPubKey: "leander"},
				TXOutput{Value: 4, ScriptPubKey: "rodrigo"},
			},
		},
		&Transaction{
			ID: Hex2Bytes("ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c"),
			Vin: []TXInput{
				TXInput{
					Txid:      Hex2Bytes("bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa"),
					OutIdx:    0,
					ScriptSig: "leander",
				},
			},
			Vout: []TXOutput{
				TXOutput{Value: 3, ScriptPubKey: "rodrigo"},
				TXOutput{Value: 2, ScriptPubKey: "leander"},
			},
		},
	},
	"block3": []*Transaction{
		&Transaction{
			ID: Hex2Bytes("536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b"),
			Vin: []TXInput{
				TXInput{
					Txid:      Hex2Bytes("ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c"),
					OutIdx:    0,
					ScriptSig: "rodrigo",
				},
			},
			Vout: []TXOutput{
				TXOutput{Value: 2, ScriptPubKey: "leander"},
				TXOutput{Value: 1, ScriptPubKey: "rodrigo"},
			},
		},
	},
	"block4": []*Transaction{
		&Transaction{
			ID: Hex2Bytes("5a27a0fd2344fd1bc585e63906e3c9a14cec9cdfd7b1e00553841ec8516535ab"),
			Vin: []TXInput{
				TXInput{
					Txid:      Hex2Bytes("ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c"),
					OutIdx:    0,
					ScriptSig: "leander",
				},
				TXInput{
					Txid:      Hex2Bytes("4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a"),
					OutIdx:    0,
					ScriptSig: "leander",
				},
			},
			Vout: []TXOutput{
				TXOutput{Value: 3, ScriptPubKey: "rodrigo"},
			},
		},
	},
}

func GetTestTransactions(block string) []*Transaction {
	return testTransactions[block]
}

func GetTestGenesisHash() []byte {
	return testTransactions["genesis"][0].ID
}

func GetTestGenesisUTXOSet() UTXOSet {
	utxo := make(UTXOSet)
	vout := testTransactions["genesis"][0].Vout
	txid := hex.EncodeToString(GetTestGenesisHash())
	utxo[txid] = vout
	return utxo
}

func TestHash(t *testing.T) {
	tx1 := Transaction{
		ID: []byte{},
		Vin: []TXInput{
			TXInput{Txid: []byte{}, OutIdx: -1, ScriptSig: GenesisCoinbaseData},
		},
		Vout: []TXOutput{
			TXOutput{Value: BlockReward, ScriptPubKey: "rodrigo"},
		},
	}
	assert.Equal(t, GetTestGenesisHash(), tx1.Hash())

	tx2 := Transaction{
		ID: []byte{},
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
	expectedTxHash := testTransactions["block1"][0].ID
	assert.Equal(t, expectedTxHash, tx2.Hash())
}

func TestIsCoinbase(t *testing.T) {

	tx := GetTestTransactions("genesis")[0]
	assert.True(t, tx.IsCoinbase())

	tx = GetTestTransactions("block1")[0]
	assert.False(t, tx.IsCoinbase())
}

func TestNewCoinbaseTX(t *testing.T) {

	// Passing data to the coinbase transaction
	tx := NewCoinbaseTX("leander", "test")
	assert.True(t, tx.IsCoinbase())
	assert.Equal(t, "test", tx.Vin[0].ScriptSig)

	// Using default data
	tx = NewCoinbaseTX("leander", "")
	assert.True(t, tx.IsCoinbase())
	assert.Equal(t, "Reward to leander", tx.Vin[0].ScriptSig)
	assert.Equal(t, BlockReward, tx.Vout[0].Value)
	assert.Equal(t, "leander", tx.Vout[0].ScriptPubKey)
}

func TestNewUTXOTransaction(t *testing.T) {
	from := "rodrigo"
	to := "leander"

	genesisTx := GetTestTransactions("genesis")[0]
	gb := NewGenesisBlock(genesisTx)
	// "From" address have 10 (i.e., genesis coinbase)
	// and "to" address have 0
	utxos := GetTestGenesisUTXOSet()

	// Reject if there is not sufficient funds
	tx1, err := NewUTXOTransaction(to, from, 5, utxos)
	assert.Errorf(t, err, "Not enough funds")
	assert.Nil(t, tx1)

	// Accept otherwise
	tx1, err = NewUTXOTransaction(from, to, 5, utxos)
	expectedTx1 := GetTestTransactions("block1")[0]
	assert.Equal(t, expectedTx1, tx1)
	assert.Nil(t, err)

	b1 := NewBlock([]*Transaction{tx1}, gb.Hash)
	utxos.Update(b1.Transactions)

	tx2, err := NewUTXOTransaction(from, to, 1, utxos)
	expectedTx2 := GetTestTransactions("block2")[0]
	assert.Equal(t, expectedTx2, tx2)
	assert.Nil(t, err)

	tx3, err := NewUTXOTransaction(to, from, 3, utxos)
	expectedTx3 := GetTestTransactions("block2")[1]
	assert.Equal(t, expectedTx3, tx3)
	assert.Nil(t, err)
}
