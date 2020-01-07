package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

// BlockChain the blockchain ledger
type BlockChain struct {
	blocks []*Block
}

// Block a single block in the chain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash compute the hash of a block
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock creates a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock add a block to the BlockChain structure
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1] // last block
	new := CreateBlock(data, prevBlock.Hash)       // create a new block
	chain.blocks = append(chain.blocks, new)       // append the new block to the chai
}

// Genesis the first block of the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain init the BlockChain with the Genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}
