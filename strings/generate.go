package strings

import (
    crand "crypto/rand"
	"encoding/hex"
	"math/rand"
)	


var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CRandomString(n int) string {
	p := make([]byte, n/2+1)
	crand.Read(p)
	return hex.EncodeToString(p)[:n]
}