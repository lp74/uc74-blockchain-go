package cec

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"

	"github.com/btcsuite/btcd/btcec"
)

/*
In quali file serve la firma:

Wallet:
il wallet ha una chiave pubblice ed una privata
- qui viene creata

Wallets
- qui viene memorizzata (wallets)

Transazioni:
- verifica della firma della transazione
occorre sia la curva, per ricavare la chiave pubblica da x e y
sia la verifica

-


*/
type securityCurve func() elliptic.Curve

// PrivateKey wraps ecdsa.PrivateKey
type PrivateKey ecdsa.PrivateKey

// PublicKey wraps ecdsa.PublicKey
type PublicKey ecdsa.PublicKey

// CEC Chain Elliptic Curve
type CEC struct {
	// privati
	curve      elliptic.Curve
	privateKey *ecdsa.PrivateKey

	// pubblici
	PublicKey *ecdsa.PublicKey
}

// NewSec costruisce un Sec
func NewSec(fn securityCurve) *CEC {
	sec := CEC{
		curve: fn(),
	}
	return &sec
}

// NewKeyPair generate new key pair
func (sec *CEC) newKey() {
	curve := btcec.S256()
	pKey, _ := btcec.NewPrivateKey(curve)
	sec.privateKey = pKey.ToECDSA()
}

func (sec *CEC) verify(publicKey PublicKey, signature []byte) {
	fmt.Println("Should verify the signature")
}

func (sec *CEC) sign(privateKey PrivateKey, data []byte) {
	fmt.Println("Should sign the data")
}
