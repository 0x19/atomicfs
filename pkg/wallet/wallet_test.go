package wallet_test

import (
	"path/filepath"
	"testing"

	"github.com/0x19/atomicfs/pkg/fs"
	"github.com/0x19/atomicfs/pkg/wallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "/tmp/atomicfs"
const testSecret = "I@mTesT..!)$_+"

func TestNewKeyStoreAccount(t *testing.T) {
	tAssert := assert.New(t)

	keystoreDir := wallet.GetKeystoreDirPath(testDataPath)
	tAssert.Equal(keystoreDir, filepath.Join(testDataPath, "keystore"))

	err := fs.CreatePathIfNotExist(testDataPath)
	tAssert.NoError(err)

	address, err := wallet.NewKeystoreAccount(testDataPath, testSecret)
	tAssert.NoError(err)
	tAssert.IsType(address, common.Address{})
	tAssert.NotEmpty(address.String())
}

func TestNewRandomKey(t *testing.T) {
	tAssert := assert.New(t)
	key, err := wallet.NewRandomKey()
	tAssert.NoError(err)
	tAssert.NotEmpty(key.Id.String())
	tAssert.NotEmpty(key.Address.String())
}
