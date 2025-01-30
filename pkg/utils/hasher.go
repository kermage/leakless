package utils

import (
	"crypto/sha256"
	"hash"
	"io"
	"os"
)

type hasher struct {
	hash.Hash
}

func (h *hasher) hash(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := io.Copy(h, file); err != nil {
		return err
	}

	return nil
}

func HashFiles(paths []string) ([]byte, error) {
	hasher := &hasher{sha256.New()}

	for _, path := range paths {
		if err := hasher.hash(path); err != nil {
			return nil, err
		}
	}

	return hasher.Sum(nil), nil
}
