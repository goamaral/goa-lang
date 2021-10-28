package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Goamaral/goa-lang/v1/parser"
)

func main() {
	// Reading source code from file to stdin
	var sourceCodeBytes []byte
	var err error
	if len(os.Args) == 1 {
		sourceCodeBytes, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error reading stdin")
			return
		}
	} else {
		sourceCodeBytes, err = ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Printf("Error reading %s\n", os.Args[1])
			return
		}
	}

	parser.Parse(string(sourceCodeBytes))
}
