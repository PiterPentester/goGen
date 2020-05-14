package abracadabra

import (
	"math/rand"
	"time"
)

const Symbols = "~!@#$%^&*()_{}<>"
const Digits = "1234567890"
// init our charset for password
const Charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	Symbols + Digits

// init seededRand
var SeededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset get length(int) & charset(str) => return generated random password string
func StringWithCharset(length int, Charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[seededRand.Intn(len(Charset))]
	}
	return string(b)
}

// String - wrapper for StringWithCharset. Get length(int) => return StringWithCharset
func String(length int) string {
	return StringWithCharset(length, Charset)
}
