package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Goamaral/goa-lang/v1/codegen"
	"github.com/Goamaral/goa-lang/v1/parser"
	"github.com/Goamaral/goa-lang/v1/parser/lexer"
)

func main() {
	var stopAtLexer, stopAtAst, stopAtCodegen bool
	outputFilePath := "./out/output.go"

	flag.BoolVar(&stopAtLexer, "lexer", false, "Stops compiler after lexer is complete")
	flag.BoolVar(&stopAtAst, "ast", false, "Stops compiler after ast is complete")
	flag.BoolVar(&stopAtCodegen, "codegen", false, "Stops compiler after codegen is complete and outputs to stdout")
	flag.Parse()

	// Reading source code from file to stdin
	var sourceCodeBytes []byte
	var err error

	sourceFileLocation := os.Args[len(os.Args)-1]
	sourceCodeBytes, err = ioutil.ReadFile(sourceFileLocation)
	if err != nil {
		fmt.Printf("Error reading %s\n", sourceFileLocation)
		return
	}

	// Lexing
	lex := lexer.New(string(sourceCodeBytes))
	lex.Parse()
	lex.Print()
	if stopAtLexer {
		return
	}

	// Building syntax tree
	syntaxTree := parser.Parse(&lex)
	if stopAtAst {
		return
	}

	// Generate code and write to file
	var outputFile *os.File
	if stopAtCodegen {
		outputFile = os.Stdout
	} else {
		_, err := os.Stat("out")
		if os.IsNotExist(err) {
			err := os.Mkdir("out", 0755)
			if err != nil {
				fmt.Println("Error creating out folder")
				return
			}
		}

		outputFile, err = os.Create(outputFilePath)
		if err != nil {
			fmt.Printf("Error writting to %s\n", outputFilePath)
		}
		defer outputFile.Close()
	}

	codegen.Generate(&syntaxTree, outputFile)
}
