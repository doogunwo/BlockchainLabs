package main

import (
	"github.com/boltdb/bolt"
)


func (bc *Blockchain) AddBlock(data string) {
	
	var lastHash []byte

	err := bc.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		lastHash = b.Get([]byte("1"))
		return nil
	})

	if err != nil {
		panic(err)
	}


	newBlock := NewBlock(data, lastHash)

	err = bc.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			panic(err)
		}
		err = b.Put([]byte("1"), newBlock.Hash)
		if err != nil {
			panic(err)
		}
		bc.tip = newBlock.Hash
		return nil
	})

	if err != nil {
		panic(err)
	}

}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	
	err = db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blocksBucket))
			if b == nil {
					genesis := NewGenesisBlock()
					b, err := tx.CreateBucket([]byte(blocksBucket))

					if err != nil {
						return err
					}
					err = b.Put(genesis.Hash, genesis.Serialize())
					err = b.Put([]byte("l"), genesis.Hash)

				
					
					tip = genesis.Hash
			} else {
					tip = b.Get([]byte("l"))
			}
			return nil
	})

	if err != nil {
		panic(err)
	}

	bc := Blockchain{tip, db}
	return &bc
}

func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{bc.tip,bc.db}
	return bci
}

func (i *BlockchainIterator) Next() *Block {
	var block *Block

	err := i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blocksBucket))
		encodedBlock := b.Get(i.currentHash)
		block = DeserializeBlock(encodedBlock)

		return nil
	})

	i.currentHash = block.PrevBlockHash
	return block
}