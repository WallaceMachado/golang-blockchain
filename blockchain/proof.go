package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// algoritimo de consenso ou prova de algoritimo - vamos implementar o tipo prova de trabalho

// Take the data from the block

// create a counter (nonce) which starts at 0

// create a hash of the data plus the counter

// check the hash to see if it meets a set of requirements

// Requirements:
// The First few bytes must contain 0s

const Difficulty = 10

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

//400000000000000000000000000000000000000000000000000000000000
//342dc11a9fd1833ed9fe18ca5627cedc56507de6698acfcafd301398cb35
//40000000000000000000000000000000000000000000000000000000000000
//22cb2a9daceeb53d15dd77a4bc896eeb6c9bacb2453d4e2be143a2be8bdeab
//400000000000000000000000000000000000000000000000000000000000
//4000000000000000000000000000000000000000000000000000000000
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	fmt.Printf("target: %x\n", target)
	target.Lsh(target, uint(256-Difficulty))
	fmt.Printf("target lsh : %x\n", target)
	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

// create a hash of the data plus the counter
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
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:]) // convert o has em big int

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
