package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

const Difficulty = 5

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

type Block struct {
	Data         []byte
	Hash         []byte
	PreviousHash []byte
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
	hash := sha256.Sum256(append([]byte(data), previousHash...))
	return Block{
		[]byte(data),
		hash[:],
		previousHash,
	}
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
