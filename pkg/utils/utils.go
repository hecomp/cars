package utils

import (
	"math/rand"
	"strings"
)

const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"

func GenId(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		idx := rand.Int63() % int64(len(letterBytes))
		sb.WriteByte(letterBytes[idx])
	}
	return sb.String()
}
