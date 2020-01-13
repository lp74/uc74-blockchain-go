package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Transaction una transazione è composta da due aggregati:
// il riferimento a transazioni precedenti
// le transazioni in ingresso (TxInput) e le transazioni in uscita (TxOutput)
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// TxInput in questa implementazione una transazione di input si compone di
// ID in riferimento a transazioni precedenti
// Out il riferimento alla transazione di uscita
// una stringa arbitraria
type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

// TxOutput si compone di un valore e della chiave pubblica del destinatario
type TxOutput struct {
	Value  int
	PubKey string
}

// SetID assegna l'ID alla transazione
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	Handle(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CoinbaseTx la transazione Coinbase è la prima transazione della catena.
// Ne viene aggiunta una anche ad ogni blocco e rappresenta l'incentivo destinato a chi ha formato il blocco
// è una transazione speciale perché non necessità di referenziare nessuna transazione precedente.
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}

	tx := Transaction{
		nil,               // nessuna transazione precedente
		[]TxInput{txin},   // transazioni in ingresso (txin)
		[]TxOutput{txout}} // transazioni in uscita (txout)
	tx.SetID()

	return &tx
}

// NewTransaction crea una nuova transazione
// from = origine
// to = destinatario
// amount = quantità
// riferimento alla catena
func NewTransaction(from, to string, amount int, chain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error: not enough funds")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		Handle(err)

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	outputs = append(outputs, TxOutput{amount, to})

	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from})
	}

	tx := Transaction{nil, inputs, outputs}
	tx.SetID()

	return &tx
}

// IsCoinbase determina se una transazione è la transazione Coinbase
func (tx *Transaction) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

// CanUnlock ?
func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

// CanBeUnlocked ?
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PubKey == data
}
