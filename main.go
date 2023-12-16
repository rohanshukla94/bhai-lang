package main

import (
	"fmt"
	"flag"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	fmt.Println("hi")

	var inputFile string
	flag.StringVar(&inputFile, "file", "test.bhai", "The input file to be compiled")

	flag.Parse()

	fmt.Printf("Input File: %s\n", inputFile)

	dat, err := os.ReadFile(inputFile)
    check(err)
    fmt.Print(string(dat))
}

// func parseFile() {

// }