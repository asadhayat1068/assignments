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

// IntToHex converts an int64 to a byte array
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
