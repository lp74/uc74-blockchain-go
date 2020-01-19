package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
	"time"
)

// Block a single block in the chain
type Block struct {
	Time          int64
	Hash          []byte
	Transactions  []*Transaction // Transazioni
	HashPrevBlock []byte
	Nonce         int
	Height        int
	// Version      int
	// Bits uint
	// hashMerkleRoot []byte
}

func (b *Block) setNull() {
	//b.Version = 0
	b.HashPrevBlock = nil
	//b.hashMerkleRoot = nil
	b.Time = 0
	//b.Bits = 0
	b.Nonce = 0
}

// HashTransactions crea l'hash delle transazioni
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte

	for _, tx := range b.Transactions {
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
		Transactions:  txs,
		HashPrevBlock: prevHash,
		Nonce:         0,
		Height:        height,
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
func (b *Block) Serialize() []byte {
	var res bytes.Buffer
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)

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
