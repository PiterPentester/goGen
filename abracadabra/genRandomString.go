package abracadabra

import (
	"math/rand"
	"time"
)

// Symbols - var for different symbols in our Charset. Contains most frequent symbols.
const Symbols = "~!@#$%^&*()_{}<>"
// Digits - var for all digits in our Charset
const Digits = "1234567890"
// Charset for password
const Charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	Symbols + Digits

// SeededRand - init seed from time
var SeededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset - get length(int) & charset(str) => return generated random password string
func StringWithCharset(length int, Charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[SeededRand.Intn(len(Charset))]
	}
	return string(b)
}

// String - wrapper for StringWithCharset. Get length(int) => return StringWithCharset
func String(length int) string {
	return StringWithCharset(length, Charset)
}
