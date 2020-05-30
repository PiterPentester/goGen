package main

import (
	"flag"
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

func goGen() {
	genMem := flag.Bool("mem", false, "generate memorable password")
	genAbr := flag.Bool("abr", false, "generate abracadabra password")
	flag.Parse()
	if *genMem {
		fmt.Println("Memorable password:", mem.GenMemorablePass(mem.GetRandWords(3)))
	} else if *genAbr {
		fmt.Println("Abracadabra password:", abr.String(10))
	} else {
		fmt.Println("Choose your pass:")
		fmt.Println("./goGen -mem (return memorable password)")
		fmt.Println("./goGen -abr (return abracadabra password)")
	}
}

func main() {
	banner()
	goGen()
}
