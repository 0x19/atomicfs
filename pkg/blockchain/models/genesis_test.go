package models_test

import (
	"testing"

	"github.com/0x19/atomicfs/pkg/blockchain/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestLoadGenesis(t *testing.T) {
	tAssert := assert.New(t)

	genesisJsonPath := GetFromTestDataPath(tAssert, "genesis.json")

	genesis, err := models.LoadGenesisFromPath(genesisJsonPath)
	tAssert.NoError(err)

	tAssert.Equal(genesis.ChainID, uint32(1))
	tAssert.Equal(genesis.ChainName, "Atomic")
	tAssert.Equal(genesis.ChainSymbol, "ATOM")
	tAssert.Equal(genesis.Timestamp.String(), "2022-05-15 00:00:00 +0000 UTC")
	tAssert.Equal(genesis.Coinbase.Hex(), "0x0000000000000000000000000000000000000000")

	tAssert.Equal(
		genesis.Allocations[common.HexToAddress("0x0000000000000000000000000000000000000001")].Balance,
		uint64(10000),
	)

	tAssert.Equal(
		genesis.Allocations[common.HexToAddress("0x0000000000000000000000000000000000000002")].Balance,
		uint64(1),
	)
}
