package blockchain

import (
	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

type BlockChain struct {
	LastHash []byte
	Database *badger.DB
}


func InitBlockChain() *BlockChain {
	var lastHashk []byte

	opts := badger.DefualtOptions
	opts.Dir = dbPath
	opt.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err := db.Update(func(txn *badger.Txn) error {
		if _, err:= txn.Get([]byte("lh")); err == badger.ErrKeyNotFound
			fmt.Println("No existing blockchain found")
			genesis := Genesis,()
	})
}

func (chain *Blockchain) addBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.hash)
	chain.Blocks = append(chain.blocks, new)
} 
