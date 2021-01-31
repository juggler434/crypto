package pkcs7

import "errors"

var InvalidPaddingError = errors.New("invalid padding")

func Pad(input []byte, blockSize int) []byte {
	r := len(input) % blockSize
	var pl int
	pl = blockSize - r // This makes it so we won't pad input with the correct length
	for i := 0; i < pl; i++ {
		input = append(input, byte(pl))
	}

	return input
}

func Unpad(input []byte) ([]byte, error) {
	if input == nil || len(input) == 0 {
		return nil, nil
	}
	pc := input[len(input)-1]
	pl := int(pc)

	err := checkPaddingIsValid(input, pl)

	if err != nil {
		return nil, err
	}

	return input[:len(input)-pl], nil

}

func checkPaddingIsValid(input []byte, paddingLength int) error {
	if len(input) < paddingLength {
		return InvalidPaddingError
	}

	p := input[len(input)-(paddingLength):]

	for _, pc := range p {
		if uint(pc) != uint(len(p)) {
			return InvalidPaddingError
		}
	}
	return nil
}
