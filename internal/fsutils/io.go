package fsutils

import (
	"os"
)

func ReadFile(path string) ([]byte, error) {
	abs, err := NormalizePath(path)
	if err != nil {
		return nil, err
	}

	if err := CheckFile(abs); err != nil {
		return nil, err
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		return nil, ErrReadFailure
	}

	return data, nil
}
