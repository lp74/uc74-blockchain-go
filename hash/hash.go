package hash

import (
	"crypto/sha256"
	"log"

	"golang.org/x/crypto/ripemd160"
)

// CHash256 una classe che genera l'hash Bitcoin's 256-bit
type CHash256 struct {
}

// Finalize restituisce l'hash 256 dei dati
func (CHash256) Finalize(data []byte) []byte {
	sha := sha256.Sum256(data)
	sha = sha256.Sum256(sha[:])
	return sha[:]
}

// CHash160 una classe che genera l'hash Bitcoin's 160-bit  (SHA-256 + RIPEMD-160).
type CHash160 struct {
}

// Finalize restituisce l'hash 160 dei dati
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

// Hash Computa l'hash 256-bit di un oggetto
func Hash(data []byte) []byte {
	return CHash256{}.Finalize(data)
}

// Hash160 Computa l'hash 160-bit di un oggetto
func Hash160(data []byte) []byte {
	return CHash256{}.Finalize(data)
}
