package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
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
type BlockHeader struct {
	// Testata (header)
	Version        int //TODO: da implementare
	Time           int64
	HashPrevBlock  []byte
	HashMerkleRoot []byte //TODO: da implementare
	Bits           uint   //TODO: da implementare
	Nonce          int
}

// Block Blocco della catena
//
// TODO: implementare:
// - Version
// - HashMerkleRoot
// - Bits
type Block struct {
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

// CreateBlock creates a new block
// in questa parte cambia la firma della funzione
// in quanto vengono inserite le transazioni e non dati arbitrari
func CreateBlock(txs []*Transaction, prevHash []byte, height int) *Block {
	block := &Block{
		Time:          time.Now().Unix(),
		Hash:          []byte{},
		HashPrevBlock: prevHash,
		Nonce:         0,

		Height:       height,
		Transactions: txs,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// Genesis the first block of the chain
// Anche il blocco di genesis cambia introducendo la transazione coinbase
func Genesis(coinbase *Transaction) *Block {
	return CreateBlock([]*Transaction{coinbase}, []byte{}, 0)
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
