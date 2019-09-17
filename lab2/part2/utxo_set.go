package main

import (
	"bytes"
	"fmt"
	"strings"
)

// UTXOSet represents a set of UTXO as an in-memory cache
// The key of the map is the transaction ID
// (encoded as string) that contains these outputs
type UTXOSet map[string][]TXOutput

// FindSpendableOutputs finds and returns unspent outputs in the UTXO Set
// to reference in inputs
func (u UTXOSet) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	// TODO(student)
	// Modify your function to use IsLockedWithKey instead of CanBeUnlockedWith
	return 0, nil
}

// FindUTXO finds UTXO in the UTXO Set for a given unlockingData key (e.g., address)
func (u UTXOSet) FindUTXO(pubKeyHash []byte) []TXOutput {
	// TODO(student)
	// Modify your function to use IsLockedWithKey instead of CanBeUnlockedWith
	return []TXOutput{}
}

// CountUTXOs returns the number of transactions outputs in the UTXO set
func (u UTXOSet) CountUTXOs() int {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return 0
}

// Update updates the UTXO Set with the new set of transactions
func (u UTXOSet) Update(transactions []*Transaction) {
	// TODO(student)
	// Modify this function if needed to comply with the new
	// input and output struct.
}

// Equal compares two UTXOSet
func (u UTXOSet) Equal(utxos UTXOSet) bool {
	if len(u) != len(utxos) {
		return false
	}

	for txid, outputs := range u {
		o, ok := utxos[txid]

		if !ok || len(outputs) != len(o) {
			return false
		}

		for i, out := range outputs {
			if out.Value != o[i].Value || !bytes.Equal(out.PubKeyHash, o[i].PubKeyHash) {
				return false
			}
		}
	}

	return true
}

func (u UTXOSet) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- UTXO SET:"))
	for txid, outputs := range u {
		lines = append(lines, fmt.Sprintf("     TxID: %s", txid))
		for i, out := range outputs {
			lines = append(lines, fmt.Sprintf("           Output %d: %v", i, out))
		}
	}

	return strings.Join(lines, "\n")
}
