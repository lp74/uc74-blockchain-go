package blockchain

import (
	"bytes"

	"github.com/lp74/uc74-blockchain-go/wallet"
)

// TxOutput si compone di un valore e della chiave pubblica del destinatario
type TxOutput struct {
	Value      int
	PubKeyHash []byte
}

// TxInput in questa implementazione una transazione di input si compone di
// ID in riferimento a transazioni precedenti
// Out il riferimento alla transazione di uscita
// una stringa arbitraria
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

// UsesKey veriies the PubKey of the TXInput transaction
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// Lock locks the TXOutput with the PubKey of the owner
func (out *TxOutput) Lock(address []byte) {
	pubKeyHash := wallet.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey given the PubKey hash checks the TXO ownership
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput returns a new TXO of a given amount and locked by the owner
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
