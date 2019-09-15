package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newMockBlockchain(address string) *Blockchain {
	genesis := newMockBlock([]*Transaction{testTransactions["tx0"]}, []byte{})
	bc := &Blockchain{[]*Block{genesis}}
	return bc
}

func addMockBlock(bc *Blockchain, newBlock *Block) {
	bc.blocks = append(bc.blocks, newBlock)
}

func TestBlockchain(t *testing.T) {
	// NewBlockchain
	bc := NewBlockchain("rodrigo")
	assert.NotNil(t, bc)
	assert.Equal(t, 1, len(bc.blocks))

	// GetGenesisBlock
	gb := bc.GetGenesisBlock()
	assert.NotNil(t, gb)
	assert.Equal(t, []byte{}, gb.PrevBlockHash, "Genesis block shouldn't has PrevBlockHash")

	// Genesis block should contains a genesis transaction
	coinbaseTx := gb.Transactions[0]
	assert.Equal(t, 1, len(gb.Transactions))
	assert.Equal(t, -1, coinbaseTx.Vin[0].OutIdx)
	assert.Equal(t, []byte{}, coinbaseTx.Vin[0].Txid)
	assert.Equal(t, GenesisCoinbaseData, coinbaseTx.Vin[0].ScriptSig)
	assert.Equal(t, BlockReward, coinbaseTx.Vout[0].Value)
	assert.Equal(t, "rodrigo", coinbaseTx.Vout[0].ScriptPubKey)
}

func TestAddBlock(t *testing.T) {
	bc := newMockBlockchain("rodrigo")
	assert.Equal(t, 1, len(bc.blocks))

	b1 := bc.AddBlock([]*Transaction{testTransactions["tx1"]})
	assert.NotNil(t, b1)
	assert.Equal(t, 2, len(bc.blocks))

	gb := bc.blocks[0]
	assert.Equalf(t, gb.Hash, b1.PrevBlockHash, "Genesis block Hash: %x isn't equal to current PrevBlockHash: %x", gb.Hash, b1.PrevBlockHash)

	b2 := bc.AddBlock([]*Transaction{testTransactions["tx2"]})
	assert.NotNil(t, b2)
	assert.Equal(t, 3, len(bc.blocks))
	assert.Equalf(t, b1.Hash, b2.PrevBlockHash, "Previous block Hash: %x isn't equal to the expected: %x", b2.PrevBlockHash, b1.Hash)
}

func TestCurrentBlock(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

	b := bc.CurrentBlock()
	assert.NotNil(t, b)
	expectedBlock := bc.blocks[0]
	assert.Equalf(t, expectedBlock.Hash, b.Hash, "Current block Hash: %x isn't the expected: %x", b.Hash, expectedBlock.Hash)

	addMockBlock(bc, newMockBlock([]*Transaction{testTransactions["tx1"]}, bc.blocks[0].Hash))

	b = bc.CurrentBlock()
	assert.NotNil(t, b)
	expectedBlock = bc.blocks[1]
	assert.Equalf(t, expectedBlock.Hash, b.Hash, "Current block Hash: %x isn't the expected: %x", b.Hash, expectedBlock.Hash)
}

func TestGetBlock(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

	b, err := bc.GetBlock(bc.blocks[0].Hash)
	assert.Nil(t, err)
	assert.NotNil(t, b)

	assert.Equalf(t, bc.blocks[0].Hash, b.Hash, "Block Hash: %x isn't the expected: %x", b.Hash, bc.blocks[0].Hash)
}

func TestMineBlockWithInvalidTxInput(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

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
	bc := newMockBlockchain("rodrigo")

	b, err := bc.MineBlock([]*Transaction{testTransactions["tx1"]})
	assert.Nil(t, err)
	assert.NotNil(t, b)
	assert.Equal(t, 2, len(bc.blocks))
	gb := bc.blocks[0]
	assert.Equalf(t, gb.Hash, b.PrevBlockHash, "Genesis block Hash: %x isn't equal to current PrevBlockHash: %x", gb.Hash, b.PrevBlockHash)

	minedBlock, err := bc.GetBlock(b.Hash)
	assert.Equal(t, b, minedBlock)

	txMinedBlock := bc.blocks[1].Transactions[0]
	assert.NotNil(t, txMinedBlock)
	assert.Equal(t, testTransactions["tx1"].ID, txMinedBlock.ID)
}

func TestVerifyTransaction(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

	assert.True(t, bc.VerifyTransaction(testTransactions["tx0"]))
	assert.True(t, bc.VerifyTransaction(testTransactions["tx1"]))
}

func TestVerifyTransactionInvalidTxInput(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

	tx := &Transaction{
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
	assert.False(t, bc.VerifyTransaction(tx))
}

func TestFindTransaction(t *testing.T) {
	bc := newMockBlockchain("rodrigo")

	// Find coinbase genesis transaction
	tx, err := bc.FindTransaction(Hex2Bytes("e2404638779673c7c3e772e12dc3343e6d38f1d71625419d12a8468522b5cc2d"))
	assert.Nil(t, err)
	assert.NotNil(t, tx)

	notFoundTx, err := bc.FindTransaction(Hex2Bytes("non-existentID"))
	assert.Error(t, err, "Transaction not found")
	assert.Nil(t, notFoundTx)
}

func TestFindUTXOSet(t *testing.T) {
	bc := newMockBlockchain("rodrigo")
	expectedUTXOs := getTestExpectedUTXOSet("genesis")[0]

	utxos := bc.FindUTXOSet()
	assert.Equal(t, expectedUTXOs, utxos)

	addMockBlock(bc, newMockBlock([]*Transaction{testTransactions["tx1"]}, bc.blocks[0].Hash))
	expectedUTXOs = getTestExpectedUTXOSet("block1")[0]

	utxos = bc.FindUTXOSet()
	assert.Equal(t, expectedUTXOs, utxos)
}
