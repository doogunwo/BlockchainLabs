package main

import (
	"github.com/boltdb/bolt"
	"bytes"
	"encoding/binary"
	"math"
	"math/big"
)

const dbFile = "blockchain_%s.db"
const blocksBucket = "blocks"
const genesisCoinbaseData = "The Times 03/Jan/2009 Chancellor on brink of second bailout for banks"


const targetBits = 24
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

type Blockchain struct {
	tip []byte
	db  *bolt.DB

}

type Block struct {
	Timestamp     int64
	Data          []byte
	prevBlockHash []byte
	Hash          []byte
	Nonce         int
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}