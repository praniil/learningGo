package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Data        string
	PrevHash    string
	CurrentHash string
}

func calculateHash(data string, prevHash string) string {
	hashInput := data + prevHash
	hash := sha256.Sum256([]byte(hashInput))
	return hex.EncodeToString(hash[:])
}

func NewBlock(prevBlock Block, data string) Block {
	var newBlock Block
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.CurrentHash
	newBlock.CurrentHash = calculateHash(newBlock.Data, prevBlock.CurrentHash)
	return newBlock
}

func HashPointer() {
	genesisBlock := Block{
		Data:     "Genesis Block",
		PrevHash: "",
		CurrentHash:     "",
	}
	genesisBlock.CurrentHash = calculateHash(genesisBlock.Data, genesisBlock.PrevHash)

	block1 := NewBlock(genesisBlock, "Transaction Data 1")
	block2 := NewBlock(block1, "Transaction Data 2")

	fmt.Println("Genesis Block")
	fmt.Println("Data:", genesisBlock.Data)
	fmt.Println("Hash:", genesisBlock.CurrentHash)
	fmt.Println()

	fmt.Println("Block 1")
	fmt.Println("Data:", block1.Data)
	fmt.Println("Previous Hash:", block1.PrevHash)
	fmt.Println("Hash:", block1.CurrentHash)
	fmt.Println()

	fmt.Println("Block 2")
	fmt.Println("Data:", block2.Data)
	fmt.Println("Previous Hash:", block2.PrevHash)
	fmt.Println("Hash:", block2.CurrentHash)
}
