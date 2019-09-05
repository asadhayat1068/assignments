package main

// UTXOSet represents a set of UTXO as an in-memory cache
// The key of the map is the transaction ID
// (i.e., encoded as string) which contains these outputs
type UTXOSet map[string][]TXOutput

// FindSpendableOutputs finds and returns unspent outputs in the UTXO Set
// to reference in inputs and the current accumulated balance
func (u UTXOSet) FindSpendableOutputs(unlockingData string, amount int) (int, map[string][]int) {
	// TODO(student)
	return 0, make(map[string][]int)
}

// FindUTXO finds UTXO in the UTXO Set for a given unlockingData key (e.g., address)
func (u UTXOSet) FindUTXO(unlockingData string) []TXOutput {
	var UTXO []TXOutput
	// TODO(student)
	// Search for UTXO that unlockingData can unlock
	return UTXO
}

// CountUTXOs returns the number of transactions outputs in the UTXO set
func (u UTXOSet) CountUTXOs() int {
	// TODO(student)
	return 0
}

// Update updates the UTXO Set with the new set of transactions
func (u UTXOSet) Update(transactions []*Transaction) {
	// TODO(student)
	// Iterate over the transactions  and update
	// the current UTXOSet with the new
	// transactions.
	//
	// TIP: Remember to remove a entry from the UTXOSet
	// in case that it was fully spent
}

// Equal compares two UTXOSet
func (u UTXOSet) Equal(utxos UTXOSet) bool {
	if len(u) != len(utxos) {
		return false
	}

	for txid, outs := range u {
		o, ok := utxos[txid]

		if !ok || len(outs) != len(o) {
			return false
		}

		for i, out := range outs {
			if out != o[i] {
				return false
			}
		}
	}

	return true
}
