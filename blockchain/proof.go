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
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Difficulty incrementando la difficoltà aumentano il numero di bytes a 0
// e sarà più difficile trovare un hash inferiore al numero dato
const Difficulty = 18

// ProofOfWork struttura che contiene il blocco e il target
type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof ritorna una struttura pow completa di target
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	pow := &ProofOfWork{b, target}

	return pow
}

// InitData metodo di ProofOfWork che ritorna i dati da sottoporre ad hash
func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),      // Big endian
			ToHex(int64(Difficulty)), // Big endian
		},
		[]byte{},
	)

	return data
}

// Run metodo di ProofOfWork che incrementa il nonce e calcola l'hash fino a trovare il target
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}

	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate valida il nonce
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

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
