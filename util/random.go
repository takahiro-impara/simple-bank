package util

import (
	"math/rand"
	"strings"
	"time"
)

var randGen *rand.Rand

const alphabet = "abcdefghiklmnopqrstuvwxyz"

func init() {
	seed := time.Now().UnixNano()
	randGen = rand.New(rand.NewSource(seed))
}

// RandomInt returns a random integer between min and max (inclusive)
func RandomInt(min, max int64) int64 {
	if min >= max {
		return min // or handle error appropriately
	}
	return min + randGen.Int63n(max-min+1)
}

// RandomString generated a random string
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn((k))]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generated a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generated a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// RandomCurrency generated a random currency code
func RandomCurrency() string {
	currencies := []string{
		"EUR",
		"USD",
		"CAD",
	}

	n := len(currencies)
	return currencies[rand.Intn(n)]
}
