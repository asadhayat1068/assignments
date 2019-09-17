package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
)

const version = byte(0x00)
const addressChecksumLen = 4

// Wallet stores private and public keys
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewWallet creates and returns a new Wallet
func NewWallet() *Wallet {
	// TODO(student)
	// Generate a new key pair and create a wallet
	return nil
}

// CreateWallet initialize a wallet from the given keys
func CreateWallet(privKey *ecdsa.PrivateKey, pubKey *ecdsa.PublicKey) *Wallet {
	// TODO(student)
	// Create a wallet with the given keys, note that the PublicKey field in the
	// Wallet struct is a byte array (the concatenation of the X and Y
	// coordinates of the ecdsa.PublicKey) this is done to be easy to hash it
	// in future operations
	return nil
}

// GetAddress returns wallet address
// https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses#How_to_create_Bitcoin_Address
func (w Wallet) GetAddress() []byte {
	// TODO(student)
	// Create a address following the logic described in the link above and
	// in the lab documentation
	return []byte{}
}

// GetStringAddress returns wallet address as string
func (w Wallet) GetStringAddress() string {
	return string(w.GetAddress())
}

// HashPubKey hashes public key
func HashPubKey(pubKey []byte) []byte {
	// TODO(student)
	// compute the SHA256 + RIPEMD160 hash of the pubkey
	// step 2 and 3 of: https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses#How_to_create_Bitcoin_Address
	// use the go package ripemd160:
	// https://godoc.org/golang.org/x/crypto/ripemd160
	return []byte{}
}

// GetPubKeyHashFromAddress returns the hash of the public key
// discarding the version and the checksum
func GetPubKeyHashFromAddress(address string) []byte {
	// TODO(student)
	// Decode the address using Base58Decode and extract the hash of the pubkey
	// Look in the picture of the documentation of the lab to understand
	// how it is stored: version + pubkeyhash + checksum
	return []byte{}
}

// ValidateAddress check if an address is valid
func ValidateAddress(address string) bool {
	// TODO(student)
	// Validate a address by decoding it, extracting the
	// checksum, re-computing it using the "checksum" function
	// and comparing both.
	return true
}

// Checksum generates a checksum for a public key
func checksum(payload []byte) []byte {
	// TODO(student)
	// Perform a double sha256 on the versioned payload
	// and return the first 4 bytes
	// Steps 5,6, and 7 of: https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses#How_to_create_Bitcoin_Address
	return []byte{}
}

func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// TODO(student)
	// Create a new cryptographic key pair using
	// the "elliptic" and "ecdsa" package.
	// Additionally, convert the PublicKey to byte
	pubKey := []byte{}
	return ecdsa.PrivateKey{}, pubKey
}

func pubKeyToByte(pubkey ecdsa.PublicKey) []byte {
	// TODO(student)
	// step 1 of: https://en.bitcoin.it/wiki/Technical_background_of_version_1_Bitcoin_addresses#How_to_create_Bitcoin_Address
	return []byte{}
}

func encodeKeyPair(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string) {
	return encodePrivateKey(privateKey), encodePublicKey(publicKey)
}

func encodePrivateKey(privateKey *ecdsa.PrivateKey) string {
	x509Encoded, _ := x509.MarshalECPrivateKey(privateKey)
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: x509Encoded})

	return string(pemEncoded)
}

func encodePublicKey(publicKey *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncodedPub)
}

func decodeKeyPair(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, *ecdsa.PublicKey) {
	return decodePrivateKey(pemEncoded), decodePublicKey(pemEncodedPub)
}

func decodePrivateKey(pemEncoded string) *ecdsa.PrivateKey {
	block, _ := pem.Decode([]byte(pemEncoded))
	privateKey, _ := x509.ParseECPrivateKey(block.Bytes)

	return privateKey
}

func decodePublicKey(pemEncodedPub string) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	genericPubKey, _ := x509.ParsePKIXPublicKey(blockPub.Bytes)
	publicKey := genericPubKey.(*ecdsa.PublicKey) // cast to ecdsa

	return publicKey
}
