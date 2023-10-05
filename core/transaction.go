package core

import (
	"crypto/rsa"

	"github.com/doogunwo/Block_chain_go/config"
	"github.com/doogunwo/Block_chain_go/util"
)

type TransactionInput struct {
	PrevtxMap   [config.HashSize]byte
	OutputIndex uint64
	signature   []byte
}

type TransactionOutput struct {
	Value   uint64
	Address rsa.PublicKey
}

type Transaction struct {
	ID      string
	Inputs  []TransactionInput
	Outputs []TransactionOutput
	Sender  rsa.PublicKey
}

func CreateTransaction(ninput int, noutput int) Transaction {

	var tran Transaction

	for i := 0; i < ninput; i++ {
		var input TransactionInput
		tran.Inputs = append(tran.Inputs, input)
	}

	for i := 0; i < noutput; i++ {
		var output TransactionOutput
		tran.Outputs = append(tran.Outputs, output)
	}

	tran.ID = util.GetShortedUniqueId()
	return tran

}
