package account

import "fmt"

//not elegent, but it works for what we need
func CreateCookie(account *Account) []byte {
	ret := fmt.Sprintf("email=%s&uid=%s&role=%s", account.email, account.uid, account.role)
	return []byte(ret)
}
