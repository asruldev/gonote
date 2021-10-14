package helpers

import (
	"golang.org/x/crypto/scrypt"
)

func Passwordhash(pw, salt []byte) ([]byte, error) {
	return scrypt.Key(pw, salt, 32768, 8, 1, 32)
}
