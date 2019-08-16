package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

var table = []struct {
	data             [][]byte
	expectedRootHash []byte
	expectedPath     [][]byte
	expectedIndex    []int64
}{
	{
		data: [][]byte{
			[]byte("node1"),
			[]byte("node2"),
			[]byte("node3"), // n4.Hash == n3.Hash => duplicated since data is odd
		},
		expectedRootHash: Hex2Bytes("4e3e44e55926330ab6c31892f980f8bfd1a6e910ff1ebc3f778211377f35227e"),
		expectedPath: HexSlice2ByteSlice([]string{
			"3b5bb1c6e7b76daba8afd89516e24140a67fc6be2ba071cc3b97d1b2e08c238d", "64b04b718d8b7c5b6fd17f7ec221945c034cfce3be4118da33244966150c4bd4",
		}), // Path for n3 => [n4, n5]
		expectedIndex: []int64{rightNode, leftNode},
	},
	{
		data: [][]byte{
			[]byte("node1"),
			[]byte("node2"),
			[]byte("node3"),
			[]byte("node4"),
			[]byte("node5"),
			[]byte("node6"),
			[]byte("node7"),
			[]byte("node8"),
		},
		expectedRootHash: Hex2Bytes("38c456cfef483f85c116a37a6c6f73791a91a53e2445533311ad5c54b1054226"),
		expectedPath: HexSlice2ByteSlice([]string{
			"d2b8f62a7e335bbd5576c8422844760f22ec378009eeea790c41e4dc45f23c33", "64b04b718d8b7c5b6fd17f7ec221945c034cfce3be4118da33244966150c4bd4",
			"4a3bef0c7511a5e0a652d37cb28c364df456605bb71e12846cf817fb9ddf8116",
		}), // Path for n3 => [n4, n9, n14]
		expectedIndex: []int64{rightNode, leftNode, rightNode},
	},
}

func TestMerkleTree(t *testing.T) {
	for i := 0; i < len(table); i++ {
		mTree := NewMerkleTree(table[i].data)
		if mTree == nil {
			t.Fatal("Got an error while creating the Merkle tree")
		}
		if bytes.Compare(mTree.MerkleRootHash(), table[i].expectedRootHash) != 0 {
			t.Errorf("error: expected hash equal to %x got %x", table[i].expectedRootHash, mTree.MerkleRootHash())
		}
	}
}

func TestMerklePath(t *testing.T) {
	for i := 0; i < len(table); i++ {
		mTree := NewMerkleTree(table[i].data)
		hash := sha256.Sum256(table[i].data[2]) // node3
		path, index, _ := mTree.MerklePath(hash[:])
		assert.Equal(t, table[i].expectedPath, path, "Merkle path is incorrect")
		assert.Equal(t, table[i].expectedIndex, index, "Merkle path index is incorrect")
	}
}

func TestMerkleProof(t *testing.T) {
	for i := 0; i < len(table); i++ {
		hash := sha256.Sum256(table[i].data[2]) // node3
		assert.True(t, VerifyProof(table[i].expectedRootHash, hash[:], MerkleProof{table[i].expectedPath, table[i].expectedIndex}), "Inclusion proof couldn't be satisfied")
	}
}

func TestNewMerkleTreeNoNodes(t *testing.T) {
	assert.PanicsWithValue(t, "No merkle tree nodes", func() { NewMerkleTree([][]byte{}) })
}

func TestMerklePathNodeNotFound(t *testing.T) {
	n := NewMerkleNode(nil, nil, []byte("other"))
	mTree := NewMerkleTree([][]byte{[]byte("node1")})
	if mTree == nil {
		t.Fatal("Got an error while creating the Merkle tree")
	}
	_, _, err := mTree.MerklePath(n.Hash)
	assert.Errorf(t, err, "Node %x not found", n.Hash)
}

func TestNewMerkleNode(t *testing.T) {
	data := [][]byte{
		[]byte("node1"),
		[]byte("node2"),
		[]byte("node3"),
	}

	// Level 1
	n1 := NewMerkleNode(nil, nil, data[0])
	n2 := NewMerkleNode(nil, nil, data[1])
	n3 := NewMerkleNode(nil, nil, data[2])
	n4 := NewMerkleNode(nil, nil, data[2])

	// Level 2
	n5 := NewMerkleNode(n1, n2, nil)
	n6 := NewMerkleNode(n3, n4, nil)

	// Level 3 (root)
	n7 := NewMerkleNode(n5, n6, nil)

	assert.Equal(
		t,
		"64b04b718d8b7c5b6fd17f7ec221945c034cfce3be4118da33244966150c4bd4",
		hex.EncodeToString(n5.Hash),
		"Level 1 hash 1 is incorrect",
	)
	assert.Equal(
		t,
		"08bd0d1426f87a78bfc2f0b13eccdf6f5b58dac6b37a7b9441c1a2fab415d76c",
		hex.EncodeToString(n6.Hash),
		"Level 1 hash 2 is incorrect",
	)
	assert.Equal(
		t,
		"4e3e44e55926330ab6c31892f980f8bfd1a6e910ff1ebc3f778211377f35227e",
		hex.EncodeToString(n7.Hash),
		"Root hash is incorrect",
	)
}
