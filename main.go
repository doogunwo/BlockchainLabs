package main

import (
	"fmt"
    "strconv"
	"block"
	"pow"
)

func Test() {
    bc := NewBlockchain()

    bc.AddBlock("Genesis Block")
    bc.AddBlock("Send 1 BTC to Ivan")
    bc.AddBlock("Send 2 more BTC to Ivan")

    for i, block := range bc.blocks {
        fmt.Printf("Block %d\n", i)
        fmt.Printf("Prev. hash: %x\n", block.prevBlockHash)
        fmt.Printf("Data: %s\n", block.Data)
        fmt.Printf("Hash: %x\n\n", block.Hash)
		
		pow := NewProofOfWork(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
    }
}


func main(){
	Test()
}