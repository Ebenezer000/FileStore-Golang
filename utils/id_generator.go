package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerateID generates a random 16-byte hexadecimal ID string
func GenerateID() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	return hex.EncodeToString(bytes)
}
