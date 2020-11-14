package hex

import "encoding/hex"

func Encode(plainText []byte) []byte {
	ret := make([]byte, hex.EncodedLen(len(plainText)))
	hex.Encode(ret, plainText)
	return ret
}

func Decode(hexBytes []byte) ([]byte, error) {
	ret := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(ret, hexBytes)
	if err != nil {
		return nil, err
	}
	return ret, nil
}




