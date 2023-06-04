package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

// RandomFloat generates a random float between min and max
func RandomFloat(min, max int) float64 {
	return float64(min + rand.Intn(max-min+1))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(k)]
		sb.WriteByte(ch)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() float64 {
	return RandomFloat(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "IDR"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
