package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTestExpectedUTXOSet(block string) []UTXOSet {
	return testUTXOs[block].expectedUTXOs
}

func getTestSpendableOutputs(utxos UTXOSet, unlockingData string) map[string][]int {
	unspentOutputs := make(map[string][]int)

	for txID, outputs := range utxos {
		for outIdx, out := range outputs {
			if out.ScriptPubKey == unlockingData {
				unspentOutputs[txID] = append(unspentOutputs[txID], outIdx)
			}
		}
	}
	return unspentOutputs
}

var testUTXOs = map[string]struct {
	utxos         UTXOSet
	expectedUTXOs []UTXOSet
}{
	"genesis": { // (0 input -> 1 output, generating "coins")
		utxos: UTXOSet{},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d": testTransactions["tx0"].Vout,
				//tx0: rodrigo create coinbase transaction and received 10 "coins"
			},
		},
	},
	"block1": { // (1 input -> 2 outputs, splitting one input)
		utxos: UTXOSet{
			"e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d": testTransactions["tx0"].Vout,
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa": testTransactions["tx1"].Vout, //tx1: Rodrigo sent 5 "coins" to Leander and get 5 as remainder
			},
		},
	},
	"block2": { // (1 input -> 2 output, with multiple txs)
		utxos: UTXOSet{
			"bce268225bc12a0015bcc39e91d59f47fd176e64ca42e4f8aecf107fe38f3bfa": testTransactions["tx1"].Vout,
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout, //tx2: Rodrigo sent 1 "coin" to Leander and get 4 as remainder
				"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": testTransactions["tx3"].Vout, //tx3: Leander sent 3 "coins" to Rodrigo and get 2 as remainder
			},
		},
	},
	"block3": { // (1 input -> 2 outputs)
		utxos: UTXOSet{
			"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout[:1], // ignore output to reduce combination of possible input sources
			"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": testTransactions["tx3"].Vout,
		},
		expectedUTXOs: []UTXOSet{
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout[:1],
				"ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c": testTransactions["tx3"].Vout[1:], // spent first output
				// Rodrigo spent his previous output that was here (i.e., TXOutput{3, rodrigoPubKeyHash}) using it as the input of the new tx below:
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": testTransactions["tx4"].Vout, //tx4: Rodrigo sent 2 "coins" to Leander and get 1 as remainder
			},
		},
	},
	"block4": { // (2 inputs -> 1 output)
		utxos: UTXOSet{
			"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout,
			"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": testTransactions["tx4"].Vout,
		},
		expectedUTXOs: []UTXOSet{ // A map is not ordered different positions of the inputs generates a different hash of the transaction, this is why here we are duplicating the utxo set only changing the tx's hash and testing if either exists.
			UTXOSet{
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout[1:], // spent first output
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": testTransactions["tx4"].Vout[1:], // spent first output
				"a99e0d65344022a0acea8267cd6f7a0b1c3162616e1c6080429f7eaa2fd26c64": testTransactions["tx5"].Vout,     //tx5: Leander sent 3 "coins" to Rodrigo
			},
			UTXOSet{ //FIXME: isn't the best approach, but will work for now.
				"536532f5e3a5043e2de1be480761602b8218bf7abf374ece8794e0ca0d5b072b": testTransactions["tx4"].Vout[1:], // spent first output
				"4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a": testTransactions["tx2"].Vout[1:], // spent first output
				"33481b10feed4a93eab2233c68cf119281ac23ed268f1389e076b426ba8b412a": testTransactions["tx5"].Vout,     //tx5: Leander sent 3 "coins" to Rodrigo
			},
		},
	},
}

func TestFindSpendableOutputsFromOneOutput(t *testing.T) {
	utxos := getTestExpectedUTXOSet("genesis")[0]
	expectedUnspentOutputs := getTestSpendableOutputs(utxos, "rodrigo")
	expectedOut := utxos["e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"][0]
	expectedValue := expectedOut.Value

	accumulatedAmount, unspentOutputs := utxos.FindSpendableOutputs("rodrigo", 5)

	assert.Equal(t, expectedValue, accumulatedAmount)
	assert.Equal(t, expectedUnspentOutputs, unspentOutputs)
}

func TestFindSpendableOutputsFromMultipleOutputs(t *testing.T) {
	utxos := getTestExpectedUTXOSet("block2")[0]
	out1 := utxos["4709fe985b8cb59ee67ec3d9ebf968ec82953806332573ba8b2b68d05d9d143a"][1]
	out2 := utxos["ac255688b1df7c2f16fa23cf02f4fe6cb1e500793cc6f9b7d58b58547bfa660c"][0]
	expectedValue := out1.Value + out2.Value
	expectedUnspentOutputs := getTestSpendableOutputs(utxos, "rodrigo")

	accumulatedAmount, unspentOutputs := utxos.FindSpendableOutputs("rodrigo", 5)

	assert.Equal(t, expectedValue, accumulatedAmount)
	assert.Equal(t, expectedUnspentOutputs, unspentOutputs)
}

func TestFindUTXO(t *testing.T) {
	// Rodrigo create a coinbase transaction, receiving 10 "coins"
	utxos := getTestExpectedUTXOSet("genesis")[0]

	utxoRodrigo := utxos.FindUTXO("rodrigo")
	assert.Equal(t, []TXOutput{TXOutput{10, "rodrigo"}}, utxoRodrigo)

	utxoLeander := utxos.FindUTXO("leander")
	assert.Equal(t, []TXOutput(nil), utxoLeander)

	// Rodrigo sent 5 "coins" to Leander
	utxos = getTestExpectedUTXOSet("block1")[0]

	utxoRodrigo = utxos.FindUTXO("rodrigo")
	assert.Equal(t, []TXOutput{TXOutput{5, "rodrigo"}}, utxoRodrigo)

	utxoLeander = utxos.FindUTXO("leander")
	assert.Equal(t, []TXOutput{TXOutput{5, "leander"}}, utxoLeander)

	// Rodrigo sent 1 "coin" to Leander and
	// Leander sent 3 "coins" to Rodrigo
	utxos = getTestExpectedUTXOSet("block2")[0]

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
	utxos := getTestExpectedUTXOSet("genesis")[0]
	assert.Equal(t, 1, utxos.CountUTXOs())

	utxos = getTestExpectedUTXOSet("block1")[0]
	assert.Equal(t, 2, utxos.CountUTXOs())

	utxos = getTestExpectedUTXOSet("block2")[0]
	assert.Equal(t, 4, utxos.CountUTXOs())
}
func TestUpdate(t *testing.T) {
	for k, m := range testUTXOs {
		utxos := m.utxos
		switch k {
		case "genesis":
			utxos.Update([]*Transaction{testTransactions["tx0"]})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block1":
			utxos.Update([]*Transaction{testTransactions["tx1"]})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block2":
			utxos.Update([]*Transaction{testTransactions["tx2"], testTransactions["tx3"]})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block3":
			utxos.Update([]*Transaction{testTransactions["tx4"]})
			assert.True(t, utxos.Equal(m.expectedUTXOs[0]))
		case "block4":
			utxos.Update([]*Transaction{testTransactions["tx5"]})
			ok1 := utxos.Equal(m.expectedUTXOs[0])
			ok2 := utxos.Equal(m.expectedUTXOs[1])
			assert.True(t, ok1 || ok2)
		}
	}
}
