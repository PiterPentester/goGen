package main

import (
	"fmt"
	"math/rand"
	"time"
	"net/http"
	//"errors"
	"io/ioutil"
	"strings"
)

const symbols = "~!@#$%^&*()_{}<>"
const digits = "1234567890"
// init our charset for password
const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" +
	symbols + digits

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

var client http.Client
// CurlToAddr - make curl request to site & save output
func CurlToAddr(addr string) string {
    resp, err := client.Get(addr)
    if err != nil {
	    return ""
    }
    defer resp.Body.Close()
    
    if resp.StatusCode == http.StatusOK {
        bodyBytes, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return ""
        }
        bodyString := string(bodyBytes)
        return bodyString
    }
    return ""
}

// ParseOutput get txt string & return list of words
func ParseOutput(txt string) []string {
	
	// text begins after "text_out"
	f := "text_out"
	i := strings.Index(txt, f)
	if i > -1 {
	    txt = txt[i+len(f)+3:len(txt)-2]
	}
	
	// split txt by <p> tag
	res := strings.Split(txt, "<p>")
	for id, elem := range res {
		// remove ".<\/p>\r" from res
		i = strings.Index(elem, ".")
		if i > -1 {
			elem = elem[:i]
	    }
	    res[id] = elem
	}
	
	var words []string
	
	// split res by " "
	for _, e := range res {
	    w := strings.Split(e, " ")
	    for _, word := range w {
			if len(word) < 3 {
			    continue
			}
			
		    words = append(words, word)
		}
	}
	
    return words
}

//
func GetRandWords(numOfWords int, words []string) []string {
    res := make([]string, numOfWords)
    for i := range res {
		res[i] = strings.Title(words[seededRand.Intn(len(words))])
	}
	return res
}

func GenMemorablePass(words []string) string {
    res := ""
    for _, w := range words {
		smbl := symbols[seededRand.Intn(len(symbols))]
	    res += w + string(smbl)
	}
	d := digits[seededRand.Intn(len(digits))]
	res += string(d)
	return res 
}

func main() {
	fmt.Println("1:", String(12))
	words := ParseOutput(CurlToAddr("http://www.randomtext.me/api/gibberish"))
	
	data := GetRandWords(3, words)
	fmt.Println("2:", data)
	fmt.Println("3:", GenMemorablePass(data))
}
