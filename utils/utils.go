package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashToken(token string) string {
	hash := sha256.Sum256([]byte(token))
	hashToken := hex.EncodeToString(hash[:])

	return hashToken
}
