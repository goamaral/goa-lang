package codegen

import (
	"fmt"

	"github.com/Goamaral/goa-lang/v1/ast"
)

var code string

func Generate(syntaxTree *ast.Ast) (code string) {
	code = "package main;\n"
	consumeNode(syntaxTree.Root)

	return code
}

func consumeNode(node *ast.Node) {
	switch node.Kind {
	case ast.FuncDef:
		code = fmt.Sprintf("%sfunc %s() {\n}\n", code, node.Value)
	}

	for _, childNode := range node.Children {
		consumeNode(childNode)
	}
}
