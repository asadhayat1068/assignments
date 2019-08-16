package main

import (
	"fmt"
)

// MerkleTree represents a Merkle tree
type MerkleTree struct {
	RootNode *Node
	Leafs    []*Node
}

// Node represents a Merkle tree node
type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Hash   []byte
}

const (
	leftNode = iota
	rightNode
)

// MerkleProof represents a merkle path required to prove element inclusion
// on the merkle tree
type MerkleProof struct {
	path  [][]byte
	index []int64
}

// NewMerkleTree creates a new Merkle tree from a sequence of data
func NewMerkleTree(data [][]byte) *MerkleTree {
	// TODO(student)
	return nil
}

// NewMerkleNode creates a new Merkle tree node
func NewMerkleNode(left, right *Node, data []byte) *Node {
	// TODO(student)
	return &Node{}
}

// MerkleRootHash return the hash of the merkle root
func (mt *MerkleTree) MerkleRootHash() []byte {
	return mt.RootNode.Hash
}

// MerklePath returns a list of hashes and indexes required to
// reconstruct the inclusion proof of a given hash
//
// @param hash represents the hashed data (e.g. transaction ID) stored on
// the leaf node
// @return the merkle path (list of hashes), a list of indexes indicating
// the node location (leftNode or rightNode), and a possible error.
func (mt *MerkleTree) MerklePath(hash []byte) ([][]byte, []int64, error) {
	// TODO(student)
	return nil, nil, fmt.Errorf("Node %x not found", hash)
}

// VerifyProof verify that the correct root hash can be retrieved by
// properly hashing the given hash along with the merkle path in the
// correct order
//
// @param rootHash is the hash of the current root of the merkle tree
// @param hash represents the hash of the data (e.g. transaction ID)
// to be verified
// @param mProof is the merkle proof that contains the list of hashes and
// their indexes required to reconstruct the root hash
func VerifyProof(rootHash []byte, hash []byte, mProof MerkleProof) bool {
	// TODO(student)
	return false
}
