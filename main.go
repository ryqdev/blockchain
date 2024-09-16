package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

const Difficulty = 16

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return &ProofOfWork{
		b,
		target,
	}
}

func (p *ProofOfWork) Run() (int64, []byte) {
	var nonce int64 = 0
	var hash [32]byte
	for nonce < math.MaxInt32 {
		data := bytes.Join(
			[][]byte{
				p.Block.Data,
				p.Block.PreviousHash,
				int2Bytes(nonce),
			},
			[]byte{},
		)
		hash = sha256.Sum256(data)
		fmt.Printf("\r0x%x", hash)
		result := new(big.Int)
		result.SetBytes(hash[:])
		if result.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()
	return nonce, hash[:]
}

func (p *ProofOfWork) validate() bool {
	data := bytes.Join(
		[][]byte{
			p.Block.Data,
			p.Block.PreviousHash,
			int2Bytes(p.Block.Nonce),
		},
		[]byte{},
	)
	hash := sha256.Sum256(data)
	result := new(big.Int)
	result.SetBytes(hash[:])
	return result.Cmp(p.Target) == -1
}

type Block struct {
	Data         []byte
	Hash         []byte
	PreviousHash []byte
	Nonce        int64
}

type Blockchain struct {
	Chain    []*Block
	LastHash []byte
}

func (b *Block) Print() {
	fmt.Printf("Previous Hash: 0x%x\n", b.PreviousHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: 0x%x\n", b.Hash)
	fmt.Println("---------------------------------------")
}

func CreateBlock(data string, previousHash []byte) Block {
	block := Block{
		Data:         []byte(data),
		PreviousHash: previousHash,
	}
	pow := NewProof(&block)
	nonce, hash := pow.Run()
	block.Nonce = nonce
	block.Hash = hash

	return block
}

func CreateChain() Blockchain {
	return Blockchain{}
}

func (c *Blockchain) AddToChain(block *Block) {
	c.Chain = append(c.Chain, block)
	c.LastHash = block.Hash
}

func (c *Blockchain) Print() {
	for _, b := range c.Chain {
		b.Print()
	}
}

func main() {
	chain := CreateChain()
	GenesisBlock := CreateBlock("Genesis", []byte{})
	chain.AddToChain(&GenesisBlock)

	block1 := CreateBlock("Alice sent 1 BTC to Bob", chain.LastHash)
	chain.AddToChain(&block1)

	chain.Print()
}

func int2Bytes(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
