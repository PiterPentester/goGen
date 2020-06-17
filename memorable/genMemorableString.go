package memorable

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/PiterPentester/goGen/abracadabra"
)

// init our client for curl requests
var client http.Client

// ErrBadCode - handle not "OK" http status code
var ErrBadCode = errors.New("bad status code")

// ErrEmptyResult - handle empty result for parsing functions
var ErrEmptyResult = errors.New("function return empty result")

// ErrBadNumOfWords - handle 0 & < 3 input
var ErrBadNumOfWords = errors.New("bad num of words")

// CurlToAddr - makes a curl request to the site & save the output as string
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

// getTextOut - cut text after "text_out" and store it to txt var
func getTextOut(txt string) (string, error) {
	if txt == "" {
		return "", ErrEmptyResult
	}
	// text begins after "text_out"
	f := "text_out"
	i := strings.Index(txt, f)
	if i > -1 {
		txt = txt[i+len(f)+3 : len(txt)-2]
	}
	return txt, nil
}

// splitByTag - split our txt var by "<p>", remove ".<\/p>\r" from res and return list of lines.
func splitByTag(txt string) ([]string, error) {
	if txt == "" {
		return []string{}, ErrEmptyResult
	}
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
	return res, nil
}

// splitBySpace - split our list of lines into a list of words and return it.
func splitBySpace(res []string) ([]string, error) {
	var words []string

	if len(res) == 0 {
		return words, ErrEmptyResult
	}

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
	return words, nil
}

// parseOutput - wrapper for parsing functions, return list of words
func parseOutput(addr string) ([]string, error) {
	var words []string
	txt, err := CurlToAddr(addr)
	if err != nil {
		return words, err
	}

	txt, err = getTextOut(txt)
	if err != nil {
		return words, err
	}

	res, _ := splitByTag(txt)
	if len(res) == 0 {
		return words, ErrEmptyResult
	}

	words, err = splitBySpace(res)
	if err != nil {
		return words, err
	}

	return words, nil
}

// genOfflineWords - in case we have no access to internet for get random words from site
func genOfflineWords(numOfWords int) ([]string, error) {
	if numOfWords < 3 {
		return []string{}, ErrBadNumOfWords
	}
	res := make([]string, numOfWords)
	for i := range res {
		w := ""
		for j := 0; j < 2; j++ {
			c := abracadabra.ConsonantsLower[abracadabra.SeededRand.Intn(len(abracadabra.ConsonantsLower))]
			v := abracadabra.VowelsLower[abracadabra.SeededRand.Intn(len(abracadabra.VowelsLower))]
			w += string(c) + string(v)
		}
		res[i] = w
	}
	return res, nil
}

// GetRandWords - get number of wanted words & create list of random choosen words with predefined length
// GetRandWords(3, words) => [word1, word2, word3]
func GetRandWords(numOfWords int) ([]string, error) {
	if numOfWords < 3 {
		return []string{}, ErrBadNumOfWords
	}
	res := make([]string, numOfWords)
	words, err := parseOutput("http://www.randomtext.me/api/gibberish")
	if err != nil {
		words, _ = genOfflineWords(numOfWords)
		for i := range res {
			res[i] = strings.Title(words[i])
		}
		return res, nil
	}
	for i := range res {
		res[i] = strings.Title(words[abracadabra.SeededRand.Intn(len(words))])
	}
	return res, nil
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

// GenMemorablePass - main function for generating memorable passwords
func GenMemorablePass(words []string) string {
	res := ""
	for _, w := range words {
		//d := abracadabra.Digits[abracadabra.SeededRand.Intn(len(abracadabra.Digits))]
		randomNum := random(10, 99)
		str := strconv.Itoa(randomNum)
		res += w + string(str)
	}
	smbl := abracadabra.Symbols[abracadabra.SeededRand.Intn(len(abracadabra.Symbols))]
	res += string(smbl)
	return res
}
