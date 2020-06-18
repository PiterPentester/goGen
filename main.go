package main

import (
	"flag"
	"fmt"

	abr "github.com/PiterPentester/goGen/abracadabra"
	mem "github.com/PiterPentester/goGen/memorable"
)

func goGen() {
	// define flags
	genMem := flag.Bool("mem", false, "generate memorable password")
	genAbr := flag.Bool("abr", false, "generate abracadabra password")
	flag.Parse()
	if *genMem {
		// get words
		wrds, _ := mem.GetRandWords(3)
		// generating memorable password
		fmt.Println("Memorable password:", mem.GenMemorablePass(wrds))
	} else if *genAbr {
		// generating abracadabra password
		fmt.Println("Abracadabra password:", abr.String(16))
	} else {
		// print help
		fmt.Println("Choose your pass:")
		fmt.Println("./goGen -mem (return memorable password)")
		fmt.Println("./goGen -abr (return abracadabra password)")
	}
}

func main() {
	goGen()
}
