package oracle

type Encrypter interface {
	Encrypt([]byte, []byte) ([]byte, error)
}
