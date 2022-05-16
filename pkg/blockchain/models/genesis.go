package models

import (
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type GenesisAllocation struct {
	Balance uint64 `json:"balance"`
}

type Genesis struct {
	ChainID     uint32                               `json:"chain_id"`
	ChainName   string                               `json:"chain_name"`
	ChainSymbol string                               `json:"chain_symbol"`
	Timestamp   time.Time                            `json:"timestamp"`
	Coinbase    common.Address                       `json:"coinbase"`
	Allocations map[common.Address]GenesisAllocation `json:"allocs"`
}

func LoadGenesisFromPath(path string) (*Genesis, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var toReturn Genesis

	if err := json.Unmarshal(content, &toReturn); err != nil {
		return nil, err
	}

	return &toReturn, nil
}
