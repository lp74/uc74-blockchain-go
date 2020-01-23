package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/dgraph-io/badger" // BadgerDB (Bitcoin-core usa LevelDB)
)

/*
Questa implementazione della struttura Block è molto semplice
Proseguento lo sviluppo sarà necessario introdurre altri elementi:

- Catena principale `MainChain`
- Catena secondaria `SecondChain`
- Blocchi orfani `Orphans`

Per la gestione della catena secondaria sarà necessario implementare dei metodi (legati a difficulty)
per ottenere il lavoro svolto dalla catena e decidere quale è la principale

*/

const (
	dbPath      = "./tmp/blocks_%s"
	genesisData = "First Transaction from Genesis"
)

// Chain the blockchain ledger
type Chain struct {
	LastHash []byte
	Database *badger.DB
}

// Iterator the iterator struct
// it keeps a reference to the DB and the last hash
type Iterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

// DBexists restituisce true se il database esiste
func DBexists(path string) bool {
	if _, err := os.Stat(path + "/MANIFEST"); os.IsNotExist(err) {
		return false
	}

	return true
}

// ContinueBlockChain restituisce la Block Chain
func ContinueBlockChain(nodeID string) *Chain {
	path := fmt.Sprintf(dbPath, nodeID)
	if DBexists(path) == false {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path

	db, err := openDB(path, opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		return err
	})
	Handle(err)

	chain := Chain{lastHash, db}

	return &chain
}

// InitBlockChain init the BlockChain with the Genesis block
func InitBlockChain(address, nodeID string) *Chain {
	path := fmt.Sprintf(dbPath, nodeID)
	if DBexists(path) {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}
	var lastHash []byte
	opts := badger.DefaultOptions
	opts.Dir = path
	opts.ValueDir = path

	db, err := openDB(path, opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)
		fmt.Println("Genesis created")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		lastHash = genesis.Hash

		return err

	})

	Handle(err)

	blockchain := Chain{lastHash, db}
	return &blockchain
}

// AddBlock add a block to the BlockChain structure
func (chain *Chain) AddBlock(block *Block) {
	err := chain.Database.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get(block.Hash); err == nil {
			return nil
		}

		blockData := block.Serialize()
		err := txn.Set(block.Hash, blockData)
		Handle(err)

		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, _ := item.Value()

		item, err = txn.Get(lastHash)
		Handle(err)
		lastBlockData, _ := item.Value()

		lastBlock := Deserialize(lastBlockData)

		if block.Height > lastBlock.Height {
			err = txn.Set([]byte("lh"), block.Hash)
			Handle(err)
			chain.LastHash = block.Hash
		}

		return nil
	})
	Handle(err)
}

// GetBestHeight restituisce l'altezza
func (chain *Chain) GetBestHeight() int {
	var lastBlock Block

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, _ := item.Value()

		item, err = txn.Get(lastHash)
		Handle(err)
		lastBlockData, _ := item.Value()

		lastBlock = *Deserialize(lastBlockData)

		return nil
	})
	Handle(err)

	return lastBlock.Height
}

// GetBlock dato l'hash restituisce un blocco o un errore
func (chain *Chain) GetBlock(blockHash []byte) (Block, error) {
	var block Block

	err := chain.Database.View(func(txn *badger.Txn) error {
		if item, err := txn.Get(blockHash); err != nil {
			return errors.New("Block is not found")
		} else {
			blockData, _ := item.Value()

			block = *Deserialize(blockData)
		}
		return nil
	})
	if err != nil {
		return block, err
	}

	return block, nil
}

// GetBlockHashes restituisce gli hash
func (chain *Chain) GetBlockHashes() [][]byte {
	var blocks [][]byte

	iter := chain.Iterator()

	for {
		block := iter.Next()

		blocks = append(blocks, block.Hash)

		if len(block.HashPrevBlock) == 0 {
			break
		}
	}

	return blocks
}

// MineBlock crea nuovi blocchi
func (chain *Chain) MineBlock(transactions []*Transaction) *Block {
	var lastHash []byte
	var lastHeight int

	for _, tx := range transactions {
		if chain.VerifyTransaction(tx) != true {
			log.Panic("Invalid Transaction")
		}
	}

	err := chain.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()

		item, err = txn.Get(lastHash)
		Handle(err)
		lastBlockData, _ := item.Value()

		lastBlock := Deserialize(lastBlockData)

		lastHeight = lastBlock.Height

		return err
	})
	Handle(err)

	newBlock := CreateNewBlock(transactions, lastHash, lastHeight+1)

	err = chain.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)

	return newBlock
}

// FindUTXO restituisce UTXO iterando la catena
func (chain *Chain) FindUTXO() map[string]TxOutputs {
	UTXO := make(map[string]TxOutputs)
	spentTXOs := make(map[string][]int)

	iter := chain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Vout {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				outs := UTXO[txID]
				outs.Outputs = append(outs.Outputs, out)
				UTXO[txID] = outs
			}
			if tx.IsCoinbase() == false {
				for _, in := range tx.Vin {
					inTxID := hex.EncodeToString(in.Prevout.PrevTxID)
					spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Prevout.OutIndex)
				}
			}
		}

		if len(block.HashPrevBlock) == 0 {
			break
		}
	}
	return UTXO
}

// FindTransaction trova una transazione attraverso l'ID (hash)
func (chain *Chain) FindTransaction(ID []byte) (Transaction, error) {
	iter := chain.Iterator()

	for {
		block := iter.Next()

		for _, tx := range block.Transactions {
			if bytes.Compare(tx.ID, ID) == 0 {
				return *tx, nil
			}
		}

		if len(block.HashPrevBlock) == 0 {
			break
		}
	}

	return Transaction{}, errors.New("Transaction does not exist")
}

// SignTransaction firma la transazione
func (chain *Chain) SignTransaction(tx *Transaction, privKey ecdsa.PrivateKey) {
	prevTXs := make(map[string]Transaction)

	for _, in := range tx.Vin {
		prevTX, err := chain.FindTransaction(in.Prevout.PrevTxID)
		Handle(err)
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	tx.Sign(privKey, prevTXs)
}

// VerifyTransaction verifica la transazione tx
// costruisce una mappa ( k, v ) = (TxI, Transaction)
// e la sottopone a verifica
func (chain *Chain) VerifyTransaction(tx *Transaction) bool {
	if tx.IsCoinbase() {
		return true
	}
	prevTXs := make(map[string]Transaction)

	for _, in := range tx.Vin {
		prevTX, err := chain.FindTransaction(in.Prevout.PrevTxID)
		Handle(err)
		prevTXs[hex.EncodeToString(prevTX.ID)] = prevTX
	}

	return tx.Verify(prevTXs)
}

func retry(dir string, originalOpts badger.Options) (*badger.DB, error) {
	lockPath := filepath.Join(dir, "LOCK")
	if err := os.Remove(lockPath); err != nil {
		return nil, fmt.Errorf(`removing "LOCK": %s`, err)
	}
	retryOpts := originalOpts
	retryOpts.Truncate = true
	db, err := badger.Open(retryOpts)
	return db, err
}

func openDB(dir string, opts badger.Options) (*badger.DB, error) {
	if db, err := badger.Open(opts); err != nil {
		if strings.Contains(err.Error(), "LOCK") {
			if db, err := retry(dir, opts); err == nil {
				log.Println("database unlocked, value log truncated")
				return db, nil
			}
			log.Println("could not unlock database:", err)
		}
		return nil, err
	} else {
		return db, nil
	}
}
