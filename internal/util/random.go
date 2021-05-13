package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())

}

//RandomInt generates a random number between min and max
func RandomUserID(min, max int) int {
	return max + rand.Intn(max-min+1)
}

//RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

//RandomContent generate a random owner name
func RandomCountry() string {
	return RandomString(10)
}
func RandomState() string {
	return RandomString(10)
}
func RandomCity() string {
	return RandomString(10)
}
func RandomStreet() string {
	return RandomString(10)
}
func RandomPostalCode() string {
	return RandomString(8)
}
