package blockchain

import "time"

// CreateNewBlock creates a new block
// in questa parte cambia la firma della funzione
// in quanto vengono inserite le transazioni e non dati arbitrari
func CreateNewBlock(txs []*Transaction, prevHash []byte, height int) *Block {
	block := &Block{
		BlockHeader: BlockHeader{
			Time:          time.Now().Unix(),
			HashPrevBlock: prevHash,
			Bits:          GetNextWorkRequired(),
			Nonce:         0,
		},
		Hash:         []byte{},
		Transactions: txs,
		Height:       height,
	}

	//
	//
	proof := NewProof(block)
	nonce, hash := proof.Run()

	block.Hash = hash[:] // Hash viene settato qui
	block.Nonce = nonce  // Nounce viene settato qui

	return block
}

// Genesis the first block of the chain
// Anche il blocco di genesis cambia introducendo la transazione coinbase
func Genesis(coinbase *Transaction) *Block {
	return CreateNewBlock([]*Transaction{coinbase}, []byte{}, 0)
}
