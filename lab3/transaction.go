package main

import (
	"crypto/ecdsa"
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
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return []byte{}
}

// Hash returns the hash of the Transaction
func (tx *Transaction) Hash() []byte {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
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
		lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("       PubKey: %x", input.PubKey))
	}

	for i, output := range tx.Vout {
		lines = append(lines, fmt.Sprintf("     Output %d:", i))
		lines = append(lines, fmt.Sprintf("       Value:  %d", output.Value))
		lines = append(lines, fmt.Sprintf("       PubKeyHash: %x", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}

// IsCoinbase checks whether the transaction is coinbase
func (tx Transaction) IsCoinbase() bool {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return true
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	return nil
}

// NewUTXOTransaction creates a new UTXO transaction
func NewUTXOTransaction(wallet *Wallet, to string, amount int, utxos UTXOSet, bc *Blockchain) (*Transaction, error) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// Modify your function to use the address instead of just strings
	// And also sign the new transaction before return
	return nil, nil
}

// TrimmedCopy creates a trimmed copy of Transaction to be used in signing
func (tx Transaction) TrimmedCopy() Transaction {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// You need to create a copy of the transaction to be signed.
	// The fields Signature and PubKey of the input need to be nil
	// since they are not included in signature.
	return Transaction{}
}

// Sign signs each input of a Transaction
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]*Transaction) {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// 1) coinbase transactions are not signed.
	// 2) Throw a Panic in case of any prevTXs (used inputs) didn't exists
	// Take a look on the tests to see the expected error message
	// 3) Create a copy of the transaction to be signed
	// 4) Sign all the previous TXInputs of the transaction tx using the
	// copy as the payload (serialized) to be signed in the ecdsa.Sig
	// (https://golang.org/pkg/crypto/ecdsa/#Sign)
	// Make sure that each input of the copy to be signed
	// have the correct PubKeyHash of each output in the prevTXs
	// Store the signature as a concatenation of R and S fields
}

// Verify verifies signatures of Transaction inputs
func (tx Transaction) Verify(prevTXs map[string]*Transaction) bool {
	// TODO(student) -- YOU DON'T NEED TO CHANGE YOUR PREVIOUS METHOD
	// 1) coinbase transactions are not signed.
	// 2) Throw a Panic in case of any prevTXs (used inputs) didn't exists
	// Take a look on the tests to see the expected error message
	// 3) Create the same copy of the transaction that was signed
	// and get the curve used for sign: P256
	// 4) Doing the opposite operation of the signing, perform the
	// verification of the signature, by recovering the R and S byte fields
	// of the Signature and the X and Y fields of the PubKey from
	// the inputs of tx. Verify the signature of each input using the
	// ecdsa.Verify function (https://golang.org/pkg/crypto/ecdsa/#Verify)
	// Note that to use this function you need to reconstruct the
	// ecdsa.PublicKey. Also notice that the ecdsa.Verify function receive
	// a byte array, you the transaction copy need to be serialized.
	// return true if all inputs have valid signature,
	// and false if any of them have an invalid signature.
	return true
}
