package account

import (
	"bytes"
	"errors"
)

func Parse(cookie []byte) (*Account, error) {
	ret := Account{}
	res := bytes.Split(cookie, []byte("&"))
	for _, chnk := range res {
		tup := bytes.Split(chnk, []byte("="))
		if len(tup) != 2 {
			return nil, errors.New("invalid input")
		}
		switch {
		case bytes.Equal(tup[0], []byte("email")):
			ret.email = tup[1]
		case bytes.Equal(tup[0], []byte("uid")):
			ret.uid = tup[1]
		case bytes.Equal(tup[0], []byte("role")):
			ret.role = tup[1]
		default:
			return nil, errors.New("invalid input")
		}
	}
	return &ret, nil
}
