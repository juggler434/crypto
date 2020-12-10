package account

import "bytes"

func Parse(cookie []byte) (*Account, error) {
	ret := Account{}

	res := bytes.Split(cookie, []byte("&"))
	for _, chnk := range res {

	}

	return &Account{}, nil
}
