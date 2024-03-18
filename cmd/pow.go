package pkg

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
	"math"
	"strconv"
    "time"
	"encoding/binary"
)

const targetBits = 24
const maxNonce = math.MaxInt64


type Block struct {
	Timestamp	int64
	Data		[]byte
	prevBlockHash []byte
	Hash		[]byte
	Nonce		int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b, target}
	return pow
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



type ProofOfWork struct {
	block *Block
	target *big.Int
}

func NewPow (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	pow := &ProofOfWork{b,target}
	return pow
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, num)

	return buff.Bytes()
}

func (pow *ProofOfWork) prepareData (nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.prevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)

	for nonce < maxNonce {
			data := pow.prepareData(nonce)
			hash = sha256.Sum256(data)
			fmt.Println("\r%x", hash)

			hashInt.SetBytes(hash[:])
			if hashInt.Cmp(pow.target) == -1 {
					break
			} else {
					nonce++
			}
	}
	fmt.Println("mining complete")
	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int
	var isValid bool
	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	isValid = hashInt.Cmp(pow.target) == -1
	return isValid
}


func main() {
	bc := NewBlockchain()

	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
			fmt.Printf("Prev. hash: %x\n", block.prevBlockHash)
			fmt.Printf("Data: %s\n", block.Data)
			fmt.Printf("Hash: %x\n", block.Hash)
			pow := NewProofOfWork(block)
    		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
			fmt.Println()
	}
}

