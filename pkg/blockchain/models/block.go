package models

import (
	"encoding/json"
	"fmt"
	"time"

	blockchainpb "github.com/0x19/atomicfs/protos/blockchain/dist/golang"
	"github.com/ethereum/go-ethereum/common"
)

type BlockHeader struct {
	Parent    common.Hash    `json:"parent"`
	Number    uint64         `json:"number"`
	Nonce     uint64         `json:"nonce"`
	Miner     common.Address `json:"address"`
	Timestamp time.Time      `json:"timestamp"`
}

type Block struct {
	Header       BlockHeader         `json:"header"`
	Transactions []SignedTransaction `json:"transactions"`
}

func (b Block) JsonEncode() ([]byte, error) {
	return json.Marshal(b)
}

func (b Block) ToProto() *blockchainpb.Block {
	return &blockchainpb.Block{
		Header: &blockchainpb.BlockHeader{
			Parent:    b.Header.Parent.Hex(),
			Number:    b.Header.Number,
			Nonce:     b.Header.Nonce,
			Miner:     b.Header.Miner.Hex(),
			Timestamp: b.Header.Timestamp.UnixNano(),
		},
		Transactions: func() []*blockchainpb.Transaction {
			transactions := make([]*blockchainpb.Transaction, len(b.Transactions))

			for _, transaction := range b.Transactions {
				transactions = append(transactions, transaction.ToProto())
			}

			return transactions
		}(),
	}
}

func (b Block) Hash() (common.Hash, error) {
	blockJson, err := b.JsonEncode()
	if err != nil {
		return common.Hash{}, err
	}

	return common.BytesToHash(blockJson), nil
}

func (b Block) GasReward() uint64 {
	reward := uint64(0)

	//for _, tx := range b.Transactions {
	//	reward += tx.GasCost()
	//}

	return reward
}

func IsBlockHashValid(hash common.Hash, miningDifficulty uint) bool {
	zeroesCount := uint(0)

	for i := uint(0); i < miningDifficulty; i++ {
		if hash.Hex() == "0x" {
			zeroesCount++
		}
	}

	if fmt.Sprintf("%x", hash[miningDifficulty]) == "0" {
		return false
	}

	return zeroesCount == miningDifficulty
}

func NewBlock(parent common.Hash, number uint64, nonce uint64, time time.Time, miner common.Address, txs []SignedTransaction) Block {
	return Block{BlockHeader{parent, number, nonce, miner, time}, txs}
}
