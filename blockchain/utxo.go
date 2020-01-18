package blockchain

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/dgraph-io/badger"
)

var (
	utxoPrefix   = []byte("utxo-")
	prefixLength = len(utxoPrefix)
)

// UTXOSet a struct that manages UTXO
type UTXOSet struct {
	Blockchain *BlockChain
}

// FindUnspentTransactions returns unspent transactions for a given PubKeyHash
//
// Per comprendere questo metodo dobbiamo capire cosa rende uno TxOutput spendibile:
// Un TxOutput è spendibile se non è referenziato dalle transazioni di TXInput di nessun blocco
// Infatti, ricordando che le TxInput hanno due campi ID e Out.
// Questi due campi identificano la transazione di riferimento presa come input ed il relativo UTXO
// Se esiste una TXInput che referenzia la coppia (ID, Out) il TXOut è speso e viene messo nella lista dei TXO spesi
//
// Dunque la funzione opera in questo modo:
// 1 - Il ciclo principale itera tutti i blocchi della catena e preleva le transazioni referenziate dal blocco
// 		1.1 - Un ciclo interno itera le transazioni del blocco
// 			2.1 - data una transazione itera gli TxInput della transazione
//       	saltando la Coinbase che per sua natura non referenzia altra transazioni
//       	poiché le transazioni TxInput contengono la coppia (ID, Out)
//			aggiorna la mappa delle spentTXOs
// 			2.2 - questo ciclo itera i TXO
//       		se non sono contenuti nella mappa degli spesi e se solo posseduti dal soggetto emittente li aggiunge alla lista dei UTXO
// # Perché itera prima gli output?
// Ricordiamo che l'iterazione dei blocchi parte dal blocco con altezza maggiore (l'ultimo)
// ... dimsissed

// FindSpendableOutputs restituisce una mappa di UTXO
func (utxoSet UTXOSet) FindSpendableOutputs(pubKeyHash []byte, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	accumulated := 0
	db := utxoSet.Blockchain.Database

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(utxoPrefix); it.ValidForPrefix(utxoPrefix); it.Next() {
			item := it.Item()
			k := item.Key()
			v, err := item.Value()
			Handle(err)
			k = bytes.TrimPrefix(k, utxoPrefix)
			txID := hex.EncodeToString(k)
			outs := DeserializeOutputs(v)

			for outIdx, out := range outs.Outputs {
				if out.IsLockedWithKey(pubKeyHash) && accumulated < amount {
					accumulated += out.Value
					unspentOuts[txID] = append(unspentOuts[txID], outIdx)
				}
			}
		}
		return nil
	})
	Handle(err)
	return accumulated, unspentOuts
}

// FindUTXO cerca e restituisce un UTXO referenziato attraverso pubKeyHash
func (utxoSet UTXOSet) FindUTXO(pubKeyHash []byte) []TxOutput {
	var UTXOs []TxOutput

	db := utxoSet.Blockchain.Database

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions

		it := txn.NewIterator(opts)
		defer it.Close()

		for it.Seek(utxoPrefix); it.ValidForPrefix(utxoPrefix); it.Next() {
			item := it.Item()
			v, err := item.Value()
			Handle(err)
			outs := DeserializeOutputs(v)

			for _, out := range outs.Outputs {
				if out.IsLockedWithKey(pubKeyHash) {
					UTXOs = append(UTXOs, out)
				}
			}
		}

		return nil
	})
	Handle(err)

	return UTXOs
}

// CountTransactions conta le transazione nell UTXO Set
func (utxoSet UTXOSet) CountTransactions() int {
	db := utxoSet.Blockchain.Database
	counter := 0

	err := db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions

		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Seek(utxoPrefix); it.ValidForPrefix(utxoPrefix); it.Next() {
			counter++
		}

		return nil
	})

	Handle(err)

	return counter
}

// Reindex indicizza l'UTXO Set
func (utxoSet UTXOSet) Reindex() {
	db := utxoSet.Blockchain.Database

	utxoSet.DeleteByPrefix(utxoPrefix)

	UTXO := utxoSet.Blockchain.FindUTXO()

	err := db.Update(func(txn *badger.Txn) error {
		for txID, outs := range UTXO {
			key, err := hex.DecodeString(txID)
			if err != nil {
				return err
			}
			key = append(utxoPrefix, key...)

			err = txn.Set(key, outs.Serialize())
			Handle(err)
		}

		return nil
	})
	Handle(err)
}

// Update aggiorna l'UTXO Set
func (utxoSet *UTXOSet) Update(block *Block) {
	db := utxoSet.Blockchain.Database

	err := db.Update(func(txn *badger.Txn) error {
		for _, tx := range block.Transactions {
			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					updatedOuts := TxOutputs{}
					inID := append(utxoPrefix, in.PrevTxID...)
					item, err := txn.Get(inID)
					Handle(err)
					v, err := item.Value()
					Handle(err)

					outs := DeserializeOutputs(v)

					for outIdx, out := range outs.Outputs {
						if outIdx != in.OutIndex {
							updatedOuts.Outputs = append(updatedOuts.Outputs, out)
						}
					}

					if len(updatedOuts.Outputs) == 0 {
						if err := txn.Delete(inID); err != nil {
							log.Panic(err)
						}

					} else {
						if err := txn.Set(inID, updatedOuts.Serialize()); err != nil {
							log.Panic(err)
						}
					}
				}
			}

			newOutputs := TxOutputs{}
			for _, out := range tx.Outputs {
				newOutputs.Outputs = append(newOutputs.Outputs, out)
			}

			txID := append(utxoPrefix, tx.ID...)
			if err := txn.Set(txID, newOutputs.Serialize()); err != nil {
				log.Panic(err)
			}
		}

		return nil
	})
	Handle(err)
}

// DeleteByPrefix cancella
func (utxoSet *UTXOSet) DeleteByPrefix(prefix []byte) {
	deleteKeys := func(keysForDelete [][]byte) error {
		if err := utxoSet.Blockchain.Database.Update(func(txn *badger.Txn) error {
			for _, key := range keysForDelete {
				if err := txn.Delete(key); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}
		return nil
	}

	collectSize := 100000
	utxoSet.Blockchain.Database.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()

		keysForDelete := make([][]byte, 0, collectSize)
		keysCollected := 0
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			key := it.Item().KeyCopy(nil)
			keysForDelete = append(keysForDelete, key)
			keysCollected++
			if keysCollected == collectSize {
				if err := deleteKeys(keysForDelete); err != nil {
					log.Panic(err)
				}
				keysForDelete = make([][]byte, 0, collectSize)
				keysCollected = 0
			}
		}
		if keysCollected > 0 {
			if err := deleteKeys(keysForDelete); err != nil {
				log.Panic(err)
			}
		}
		return nil
	})
}
