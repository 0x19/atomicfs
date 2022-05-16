package models_test

import (
	"os"
	"path/filepath"

	"github.com/stretchr/testify/assert"
)

func GetFromTestDataPath(tAssert *assert.Assertions, path string) string {
	pwd, err := os.Getwd()
	tAssert.NoError(err)

	return filepath.Join(pwd, "../../..", "test", "data", path)
}
