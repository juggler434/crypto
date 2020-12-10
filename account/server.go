package account

import (
	"github.com/juggler434/crypto/aes128/ecb"
	"github.com/juggler434/crypto/aes128/oracle"
)

type UserService struct {
	key []byte
}

func NewUserService() *UserService {
	k, _ := oracle.GenerateKey()
	return &UserService{key: k}
}

func (us *UserService) GetUser(email []byte) ([]byte, error) {
	c, err := Get(email)
	if err != nil {
		return nil, err
	}
	ret, err := ecb.Encrypt(c, us.key)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
