package account

import "github.com/juggler434/crypto/padding"

func Attack(userService *UserService) []byte {
	bs := findBlockSize(userService)
	em := createAttackerEmail(bs)
	eb, _ := userService.GetUser(em)

	pad := createAdminPaddedEmail(bs)
	encryptedAdminCookie, _ := userService.GetUser(pad)
	encryptedAdmin := encryptedAdminCookie[bs : bs+bs]

	replaceRole(eb, encryptedAdmin, bs)
	return eb
}

func findBlockSize(userService *UserService) int {
	ie := []byte("A@evil.com")
	io, _ := userService.GetUser(ie)

	for {
		ie = append([]byte("A"), ie...)
		r, _ := userService.GetUser(ie)
		if len(r) > len(io) {
			return len(r) - len(io)
		}
	}
}

func createAttackerEmail(blockSize int) []byte {
	reqInputLen := len([]byte("email=&uid=10&role="))

	r := reqInputLen % blockSize
	el := blockSize - r
	email := make([]byte, el)
	for i := 0; i < len(email); i++ {
		email[i] = byte('A')
	}
	return email
}

func createAdminPaddedEmail(blockSize int) []byte {
	paddedAdmin := pkcs7.Pad([]byte("admin"), blockSize)
	emailHeaderLength := len([]byte("email="))
	pad := make([]byte, abs(blockSize-emailHeaderLength))
	for i := 0; i < len(pad); i++ {
		pad[i] = byte('A')
	}

	pad = append(pad, paddedAdmin...)
	return pad
}

func replaceRole(cookie, encryptedAdmin []byte, blockSize int) []byte {
	for i, j := len(cookie)-blockSize, 0; i < len(cookie); i, j = i+1, j+1 {
		cookie[i] = encryptedAdmin[j]
	}
	return cookie
}

func abs(input int) int {
	if input < 0 {
		return -input
	}
	return input
}
