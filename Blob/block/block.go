package block

import (
    "bytes"
    "crypto/sha256"
    "strconv"
    "time"
    "fmt"
)
type Block struct {
	Timestamp int64
	Data []byte
	PrevBlockHash []byte
	Hash []byte
}

type Blockchain struct {
	blocks []*Block
}

 func (b *Block) SetHash() {
	Timestamp 	:=	[]byte (strconv.FormatInt(b.Timestamp,10))
	headers 	:=	bytes.Join([][]byte{b.PrevBlockHash,b.Data,Timestamp},[]byte{})
	hash 		:=	sha256.Sum256(headers)
	b.Hash = hash[:]
 }

 func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string){
	prevBlock	:= bc.blocks[len(bc.blocks)-1]
	newBlock	:= NewBlock(data,prevBlock.Hash)	 	
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block",[]byte{})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{ []*Block{NewGenesisBlock()}}
}

func Block_test() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			fmt.Println()
	}
}

