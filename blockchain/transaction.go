package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/lp74/uc74-blockchain-go/wallet"
)

// Transaction una transazione è composta da due aggregati:
// il riferimento a transazioni precedenti
// le transazioni in ingresso (TxInput) e le transazioni in uscita (TxOutput)
type Transaction struct {
	ID   []byte
	Vin  []CTxIn
	Vout []CTxOut
}

// Hash hash della transazione
// serializza la transazione e ne restituisce l'hash SHA-256
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

// Serialize serializza la transazione
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}

	return encoded.Bytes()
}

// DeserializeTransaction de-serializza la transazione
func DeserializeTransaction(data []byte) Transaction {
	var transaction Transaction

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&transaction)
	Handle(err)
	return transaction
}

// CoinbaseTx la transazione Coinbase è la prima transazione della catena.
// Ne viene aggiunta una ad ogni blocco e rappresenta l'incentivo destinato a chi ha formato il blocco
// è una transazione speciale perché non necessità di referenziare nessuna transazione precedente.
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		randData := make([]byte, 24)
		_, err := rand.Read(randData)
		Handle(err)
		data = fmt.Sprintf("%x", randData)
	}

	txin := CTxIn{
		PrevTxID:  []byte{},
		OutIndex:  -1,
		PubKey:    []byte(data),
		Signature: nil,
	}
	txout := NewTXOutput(20, to)

	tx := Transaction{
		ID:   nil,
		Vin:  []CTxIn{txin},
		Vout: []CTxOut{*txout},
	}
	tx.ID = tx.Hash()

	return &tx
}

// NewTransaction genera una nuova transazione
//
// - from: = indirizzo sorgente
// - to: indirizzo destinatario
// - amount = valore
// - chain: riferimento alla block chain
//
// * Recupera il wallet e preleva l'indirizzo del soggetto emittente
// * solo il soggetto emittente detiene la chiave privata all'interno del wallet
// * crea il pubKeyHash a partire dalla chiave pubblica [REV KEY CHECKSUM]
// * cerca gli UTXO necessari per spendere il valore amout attraverso la pubKeyHash e ne computa il valore totale (se inferiore al necessario termina)
// * itera gli UTXO. Gli UTXO sono raggruppati per transazione, dunque spendableOutputs è uma mappa chiave valore { "TXID" : [ TXO ] }
// * genera gli input della transazione
// * genera gli output della transazione ponendo il PubKeyHash del soggetto destinatario
// * firma la transazione
func NewTransaction(w *wallet.Wallet, to string, amount int, UTXO *UTXOSet) *Transaction {
	var vin []CTxIn
	var vout []CTxOut

	pubKeyHash := wallet.PublicKeyHash(w.PublicKey)
	acc, validOutputs := UTXO.FindSpendableOutputs(pubKeyHash, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			txIn := CTxIn{
				PrevTxID:  txID,
				OutIndex:  out,
				PubKey:    w.PublicKey,
				Signature: nil,
			}
			vin = append(vin, txIn)
		}
	}

	from := fmt.Sprintf("%s", w.Address())

	vout = append(vout, *NewTXOutput(amount, to))

	if acc > amount {
		vout = append(vout, *NewTXOutput(acc-amount, from))
	}

	tx := Transaction{
		ID:   nil,
		Vin:  vin,
		Vout: vout,
	}
	tx.ID = tx.Hash()
	UTXO.Blockchain.SignTransaction(&tx, w.PrivateKey)

	return &tx
}

// IsCoinbase determina se una transazione è la transazione Coinbase
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].PrevTxID) == 0 && tx.Vin[0].OutIndex == -1
}

// Sign firma la transazione
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.IsCoinbase() {
		return
	}

	for _, in := range tx.Vin {
		if prevTXs[hex.EncodeToString(in.PrevTxID)].ID == nil {
			log.Panic("ERROR: Previous transaction is not correct")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inID, in := range txCopy.Vin {
		prevTX := prevTXs[hex.EncodeToString(in.PrevTxID)]
		txCopy.Vin[inID].Signature = nil
		txCopy.Vin[inID].PubKey = prevTX.Vout[in.OutIndex].PubKeyHash

		dataToSign := fmt.Sprintf("%x\n", txCopy)

		r, s, err := ecdsa.Sign(rand.Reader, &privKey, []byte(dataToSign))
		Handle(err)
		signature := append(r.Bytes(), s.Bytes()...)

		tx.Vin[inID].Signature = signature
		txCopy.Vin[inID].PubKey = nil
	}
}

// Verify verifica la transazione
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}

	for _, in := range tx.Vin {
		if prevTXs[hex.EncodeToString(in.PrevTxID)].ID == nil {
			log.Panic("Previous transaction not correct")
		}
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	for inId, in := range tx.Vin {
		prevTx := prevTXs[hex.EncodeToString(in.PrevTxID)]
		txCopy.Vin[inId].Signature = nil
		txCopy.Vin[inId].PubKey = prevTx.Vout[in.OutIndex].PubKeyHash

		r := big.Int{}
		s := big.Int{}

		sigLen := len(in.Signature)
		r.SetBytes(in.Signature[:(sigLen / 2)])
		s.SetBytes(in.Signature[(sigLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(in.PubKey)
		x.SetBytes(in.PubKey[:(keyLen / 2)])
		y.SetBytes(in.PubKey[(keyLen / 2):])

		dataToVerify := fmt.Sprintf("%x\n", txCopy)

		rawPubKey := ecdsa.PublicKey{Curve: curve, X: &x, Y: &y}
		if ecdsa.Verify(&rawPubKey, []byte(dataToVerify), &r, &s) == false {
			return false
		}
		txCopy.Vin[inId].PubKey = nil
	}

	return true
}

// TrimmedCopy TODO: clarify
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []CTxIn
	var outputs []CTxOut

	for _, in := range tx.Vin {
		inputs = append(inputs, CTxIn{
			PrevTxID:  in.PrevTxID,
			OutIndex:  in.OutIndex,
			PubKey:    nil,
			Signature: nil,
		})
	}

	for _, out := range tx.Vout {
		outputs = append(outputs, CTxOut{out.Value, out.PubKeyHash})
	}

	txCopy := Transaction{
		ID:   tx.ID,
		Vin:  inputs,
		Vout: outputs,
	}

	return txCopy
}

// String restituisce una rappresentazione testuale della transazione
func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("--- Transaction %x:", tx.ID))
	for i, input := range tx.Vin {
		lines = append(lines, fmt.Sprintf("     Input %d:", i))
		lines = append(lines, fmt.Sprintf("       TXID:     %x", input.PrevTxID))
		lines = append(lines, fmt.Sprintf("       Out:       %d", input.OutIndex))
		lines = append(lines, fmt.Sprintf("       Signature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("       PubKey:    %x", input.PubKey))
	}

	for i, output := range tx.Vout {
		lines = append(lines, fmt.Sprintf("     Output %d:", i))
		lines = append(lines, fmt.Sprintf("       Value:  %d", output.Value))
		lines = append(lines, fmt.Sprintf("       Script: %x", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}
