package util

import (
	"math/rand"
	"time"
)

const numberCharset = "0123456789"

func GenerateCode(length int) string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	code := make([]byte, length)
	for i := range code {
		code[i] = numberCharset[rnd.Intn(len(numberCharset))]
	}
	return string(code)
}
