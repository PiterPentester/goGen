package abracadabra

import (
	"math/rand"
	"time"
)

// ConsonantsLower - lower consonants in language
const ConsonantsLower = "bcdfghjklmnpqrstvwxz"

// VowelsLower - lower vowels in language
const VowelsLower = "aeiouy"

// ConsonantsUpper - upper consonants in language
const ConsonantsUpper = "BCDFGHJKLMNPQRSTVWXZ"

// VowelsUpper - upper vowels in language
const VowelsUpper = "AEIOUY"

// Symbols - different symbols in our Charset. Contains most frequent symbols.
const Symbols = "~!@#$%^&*()_{}<>"

// Digits - all digits in our Charset
const Digits = "1234567890"

// Upper - upper letters in language
const Upper = "AEIOUYBCDFGHJKLMNPQRSTVWXZ"

// Lower - lower letters in language
const Lower = "aeiouybcdfghjklmnpqrstvwxz"

// Charset for password
const Charset = Upper + Lower + Symbols + Digits

// SeededRand - init seed from time
var SeededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset - get length(int) & charset(str) => return generated random password string
func StringWithCharset(length int, Charset string) string {
	b := make([]byte, length)
	var l, u, d, s int = 0, 0, 0, 0
	for j := 0; j < len(b); {
		//generate random i
		i := SeededRand.Intn(len(b))
		//insert Lower
		if i%2 == 0 && l < 2 {
			b[j] = Lower[SeededRand.Intn(len(Lower))]
			l++
			j++
		} else if i%5 == 0 && u < 2 {
			//insert Upper
			b[j] = Upper[SeededRand.Intn(len(Upper))]
			u++
			j++
		} else if i%3 == 0 && d < 2 {
			//insert Digits
			b[j] = Digits[SeededRand.Intn(len(Digits))]
			d++
			j++
		} else if i%1 == 0 && s < 2 {
			//insert Symbols
			b[j] = Symbols[SeededRand.Intn(len(Symbols))]
			s++
			j++
		} else {
			//insert random
			b[j] = Charset[SeededRand.Intn(len(Charset))]
			j++
		}
	}
	if len(string(b)) > 7 {
		return string(b)
	}
	return String(16)
}

// String - wrapper for StringWithCharset. Get length(int) => return StringWithCharset
func String(length int) string {
	return StringWithCharset(length, Charset)
}
