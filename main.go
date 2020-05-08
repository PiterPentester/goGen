package main

import (
	"fmt"
	"math/rand"
	"time"
)

// init our charset for password
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	"~!@#$%^&*()_{}<>"

// init seededRand
var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset get len(int) & charset(str) => return generated password string
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// String - wrapper for StringWithCharset. Get len(int) & return StringWithCharset
func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	fmt.Println("1:", String(12))
	//fmt.Println("2:", rand.Int())
	//fmt.Println("3:", rand.Int())
}
