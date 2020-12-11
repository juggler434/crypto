package account

import (
	"bytes"
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

func (us *UserService) CheckAdminPermission(encryptedCookie []byte) (bool, error) {
	uc, err := ecb.Decrypt(encryptedCookie, us.key)
	if err != nil {
		return false, err
	}

	a, err := Parse(uc)
	if err != nil {
		return false, err
	}
	isAdmin := bytes.Equal(a.role, []byte("admin"))
	return isAdmin, nil
}
