package aes128

import "github.com/juggler434/crypto/aes128/oracle"

const (
	inputText = "cool&admin&true"
	index1    = 20
	index2    = 26
)

func BitflipAttack(encrypter *oracle.CommentEncrypter) (bool, error) {
	et, err := encrypter.Encrypt([]byte(inputText))
	if err != nil {
		return false, err
	}

	et[index1] = et[index1] ^ '&' ^ ';'
	et[index2] = et[index2] ^ '&' ^ '='

	ia, err := encrypter.CheckIsAdmin(et)
	if err != nil {
		return false, err
	}

	return ia, nil

}
