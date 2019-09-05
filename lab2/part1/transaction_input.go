package main

// TXInput represents a transaction input
type TXInput struct {
	Txid      []byte
	OutIdx    int
	ScriptSig string
}

// CanUnlockOutputWith checks whether the address initiated the transaction
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	// TODO(student)
	return true
}
