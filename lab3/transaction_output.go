package main

import (
	"fmt"
)

// TXOutput represents a transaction output
type TXOutput struct {
	Value      int    // The amount
	PubKeyHash []byte // The hash of the public key (used to "lock" the output)
}

// Lock locks the transaction to a specific address
// Only this address owns this transaction
func (out *TXOutput) Lock(address string) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// "Lock" the TXOutput to a specific PubKeyHash
	// based on the given address
}

// IsLockedWithKey checks if the output can be used by the owner of the pubkey
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return true
}

// NewTXOutput create a new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// Create a new locked TXOutput
	return nil
}

func (out TXOutput) String() string {
	return fmt.Sprintf("{%d, %x}", out.Value, out.PubKeyHash)
}
