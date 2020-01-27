package blockchain

/*
PoW

Dato un blocco:
- creare un nonce (1)
- creare l'hash dei dati del blocco addizionano il valore nonce
- controllare se l'hash rispetta le richieste: i primi x bytes devono contenere 0

1. in crittografia il termine nonce indica un numero, in genere causale o pseudo-casuale, che ha un utilizzo unico.
Nonce deriva dall'espressione "for the nince" che significa per l'occasione.
Nel caso specifico il nonce è un valore che viene incrementato fin
*/

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/lp74/uc74-blockchain-go/hash"
)

var (
	bigOne    = big.NewInt(1)
	oneLsh256 = new(big.Int).Lsh(bigOne, 256)
)

// Bits incrementando la difficoltà aumentano il numero di bytes a 0
// e sarà più difficile trovare un hash inferiore al numero dato
// [come calcolare nBits](https://bitcoin.stackexchange.com/questions/2924/how-to-calculate-new-bits-value)
const Bits = 12 // in bitcoin il blocco genesis ha nBits = 0x1d00ffff rappresentazione Compat

// ProofOfWork struttura che contiene il blocco e il target
type ProofOfWork struct {
	Block  *Block
	Target *big.Int // TODO: implementare Compat, chainParams e ricavare la difficoltà dai blocchi
}

// NewProof ritorna una struttura pow completa di target
func NewProof(block *Block) *ProofOfWork {
	target := bigOne
	target.Lsh(target, uint(256-Bits)) // shift a sinistra di 256 - CompatToBig(Bits) = 256 - 12

	pow := &ProofOfWork{block, target}

	return pow
}

// InitData metodo di ProofOfWork che ritorna i dati da sottoporre ad hash
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.HashPrevBlock,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(Bits)),
		},
		[]byte{},
	)

	return data
}

// Run metodo di ProofOfWork che incrementa il nonce e calcola l'hash fino a trovare il target
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var sha []byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		sha = hash.Hash(data)

		fmt.Printf("\r%x", sha)
		intHash.SetBytes(sha[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, sha[:]
}

// GetNextWorkRequired utilizza nBits dell'ultimo blocco
// per computare nBits del prossimo blocco
func GetNextWorkRequired() uint {
	return 12
}

// CalculateNextWorkRequired calcola la difficoltà
func CalculateNextWorkRequired(block *Block) uint {
	return 12
}

// CheckProofOfWork valida il nonce
func (pow *ProofOfWork) CheckProofOfWork() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	sha := hash.Hash(data)
	intHash.SetBytes(sha[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex dato un int 64 lo scrive in formato BigEndian
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
