package blockchain

import (
	"bytes"

	"github.com/lp74/uc74-blockchain-go/wallet"
)

// TxInput 
// - ID: referenzia una Transazione precedente (escluso TxInput Coinbase) 
// - Out: il riferimento alla transazione di uscita che viene utilizzata
// - Signature: la firma della transazione fatta a mezzo della chiave privata di colui che trasferisce
// - PubKey: Chiave Pubblica del soggetto che emette la transazione
type TxInput struct {
	ID        []byte
	Out       int
	Signature []byte
	PubKey    []byte
}

// TxOutput
// - Value: il valore del TXO
// - PubKeyHash: l'hash della Chiave Pubblica del soggetto destinatario
type TxOutput struct {
	Value      int
	PubKeyHash []byte
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
