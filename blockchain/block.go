package blockchain

// BlockChain the blockchain ledger
type BlockChain struct {
	Blocks []*Block
}

// Block a single block in the chain
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// CreateBlock creates a new block
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock add a block to the BlockChain structure
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1] // last block
	new := CreateBlock(data, prevBlock.Hash)       // create a new block
	chain.Blocks = append(chain.Blocks, new)       // append the new block to the chai
}

// Genesis the first block of the chain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain init the BlockChain with the Genesis block
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
