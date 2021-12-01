package codegen

import (
	"bufio"
	"os"

	"github.com/Goamaral/goa-lang/v1/ast"
)

func Generate(syntaxTree *ast.Ast, outputFile *os.File) {
	if outputFile == nil {
		outputFile = os.Stdout
	}
	writer := bufio.NewWriter(outputFile)

	writer.WriteString("package main;\n\n")

	syntaxTree.Root.CodeGen(writer)
	writer.Flush()
}
