package wallet

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"log"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

// Wallet struttura atta a gestire la chiave privata
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// Address restituisce un indirizzo per la Block chain
// TODO: migliorare l'aderenza a Bitcoin
func (w Wallet) Address() []byte {
	ripemd160 := PublicKeyHash(w.PublicKey)

	versionedRimpemd160 := append([]byte{version}, ripemd160...)
	checksum := CheckSumSlice(versionedRimpemd160)

	fullHash := append(versionedRimpemd160, checksum...)
	address := Base58Encode(fullHash)

	return address
}

func p256Strategy() elliptic.Curve {
	return elliptic.P256()
}

func secp256k1Strategy() elliptic.Curve {
	return btcec.S256()
}

// EllipticCurve returns the elliptic curve of choice
func EllipticCurve() elliptic.Curve {
	return p256Strategy()
}

// NewKeyPair genera e restituisce una coppia:
// *chiave privata
// chiave pubblica
// utilizzando ECDSA
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	ellipticCurve := EllipticCurve()

	privateKey, err := ecdsa.GenerateKey(ellipticCurve, rand.Reader)
	if err != nil {
		log.Panic(err)
	}
	X := privateKey.PublicKey.X.Bytes()
	Y := privateKey.PublicKey.Y.Bytes()
	//fmt.Println(len(X), X)
	//fmt.Println(len(Y), Y)
	publicKey := append(
		X,    // 32 bytes (P256)
		Y..., // 32 bytes (P256)
	) // 64 bytes => 64 * 8 bits = 512 bits (perchè usiamo P256 o secp256k)
	return *privateKey, publicKey
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}

	return &wallet
}

func PublicKeyHash(pubKey []byte) []byte {
	pubHash := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHash[:])
	if err != nil {
		log.Panic(err)
	}

	publicRipMD := hasher.Sum(nil)

	return publicRipMD
}

// CheckSumSlice calcola il checksum
// restituendono i primi checksumLenght byte
func CheckSumSlice(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	return secondHash[:checksumLength]
}

func ValidateAddress(address string) bool {
	pubKeyHash := Base58Decode([]byte(address))
	actualChecksum := pubKeyHash[len(pubKeyHash)-checksumLength:]           // checksum del pubKeyHash: ultimi 4 byte
	version := pubKeyHash[0]                                                // la versione è il primo byte
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-checksumLength]             // la chiave sta in mezzo
	targetChecksum := CheckSumSlice(append([]byte{version}, pubKeyHash...)) // computa il target prendendo la versione e la chiave interna

	return bytes.Compare(actualChecksum, targetChecksum) == 0 // compara
}
