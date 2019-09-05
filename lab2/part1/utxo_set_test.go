package main

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSpendableOutputsFromOneOutput(t *testing.T) {
	utxos := GetTestGenesisUTXOSet()

	expectedUnspentOutputs := make(map[string][]int)
	txid := hex.EncodeToString(GetTestGenesisHash())
	expectedOutIdx := 0
	expectedUnspentOutputs[txid] = append(expectedUnspentOutputs[txid], expectedOutIdx)
	expectedValue := testTransactions["genesis"][0].Vout[0].Value

	accumulatedAmount, unspentOutputs := utxos.FindSpendableOutputs("rodrigo", 5)

	assert.Equal(t, expectedValue, accumulatedAmount)
	assert.Equal(t, expectedUnspentOutputs, unspentOutputs)
}

func TestFindSpendableOutputsFromMultipleOutputs(t *testing.T) {
	utxos := GetTestGenesisUTXOSet()
	utxos.Update(GetTestTransactions("block1"))
	utxos.Update(GetTestTransactions("block2"))

	expectedUnspentOutputs := make(map[string][]int)
	expectedUnspentOutputs["4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a"] = []int{1}
	expectedUnspentOutputs["ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c"] = []int{0}

	// Get the rodrigo's unspent TXOutput
	out1 := GetTestTransactions("block2")[0].Vout[1].Value
	out2 := GetTestTransactions("block2")[1].Vout[0].Value
	expectedValue := out1 + out2
	accumulatedAmount, unspentOutputs := utxos.FindSpendableOutputs("rodrigo", 5)

	assert.Equal(t, expectedValue, accumulatedAmount)
	assert.Equal(t, expectedUnspentOutputs, unspentOutputs)
}

func TestFindUTXO(t *testing.T) {
	// Rodrigo create a coinbase transaction, receiving 10 "coins"
	utxos := GetTestGenesisUTXOSet()

	utxoRodrigo := utxos.FindUTXO("rodrigo")
	assert.Equal(t, []TXOutput{TXOutput{10, "rodrigo"}}, utxoRodrigo)

	utxoLeander := utxos.FindUTXO("leander")
	assert.Equal(t, []TXOutput(nil), utxoLeander)

	// Rodrigo sent 5 "coins" to Leander
	utxos.Update(GetTestTransactions("block1"))

	utxoRodrigo = utxos.FindUTXO("rodrigo")
	assert.Equal(t, []TXOutput{TXOutput{5, "rodrigo"}}, utxoRodrigo)

	utxoLeander = utxos.FindUTXO("leander")
	assert.Equal(t, []TXOutput{TXOutput{5, "leander"}}, utxoLeander)

	// Rodrigo sent 1 "coin" to Leander and
	// Leander sent 3 "coins" to Rodrigo
	utxos.Update(GetTestTransactions("block2"))

	utxoRodrigo = utxos.FindUTXO("rodrigo")
	assert.ElementsMatch(t, []TXOutput{
		TXOutput{4, "rodrigo"},
		TXOutput{3, "rodrigo"},
	}, utxoRodrigo)
	assert.Equal(t, 2, len(utxoRodrigo))

	utxoLeander = utxos.FindUTXO("leander")
	assert.ElementsMatch(t, []TXOutput{
		TXOutput{2, "leander"},
		TXOutput{1, "leander"},
	}, utxoLeander)
	assert.Equal(t, 2, len(utxoLeander))
}

func TestCountUTXOs(t *testing.T) {
	utxos := GetTestGenesisUTXOSet()
	assert.Equal(t, 1, utxos.CountUTXOs())

	utxos.Update(GetTestTransactions("block1"))
	assert.Equal(t, 2, utxos.CountUTXOs())

	utxos.Update(GetTestTransactions("block2"))
	assert.Equal(t, 4, utxos.CountUTXOs())
}

