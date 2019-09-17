package main

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte // The ID (i.e. Hash) of the transaction.
	OutIdx    int    // The index of the output
	Signature []byte // The signature of this input.
	PubKey    []byte // The raw public key (not hashed)
}

// UsesKey checks whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	// TODO(student)
	// Check if the given lockingHash came from
	// the same PubKey of the input
	return true
}
