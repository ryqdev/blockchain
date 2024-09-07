package main

import (
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Data         string
	Hash         string
	PreviousHash string
}

type Blockchain struct {
	Chain    []*Block
	LastHash string
}

func (b *Block) Print() {
	fmt.Printf("Previous Hash: %x\n", b.PreviousHash)
	fmt.Printf("Data: %s\n", b.Data)
	fmt.Printf("Hash: %x\n", b.Hash)
	fmt.Println("---------------------------------------")
}

func CreateBlock(data, previousHash string) Block {
	hash := sha256.Sum256([]byte(data + previousHash))
	return Block{
		data,
		string(hash[:]),
		string(previousHash),
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
	GenesisBlock := CreateBlock("Genesis", "")
	chain.AddToChain(&GenesisBlock)

	block1 := CreateBlock("Alice sent 1 BTC to Bob", chain.LastHash)
	chain.AddToChain(&block1)

	chain.Print()
}
