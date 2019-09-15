package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestInputsTX(tx string) []TXInput {
	return testTransactions[tx].Vin
}

var testTransactions = map[string]*Transaction{
	"tx0": &Transaction{
		ID: Hex2Bytes("e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"),
		Vin: []TXInput{
			TXInput{Txid: nil, OutIdx: -1, ScriptSig: GenesisCoinbaseData},
		},
		Vout: []TXOutput{
			TXOutput{Value: BlockReward, ScriptPubKey: "rodrigo"},
		},
	},
	"tx1": &Transaction{
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
	"tx2": &Transaction{
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
	"tx3": &Transaction{
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
	"tx4": &Transaction{
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
	"tx5": &Transaction{
		ID: Hex2Bytes("33481b10feed4a93eab2233c68cf119281ac23ed268f1389e076b426ba8b412a:"),
		Vin: []TXInput{
			TXInput{
				Txid:      Hex2Bytes("536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b"),
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
	assert.Equal(t, testTransactions["tx0"].ID, tx1.Hash())

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
	assert.Equal(t, testTransactions["tx1"].ID, tx2.Hash())
}

func TestIsCoinbase(t *testing.T) {

	tx := testTransactions["tx0"]
	assert.True(t, tx.IsCoinbase())

	tx = testTransactions["tx1"]
	assert.False(t, tx.IsCoinbase())
}

func TestNewCoinbaseTX(t *testing.T) {

	// Passing data to the coinbase transaction
	tx := NewCoinbaseTX("leander", "test")
	assert.Equal(t, -1, tx.Vin[0].OutIdx)
	assert.Equal(t, []byte{}, tx.Vin[0].Txid)
	assert.Equal(t, "test", tx.Vin[0].ScriptSig)
	assert.Equal(t, BlockReward, tx.Vout[0].Value)
	assert.Equal(t, "leander", tx.Vout[0].ScriptPubKey)

	// Using default data
	tx = NewCoinbaseTX("leander", "")
	assert.Equal(t, -1, tx.Vin[0].OutIdx)
	assert.Equal(t, []byte{}, tx.Vin[0].Txid)
	assert.Equal(t, "Reward to leander", tx.Vin[0].ScriptSig)
	assert.Equal(t, BlockReward, tx.Vout[0].Value)
	assert.Equal(t, "leander", tx.Vout[0].ScriptPubKey)
}

func TestNewUTXOTransaction(t *testing.T) {
	from := "rodrigo"
	to := "leander"

	// "From" address have 10 (i.e., genesis coinbase)
	// and "to" address have 0
	utxos := UTXOSet{
		"e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d": testTransactions["tx0"].Vout,
	}

	// Reject if there is not sufficient funds
	tx1, err := NewUTXOTransaction(to, from, 5, utxos)
	assert.Errorf(t, err, "Not enough funds")
	assert.Nil(t, tx1)

	// Accept otherwise
	tx1, err = NewUTXOTransaction(from, to, 5, utxos)
	assert.Equal(t, testTransactions["tx1"], tx1)
	assert.Nil(t, err)

	utxos = UTXOSet{
		"bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa": testTransactions["tx1"].Vout,
	}

	tx2, err := NewUTXOTransaction(from, to, 1, utxos)
	assert.Equal(t, testTransactions["tx2"], tx2)
	assert.Nil(t, err)

	tx3, err := NewUTXOTransaction(to, from, 3, utxos)
	assert.Equal(t, testTransactions["tx3"], tx3)
	assert.Nil(t, err)
}
