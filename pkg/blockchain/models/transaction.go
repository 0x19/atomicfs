package models

import (
	"time"

	blockchainpb "github.com/0x19/atomicfs/protos/blockchain/dist/golang"
	"github.com/ethereum/go-ethereum/common"
)

type Transaction struct {
	Sender    common.Address `json:"sender"`
	Recipient common.Address `json:"recipient"`
	Amount    uint64         `json:"amount"`
	Nonce     uint64         `json:"nonce"`
	Gas       uint64         `json:"gas"`
	GasPrice  uint64         `json:"gas_price"`
	Data      string         `json:"data"`
	Timestamp time.Time      `json:"timestamp"`
}

func (t Transaction) IsReward() bool {
	return t.Data == "reward"
}

type SignedTransaction struct {
	Transaction Transaction `json:"transaction"`
	Signature   []byte      `json:"signature"`
}

func (st *SignedTransaction) ToProto() *blockchainpb.Transaction {
	return &blockchainpb.Transaction{}
}
