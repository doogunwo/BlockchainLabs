package main


import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Transaction struct {
	Sender string
	Recipient string
	Amount int
}


type Block struct {
	Index int
	Timestamp int64
	Transaction []Transaction
	PrevHash string
	Hash string
}

type Blockchain struct {
	Blocks []Block
}

func createBlock(prevBlock Block, tx []Transaction) Block {
	var newBlock Block

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Transactions = transactions
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}

func calculateHash(block Block) string {
	record := string(block.Index) + string(block.Timestamp) + block.PrevHash + stringifyTransactions(block.Transactions)
	hash := sha256.new()
	hash.Write([]byte(record))

	hashInBytes := hash.Sum(nil)
	return hex.EncodeToString(hashInBytes)
}

func stringifyTransactions(transactions []Transaction) string {
	var result string
	for _, tx := range transactions {
		result = result + (tx.Sender + tx.Recipient + string(tx.Amount)
	}

	retrun result
}
