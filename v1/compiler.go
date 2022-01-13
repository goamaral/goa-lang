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
	var stopAtLexer, stopAtAst bool
	flag.BoolVar(&stopAtLexer, "lexer", false, "Stops compiler after lexer is complete")
	flag.BoolVar(&stopAtAst, "ast", false, "Stops compiler after ast is complete")
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
	codegen.Generate(&syntaxTree, nil)
}
