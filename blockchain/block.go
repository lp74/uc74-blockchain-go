package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

/*
Questa implementazione della struttura Block è molto semplice
Proseguento lo sviluppo sarà necessario introdurre altri elementi:

- catena principale
- catena secondaria
- orfani

Occorrera anche prendere in considerazione il calcolo della catena con il maggior lavoro svolto,
in maniera tale da decidere quale catena è la principale

*/

// BlockHeader testata del blocco
//
// TESTATA (BLOCK HEADER)
// Version:        4 bytes        Versione del blocco
// HashPrevBlock: 32 bytes        Hash del blocco che viene referenziato da questo blocco (precedente)
// HashMerkleRoot 32 bytes        Hash di tutte le transazioni del blocco ottenuto tramite l'albero di Merkle
// Time:           4 bytes        A timestamp recording when this block was created (Will overflow in 2106[2])
// Bits:           4 bytes        Valore della difficoltà (TARGET) calcolata per questo blocco
// Nonce:          4 bytes        Il Nonce usato per generare il blocco: è utilizzato dall'algoritmo di Proof of Works per generare l'hash del blocco in conformità con il target
//
// TODO: implementare:
// - Version
// - HashMerkleRoot
// - Bits
type BlockHeader struct {
	// Testata (header)
	Version        int
	Time           int64
	HashPrevBlock  []byte
	HashMerkleRoot []byte
	Bits           uint
	Nonce          int
}

// Block Blocco della catena
type Block struct {
	// WIP Testata
	BlockHeader BlockHeader
	// Testata (header)
	Version        int
	Time           int64
	HashPrevBlock  []byte // 32 bytes
	HashMerkleRoot []byte // 32 bytes
	Bits           uint
	Nonce          int

	//
	Hash         []byte         // 32 bytes
	Transactions []*Transaction // Transazioni
	Height       int
}

func (block *Block) setNull() {
	//b.Version = 0
	block.HashPrevBlock = nil
	//b.hashMerkleRoot = nil
	block.Time = 0
	//b.Bits = 0
	block.Nonce = 0
}

// HashTransactions crea l'hash delle transazioni
func (block *Block) HashTransactions() []byte {
	var txHashes [][]byte

	for _, tx := range block.Transactions {
		txHashes = append(txHashes, tx.Hash())
	}
	tree := NewMerkleTree(txHashes)

	return tree.RootNode.Data
}

// Serialize serialize a block
func (block *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(block)

	Handle(err)

	return res.Bytes()
}

// Deserialize deserialize a block
func Deserialize(data []byte) *Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)

	Handle(err)

	return &block
}

// Handle error hanlder
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
