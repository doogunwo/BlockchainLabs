package core

import (
	"crypto/rsa"
	"encoding/binary"

	"github.com/doogunwo/Block_chain_go/config"
	"github.com/doogunwo/Block_chain_go/util"
)

type unit256 struct {
	data [4]uint64
}

type Block struct {
	hash          [config.HashSize]byte
	prevBlockHash [config.HashSize]byte
	blockIdx      uint64

	BlockValue   uint64
	timeStamps   uint64
	minerAddress rsa.PublicKey
	nuance       uint64

	Transactions []Transaction
}

func createBlock(prevBlockHash [config.HashSize]byte, blockIdx uint64, timeStampMs uint64, minerAddress *rsa.PublicKey, transactions []Transaction) *Block {
	var block Block
	block.prevBlockHash = prevBlockHash
	block.blockIdx = blockIdx
	block.timeStamps = timeStampMs
	block.minerAddress = *minerAddress

	/* Create a special transaction to reward miner (always as transaction 0) */
	block.Transactions = append(block.Transactions, CreateTransaction(1, 1))
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, blockIdx)
	copy(block.Transactions[0].Inputs[0].PrevtxMap[:], b)

	/* TODO: 100 coins, should be adjusted based on timeStamp */
	block.Transactions[0].Outputs[0].Value = config.MinerRewardBase
	block.Transactions[0].Outputs[0].Address = *minerAddress

	/* Add real transactions */
	block.AddTransactions(transactions)

	util.GetBlockLogger().Infof("Added reward transaction: %s\n", block.Transactions[0].Print())
	return &block
}
