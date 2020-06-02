package abracadabra

import (
	"math/rand"
	"time"
)

// Symbols - different symbols in our Charset. Contains most frequent symbols.
const Symbols = "~!@#$%^&*()_{}<>"

// Digits - all digits in our Charset
const Digits = "1234567890"

// VowelsUpper - upper vowels in language
const VowelsUpper = "AEIOUY"

// VowelsLower - lower vowels in language
const VowelsLower = "aeiouy"

// ConsonantsUpper - upper consonants in language
const ConsonantsUpper = "BCDFGHJKLMNPQRSTVWXZ"

// ConsonantsLower - lower consonants in language
const ConsonantsLower = "bcdfghjklmnpqrstvwxz"

// Charset for password
const Charset = VowelsUpper + VowelsLower + ConsonantsUpper + ConsonantsLower + Symbols + Digits

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
