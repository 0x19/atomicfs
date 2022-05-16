package fs

import (
	"errors"
	"os"
	"path"
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

func GetDirFromHome(p ...string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	toUse := []string{home}
	return path.Join(append(toUse, p...)...), nil
}
