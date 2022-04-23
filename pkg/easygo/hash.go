package easy

import "crypto/sha512"

func HashSecret(password string) [48]byte {
	return sha512.Sum384([]byte(password))
}
