package main

import (
	"math"
	"math/big"
)

var (
	maxNonce = math.MaxInt64
)

// TARGETBITS define the mining difficulty
const TARGETBITS = 8

// ProofOfWork represents a block mined with a target difficulty
type ProofOfWork struct {
	block  *Block
	target *big.Int // 2 ** (256 - targetbits)
}

// NewProofOfWork builds a ProofOfWork
func NewProofOfWork(block *Block) *ProofOfWork {
	// TODO(student)
	return &ProofOfWork{}
}

// setupHeader prepare the header of the block
func (pow *ProofOfWork) setupHeader() []byte {
	// TODO(student)
	return []byte{}
}

// addNonce adds a nonce to the header
func addNonce(nonce int, header []byte) []byte {
	// TODO(student)
	return []byte{}
}

// Run performs the proof-of-work
func (pow *ProofOfWork) Run() (int, []byte) {
	// TODO(student)
	return 0, []byte{}
}

// Validate validates block's Proof-Of-Work
// This function just validates if the block header hash
// is less than the target.
func (pow *ProofOfWork) Validate() bool {
	// TODO(student)
	return true
}
