package util

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Password(password string) string {

	h := md5.New()
	h.Write([]byte(password))

	passwordHash := h.Sum(nil)
	enString := hex.EncodeToString(passwordHash)

	return enString
}
