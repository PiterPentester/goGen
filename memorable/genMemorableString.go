package memorable

import (
	"math/rand"
	"time"
	"net/http"
	"errors"
	"io/ioutil"
	"strings"
)

var client http.Client

// Init errors
var ErrBadCode = errors.New("bad status code")

// CurlToAddr - makes a curl request to the site & save the output
func CurlToAddr(addr string) (string, error) {
    resp, err := client.Get(addr)
    if err != nil {
	    return "", err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return "", ErrBadCode
    }
    
    bodyBytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }
    
    bodyString := string(bodyBytes)
    return bodyString, nil
}

func getTextOut(txt string) string {
	// text begins after "text_out"
	f := "text_out"
	i := strings.Index(txt, f)
	if i > -1 {
	    txt = txt[i+len(f)+3:len(txt)-2]
	}
	return txt
}

func splitByTag(txt string) []string {
    // split txt by <p> tag
	res := strings.Split(txt, "<p>")
	for id, elem := range res {
		// remove ".<\/p>\r" from res
		i := strings.Index(elem, ".")
		if i > -1 {
			elem = elem[:i]
	    }
	    res[id] = elem
	}
	return res
}

func splitBySpace(res []string) []string {
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

// ParseOutput get txt string & return list of words
func ParseOutput() (string, error) {
	txt, err := CurlToAddr("http://www.randomtext.me/api/gibberish")
	if err != nil {
	    return "", err
	}
	
    txt = getTextOut(txt)
    res := splitByTag(txt)
    words := splitBySpace(res)
    
}

// GetRandWords: numOfWords(int), words([]string) => []string 
// GetRandWords(3, words) => [word1, word2, word3] 
func GetRandWords(numOfWords int, words []string) []string {
    res := make([]string, numOfWords)
    for i := range res {
		res[i] = strings.Title(words[seededRand.Intn(len(words))])
	}
	return res
}

// GenMemorablePass 
func GenMemorablePass(words []string) string {
    res := ""
    for _, w := range words {
		smbl := symbols[seededRand.Intn(len(Symbols))]
	    res += w + string(smbl)
	}
	d := digits[seededRand.Intn(len(Digits))]
	res += string(d)
	return res 
}
