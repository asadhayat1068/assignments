package main

import (
	"encoding/hex"
)

// HexSlice2ByteSlice returns a slice of hex string hashes as byte slice.
func HexSlice2ByteSlice(str []string) [][]byte {
	var slice [][]byte
	for _, s := range str {
		slice = append(slice, Hex2Bytes(s))
	}
	return slice
}

// Hex2Bytes returns a hex string hash as byte slice.
func Hex2Bytes(str string) []byte {
	h, _ := hex.DecodeString(str)
	return h
}
