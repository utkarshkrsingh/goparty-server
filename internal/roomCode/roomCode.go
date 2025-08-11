// Package roomcode provides functionalities for room code management
package roomcode

import (
	"crypto/rand"
	"math/big"
)

// GenerateCode generates a random alphanumeric code of a specified length (6 characters)
func GenerateCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	code := make([]byte, 6)
	for i := range 6 {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[idx.Int64()]
	}
	return string(code), nil
}
