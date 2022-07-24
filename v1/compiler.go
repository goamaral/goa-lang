package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Goamaral/goa-lang/v1/codegen"
	"github.com/Goamaral/goa-lang/v1/parser"
	"github.com/Goamaral/goa-lang/v1/parser_yacc"
)

func main() {
	var stopAtLexer, stopAtAst, stopAtCodegen, inDebugMode bool
	outputFilePath := "./out/output.go"

	flag.BoolVar(&stopAtLexer, "lexer", false, "Stops compiler after lexer is complete")
	flag.BoolVar(&stopAtAst, "ast", false, "Stops compiler after ast is complete")
	flag.BoolVar(&stopAtCodegen, "codegen", false, "Stops compiler after codegen is complete and outputs to stdout")
	flag.BoolVar(&inDebugMode, "debug", false, "Enable debug mode for more verbose output")
	flag.Parse()

	inDebugMode = stopAtLexer || stopAtAst || stopAtCodegen || inDebugMode

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
	lexer := parser_yacc.NewLexer(string(sourceCodeBytes))
	lexer.Parse()
	if inDebugMode {
		lexer.Print()
	}
	if stopAtLexer {
		return
	}

	// Building syntax tree
	syntaxTree, ok := parser.BuildAst(parser_yacc.NewYaccLexer(lexer), inDebugMode)
	if inDebugMode && ok {
		syntaxTree.Print()
	}
	if !ok || stopAtAst {
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

	codegen.Generate(syntaxTree, outputFile)
}
