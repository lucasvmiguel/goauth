package token

import (
	"math/rand"
	"time"
)

const (
	tokenSize = 20
	chars     = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890abcdefghijklmnopqrstuvwxyz"
)

//Generate a random token(size = 20 | format = string)
func Generate() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var token [tokenSize]byte

	for i := 0; i < tokenSize; i++ {
		token[i] = chars[rand.Intn(len(chars))]
	}

	return string(token[:])
}
