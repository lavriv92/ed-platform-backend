package shared

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncodePassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	encodedPassword := hex.EncodeToString(h.Sum(nil))
	return encodedPassword
}