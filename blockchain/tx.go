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
// 	-	PubKey: Chiave Pubblica del soggetto che emette la transazione (deve combaciare con UTXO referenziato)
// 		in realta questa è una semplificazione; in Bitcoin questo campo è sostituito da ScriptSig
// 	-	Signature: la firma dell'HASH della transazione fatta a mezzo della chiave privata di colui che trasferisce
//		L'algoritmo usato in questo codice per firmare è ECDSA
// 		[Elliptic Curve Digital Signature Algorithm](https://en.bitcoin.it/wiki/Elliptic_Curve_Digital_Signature_Algorithm)
type TxInput struct {
	ID        []byte // potrebbe essere chiamato prevTxID
	Out       int    // potrebbe essere chiamato Index
	PubKey    []byte // dovrebbe essere ScriptSig <sig><pubKey>, qui è la chiave pubblica PubKey del soggetto emittente
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
// Maggiori informazioni sul linguaggio Script possono essere tovare [qui](https://en.bitcoin.it/wiki/Script)

/*
# Esempio di script

Questo esempio ci aiuta a comprendere due cose:
1. come funziona script
2. come processare e verificare un pagamento standard nel nostro sistema semplificato.

Come si vede dallo script sotto riportato, l'output è spendibile qualora la firma (della nuova transazione) possa esser verificata (check)
usando la chiave pubblica.
Rifrasando. Lo scriptPubKey richede di

1. duplicare la chiave pubblica del soggetto emittente
2. eseguire HASH160
3. aggiungere allo stack il pubKeyHash fornito dal scriptPubKey (equivale a TxOutput.Lock)
4. confrontare l'uguaglianza (equivale a TxOutput.IsLockedWithKey)
5. verificare della firma (Hash della transazione) con la chiave pubblica.

Poiché la chiave pubblica del soggetto emittente (posta in scriptSig):
- deve corrispondere per mezzo dell'HASH160 a quella dell'UTXO (quindi corrisponde a quella del soggetto destinatario dell'UTXO usato come input)
- deve poter verificare la firma della transazione (quindi la firma fatta per mezzo della chiave privata è verificabile permezzo della chiave pubblica)

e poiché la transazione corrente:
- è implicitamente contenuta nella firma per mezzo del suo HASH
- referenzia transazioni precedenti per mezzo dei loro HASH

La combinazione di tutti questi fattori mi intitola a spendere l'UTXO (sono il legittimo proprietario).


Naturalmente la catena delle transizioni non può essere riscritta senza riscrivere la catena.
Ciò richiederebbe un attacco svolto con almeno in 51% della capacità computazionale dell'intera rete Bitcoin

scriptPubKey:
P2PKH OP_DUP OP_HASH160 <pubKeyHash> OP_EQUALVERIFY OP_CHECKSIG

scriptSig:
<signature><pubKey>

=>

<signature><pubKey>

OP_DUP:
<signature><pubKey>
<pubKey>

OP_HASH160:
<signature><pubKey>
<pubKeyHashA>

<pubKeyHash>:
<signature><pubKey>
<pubKeyHashA>
<pubKeyHash>

OP_EQUALVERIFY:
<signature><pubKey>

OP_CHECKSIG:
true

*/

// UsesKey verifies the PubKey (Hash) of the TXInput transaction
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
func (txo *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(txo.PubKeyHash, pubKeyHash) == 0
}

// NewTXOutput returns a new TXO of a given amount and locked by the owner
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
