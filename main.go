package main

import (
	"fmt"
	abr "github.com/PiterPentester/goGen/abracadabra"
	mem "github.com/PiterPentester/goGen/memorable"
)

func banner() {
    fmt.Println("goGen - simple passwords generator on golang")
	fmt.Println("")
	fmt.Println("Author @PiterPentester")
	fmt.Println("")
	fmt.Println("version 0.0.1alpha")
	fmt.Println("")
}

func checkLen(n int) bool {
    if n < 0 {
	    return false
	}
	return true
}

func goGen() {
    fmt.Println("What type of password do you want to generate?")
	fmt.Println("Choose num:")
	fmt.Println("1 - Abracadabra password")
	fmt.Println("2 - Memorable password")
	var answer, length int
	fmt.Scanf("%d", &answer)
	if answer == 1 {
	    fmt.Println("You choose: abracadabra")
	    fmt.Println("Enter length of password(>0):")
	    fmt.Scanf("%d", &length)
	    if checkLen(length) {
		    fmt.Println("Abracadabra password:", abr.String(length))
		} else {
		    fmt.Println("Are you stupid?!")
		}
	} else if answer == 2 {
	    fmt.Println("You choose: memorable")
	    fmt.Println("Enter num of words in password(>0):")
	    fmt.Scanf("%d", &length)
	    if checkLen(length) {
		    fmt.Println("Memorable password:", mem.GenMemorablePass(mem.GetRandWords(length)))
		} else {
		    fmt.Println("Are you stupid?!")
		}
	} else {
	    fmt.Println("Bad input!!!")
	}
}


func main() {
	banner()
	goGen()
}
