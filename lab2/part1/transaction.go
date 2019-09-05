package main

import (
	"fmt"
	"strings"
)

// Transaction represents a Bitcoin transaction
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// Serialize returns a serialized Transaction
func (tx Transaction) Serialize() []byte {
	// TODO(student)
	// Encode the Transaction struct
	// TIP: https://golang.org/pkg/encoding/gob/
	return []byte{}
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	// TODO(student)
	// Hash the serialized representation of a transaction
	return []byte{}
}

// String returns a human-readable representation of a transaction
func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))

	for i, input := range tx.Vin {
		lines = append(lines, fmt.Sprintf("     Input %d:", i))
		lines = append(lines, fmt.Sprintf("       TXID:      %x", input.Txid))
		lines = append(lines, fmt.Sprintf("       OutIdx:    %d", input.OutIdx))
		lines = append(lines, fmt.Sprintf("       ScriptSig: %s", input.ScriptSig))
	}

	for i, output := range tx.Vout {
		lines = append(lines, fmt.Sprintf("     Output %d:", i))
		lines = append(lines, fmt.Sprintf("       Value:  %d", output.Value))
		lines = append(lines, fmt.Sprintf("       ScriptPubKey: %s", output.ScriptPubKey))
	}

	return strings.Join(lines, "\n")
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	// TODO(student)
	// TIP: What differentiate a coinbase transaction from a normal transaction?
	// Remember that OutIdx represents the position of an output referred by the input
	return true
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	// TODO(student)
	// Create a new coinbase using the given data field
	// or the default "Reward to ADDRESS" in case of
	// data was empty.
	// ADDRESS represents the miner address
	return nil
}

// NewUTXOTransaction creates a new UTXO transaction
func NewUTXOTransaction(from, to string, amount int, utxos UTXOSet) (*Transaction, error) {
	// TODO(student)
	// 1) Find valid spendable outputs and the current balance of the sender
	// 2) The sender has sufficient funds? If not return the error:
	// "Not enough funds"
	// 3) Build a list of inputs based on the current valid outputs
	// 4) Build a list of new outputs, creating a "change" output if necessary
	// 5) Create a new transaction with the input and output list.
	return nil, nil
}
