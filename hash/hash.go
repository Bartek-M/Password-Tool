package hash

import (
	"crypto/sha512"
	"encoding/hex"
)

func Hash(passw string) string {
	hasher := sha512.New()
	hasher.Write([]byte(passw))

	hashSum := hasher.Sum(nil)
	hexHash := hex.EncodeToString(hashSum)

	return hexHash
}
