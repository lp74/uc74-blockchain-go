package hash

import (
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

// CHash256 A hasher class for Bitcoin's 256-bit hash (double SHA-256).
type CHash256 struct {
}

// Finalize ritorna l'hash 256
func (CHash256) Finalize(data []byte) []byte {
	sha := sha256.Sum256(data)
	// hash = sha256.Sum256(hash[:])
	return sha[:]
}

// CHash160 A hasher class for Bitcoin's 160-bit hash (SHA-256 + RIPEMD-160).
type CHash160 struct {
}

// Finalize ritorna l'hash 160
func (CHash160) Finalize(data []byte) []byte {
	sha := CHash256{}.Finalize(data)

	hasher := ripemd160.New()
	_, err := hasher.Write(sha[:])
	if err != nil {
		log.Panic(err)
	}
	sha = hasher.Sum(nil)

	return sha[:]
}

// Hash Compute the 256-bit hash of an object.
func Hash(data []byte) []byte {
	return CHash256{}.Finalize(data)
}

// Hash160 Compute the 160-bit hash of an object.
func Hash160(data []byte) []byte {
	return CHash256{}.Finalize(data)
}
