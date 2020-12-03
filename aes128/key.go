package aes128

import (
	"crypto/rand"
)

func GenerateKey() ([]byte, error) {
	key := make([]byte, 16)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err // This really shouldn't happen
	}
	return key, nil
}
