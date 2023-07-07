package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// deprecated
	// rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random int.
func RandomInt(min, max int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return min + r.Int63n(max-min+1)
}

// RandomString returns a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < n; i++ {
		c := alphabet[r.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random money value
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomEmail returns a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
