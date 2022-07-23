package codegen

import (
	"bufio"
	"os"

	"github.com/Goamaral/goa-lang/v1/ast"
)

func Generate(syntaxTree *ast.Ast, outputFile *os.File) {
	writer := bufio.NewWriter(outputFile)

	if outputFile == os.Stdout {
		writer.WriteString("===== CODEGEN =====\n")
	}

	syntaxTree.Package.CodeGen(writer, 0)
	writer.Flush()
}
