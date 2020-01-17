// Copyright 2020 Luca Polverini

package blockchain

import (
	"bytes"

	"github.com/lp74/uc74-blockchain-go/wallet"
)

// TxOutput
// 	-	Value: il valore del TXO
// 	-	PubKeyHash: l'hash della Chiave Pubblica del soggetto destinatario
// 		in realta questa è una semplificazione; in Bitcoin questo campo è sostituito da ScriptPubKey
type TxOutput struct {
	Value      int
	PubKeyHash []byte // dovrebbe essere scriptPubKey, qui è la PubKeyHash del soggetto destinatario
}

// TxInput
// 	-	ID: referenzia una Transazione precedente (escluso TxInput Coinbase)
// 	-	Out: indice della transazione di uscita TxOutput della transazione referenziata
// 	-	PubKey: Chiave Pubblica del soggetto che emette la transazione
// 		in realta questa è una semplificazione; in Bitcoin questo campo è sostituito da ScriptSig
// 	-	Signature: la firma dell'HASH della transazione fatta a mezzo della chiave privata di colui che trasferisce
//		L'algoritmo usato in questo codice per firmare è ECDSA
// 		[Elliptic Curve Digital Signature Algorithm](https://en.bitcoin.it/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
type TxInput struct {
	ID        []byte // potrebbe essere chiamato prevTxID
	Out       int    // potrebbe essere chiamato Index
	PubKey    []byte // dovrebbe essere ScriptSig, qui è la chiave pubblica PubKeyHash del TxOutput utilizzato, dunque detenuto dal soggetto emittente
	Signature []byte
}

// Nota:
// l'ordine con i quali sono dichiarati TxOutput e TxInput non è casuale.
// Eccezione fatta per la transazione Coinbase, un TxInput non può esistere senza un TxOutput spendibile.
// Conseguentemente il TxOutput, concettualmente, precede il TxInput

// ---------------------         ---------------------
//    Transazione X          	    Transazione Y
// ---------------------         ---------------------
// TxInput   |0:TxOutput   <--|  TxInput  |0:TxOutput
//           |1:TxOutput      |  ID: X    |
// ---------------------      |  Out: 0   |
//                               --------------------

// Una transazione, tramite i TxInput può referenziare n (spendibili) TXOutput contenuti in m Transazioni

// Nota:
// scriptPubKey è il predicato (contenuto in TxOuput)
// scriptSig aiuta a soddisfare il predicato (contenuto in TxInput)
// Per comprendere è sufficiente questa domanda
// Domanda: Quando puoi spendere un TxOutput?
// Risposta: Quando conosci la sciptSig
// Questa parte sarà analizzata più avanti.

// UsesKey veriies the PubKey (Hash) of the TXInput transaction
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
}

// Lock blocca il TxOutput con la PubKey (HASH) del destinatario
// questo metodo di TxOutput, dato un indirizzo Bitcoin, ne ricava il PubKeyHash
// notare che dalla decodifica dell'indirizzo Base58 sono rimossi la versione e il checksum (1° e ultimi 4 byte della decodifica)
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
