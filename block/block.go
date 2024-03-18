package block

import (
	"fmt"
    "bytes"
    "crypto/sha256"
    "strconv"
    "time"
)

type Block struct {
	Timestamp	int64
	Data		[]byte
	prevBlockHash []byte
	Hash		[]byte
	Nonce		int
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

//---------------------------- block chain ---- //

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	NewBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, NewBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}