var testUTXOs = map[string]struct {
	utxos         UTXOSet
	expectedUTXOs []UTXOSet
}{
	"coinbase": { // (0 input -> 1 output, generating "coins")
		utxos: UTXOSet{},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d": []TXOutput{TXOutput{10, "rodrigo"}}, //tx0: Rodrigo create coinbase transaction and received 10 "coins"
			},
		},
	},
	"block1": { // (1 input -> 2 outputs, splitting one input)
		utxos: UTXOSet{
			"e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d": []TXOutput{TXOutput{10, "rodrigo"}},
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa": []TXOutput{TXOutput{5, "leander"}, TXOutput{5, "rodrigo"}}, //tx1: Rodrigo sent 5 "coins" to Leander
			},
		},
	},
	"block2": { // (1 input -> 2 output, with multiple txs)
		utxos: UTXOSet{
			"bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa": []TXOutput{TXOutput{5, "leander"}, TXOutput{5, "rodrigo"}},
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{1, "leander"}, TXOutput{4, "rodrigo"}}, //tx2: Rodrigo sent 1 "coin" to Leander
				"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": []TXOutput{TXOutput{3, "rodrigo"}, TXOutput{2, "leander"}}, //tx3: Leander sent 3 "coins" to Rodrigo
			},
		},
	},
	"block3": { // (1 input -> 2 outputs, with distinct input source)
		utxos: UTXOSet{
			"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{1, "leander"}, TXOutput{4, "rodrigo"}},
			"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": []TXOutput{TXOutput{3, "rodrigo"}, TXOutput{2, "leander"}},
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{ //Depending of the input taken the produced outputs will be different
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{1, "leander"}, TXOutput{4, "rodrigo"}},
				"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": []TXOutput{TXOutput{2, "leander"}},
				// Rodrigo spent his previous output that was here (i.e., TXOutput{3, "rodrigo"}) using it as the input of the new tx below:
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": []TXOutput{TXOutput{2, "leander"}, TXOutput{1, "rodrigo"}}, //tx4: Rodrigo sent 2 "coins" to Leander
			},
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{1, "leander"}},
				"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": []TXOutput{TXOutput{3, "rodrigo"}, TXOutput{2, "leander"}},
				"b60f1d3288b50402564fc5408f26dbf9726d1d79df5f13aff225641e05619491": []TXOutput{TXOutput{2, "leander"}, TXOutput{2, "rodrigo"}}, //tx4: Rodrigo sent 2 "coins" to Leander
			},
		},
	},
	"block4": { // (2 inputs -> 1 output, with different input order)
		utxos: UTXOSet{
			// The TXOutput{1, "leander"} was removed from the tx below to limit the number of possible outputs
			"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{4, "rodrigo"}},
			"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": []TXOutput{TXOutput{2, "leander"}},
			"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": []TXOutput{TXOutput{2, "leander"}, TXOutput{1, "rodrigo"}},
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{ //Depending of the order of the inputs IDs taken the resulted tx hash will be different. But for the given UTXOSet above,one of the two resulted UTXOSet below should be true.
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{4, "rodrigo"}},
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": []TXOutput{TXOutput{1, "rodrigo"}},
				"ac1d8ce2998fef93ada0555fa27375db4a3b5f1082410a251ad7d7bfdbd63a58": []TXOutput{TXOutput{3, "rodrigo"}, TXOutput{1, "leander"}}, //tx5: Leander sent 3 "coins" to Rodrigo
			},
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": []TXOutput{TXOutput{4, "rodrigo"}},
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": []TXOutput{TXOutput{1, "rodrigo"}},
				"77781fcc1960093ae86bb24c364c564e41e18ce881af350c132a82bdaab0c9f0": []TXOutput{TXOutput{3, "rodrigo"}, TXOutput{1, "leander"}}, //tx5: Leander sent 3 "coins" to Rodrigo
			},
		},
	},
}

func TestUpdate(t *testing.T) {
	for k, m := range testUTXOs {
		utxos := m.utxos
		switch k {
		case "coinbase":
			tx := NewCoinbaseTX("rodrigo", GenesisCoinbaseData)
			gb := NewGenesisBlock(tx)
			utxos.Update(gb.Transactions)
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block1":
			tx, err := NewUTXOTransaction("rodrigo", "leander", 5, utxos)
			assert.Nil(t, err)
			utxos.Update([]*Transaction{tx})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block2":
			tx1, err := NewUTXOTransaction("rodrigo", "leander", 1, utxos)
			assert.Nil(t, err)
			tx2, err := NewUTXOTransaction("leander", "rodrigo", 3, utxos)
			assert.Nil(t, err)
			utxos.Update([]*Transaction{tx1, tx2})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block3":
			tx, err := NewUTXOTransaction("rodrigo", "leander", 2, utxos)
			assert.Nil(t, err)
			utxos.Update([]*Transaction{tx})
			ok1 := utxos.Equal(m.expectedUTXOs[0])
			ok2 := utxos.Equal(m.expectedUTXOs[1])
			assert.True(t, ok1 || ok2)
		case "block4":
			tx, err := NewUTXOTransaction("leander", "rodrigo", 3, utxos)
			assert.Nil(t, err)
			utxos.Update([]*Transaction{tx})
			ok1 := utxos.Equal(m.expectedUTXOs[0])
			ok2 := utxos.Equal(m.expectedUTXOs[1])
			assert.True(t, ok1 || ok2)
		}
	}
}
