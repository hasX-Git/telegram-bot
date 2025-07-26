package internal

import (
	"crypto/sha256"
	"encoding/hex"
)

func hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
