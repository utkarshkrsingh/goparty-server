// Package roomcode generates the room code for room
package roomcode

import (
	"crypto/rand"
	"math/big"
)

// GenerateCode generates a random alphanumeric code of a specified length (6 characters)
func GenerateCode() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, 7)
	for i := range 7 {
		if i == 3 {
			code[i] = '-'
			continue
		}
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		code[i] = charset[idx.Int64()]
	}
	return string(code), nil
}
