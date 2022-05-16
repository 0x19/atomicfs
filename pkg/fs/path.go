package fs

import (
	"errors"
	"os"
)

// CreatePathIfNotExist - Creates the path if one is not present on the machine
func CreatePathIfNotExist(path string) error {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}
