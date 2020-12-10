package account

import (
	"errors"
	"fmt"
)

func Get(email []byte) ([]byte, error) {
	err := checkForInvalidCharacters(email)
	if err != nil {
		return nil, err
	}

	ret := fmt.Sprintf("email=%s&uid=10&role=user", email) //Overly simple, but works for now
	return []byte(ret), nil
}

func checkForInvalidCharacters(email []byte) error {
	for _, b := range email {
		if b == '&' {
			return errors.New("invalid character &")
		}
		if b == '=' {
			return errors.New("invalid character =")
		}
	}
	return nil
}
