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

	syntaxTree.Root.CodeGen(writer, 0)
	writer.Flush()
}
