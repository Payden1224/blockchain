package blockchain

import (
	"bytes"
	"crypto/sha256"
)

type BlockChain struct {
	Blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte 
}

func CreateBlock(data string, prevHash []byte) *Block {
	Block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	
	return block
}

func (chain *Blockchain) addBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.hash)
	chain.Blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []bytes{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", Block.PrevHash)
		fmt.Printf("Data in Block: %s\n", Block.Data)
		fmt.Printf("Hash: %x\n", Block.Hash)
	}

}