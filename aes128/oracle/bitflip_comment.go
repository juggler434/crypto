package oracle

import (
	"bytes"
	"github.com/juggler434/crypto/aes128/cbc"
)

const (
	commentPrefix  = "comment1=cooking%20MCs;userdata="
	commentPostfix = ";comment2=%20like%20a%20pound%20of%20bacon"
	isAdmin        = "admin=true"
)

type CommentEncrypter struct {
	key                  []byte
	initializationVector []byte
}

func NewCommentEncrypter() *CommentEncrypter {
	key, _ := GenerateKey()
	iv, _ := GenerateKey()
	return &CommentEncrypter{
		key:                  key,
		initializationVector: iv,
	}
}

func (ce *CommentEncrypter) Encrypt(input []byte) ([]byte, error) {
	si := sanitizeInput(input)

	i := append([]byte(commentPrefix), si...)
	i = append(i, []byte(commentPostfix)...)

	return cbc.Encrypt(i, ce.key, ce.initializationVector)
}

func (ce *CommentEncrypter) CheckIsAdmin(input []byte) (bool, error) {
	di, err := cbc.Decrypt(input, ce.key, ce.initializationVector)
	if err != nil {
		return false, err
	}

	tups := bytes.Split(di, []byte(";"))
	for _, tup := range tups {
		if bytes.Equal(tup, []byte(isAdmin)) {
			return true, nil
		}
	}

	return false, nil
}

func sanitizeInput(input []byte) []byte {
	si := bytes.Replace(input, []byte("="), []byte("\"=\""), -1)
	si = bytes.Replace(si, []byte(";"), []byte("\";\""), -1)
	return si
}
