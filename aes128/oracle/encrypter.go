package oracle

type Encrypter interface {
	Encrypt([]byte) ([]byte, error)
}
