package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Goamaral/goa-lang/v1/internal/ast"
)

type codeGenerator struct {
	writer     *bufio.Writer
	identation int
}

func GenerateCode(syntaxTree *ast.Ast, outputFile *os.File) {
	cg := codeGenerator{writer: bufio.NewWriter(outputFile)}

	if outputFile == os.Stdout {
		cg.writer.WriteString("===== CODEGEN =====\n")
	}

	cg.ProcessAst(syntaxTree)
}

func (cg *codeGenerator) WriteIdentation() {
	for i := 0; i < cg.identation; i++ {
		cg.writer.WriteByte('\t')
	}
}

func (cg *codeGenerator) ProcessChildren(n *ast.Node) {
	for _, childNode := range n.Children {
		cg.ProcessNode(childNode)
	}
}

func (cg *codeGenerator) ProcessNode(n *ast.Node) {
	switch n.Kind {
	case ast.Package:
		cg.ProcessPackage(n)
	case ast.FuncDef:
		cg.ProcessFuncDef(n)
	case ast.FuncDefBody:
		cg.ProcessFuncDefBody(n)
	case ast.GoaFuncCall:
		cg.ProcessGoaFuncCall(n)
	case ast.FuncCallArgs:
		cg.ProcessFuncCallArgs(n)
	default:
		panic(fmt.Sprintf("Don't knwo how to process node of kind %s", n.Kind.String()))
	}
}

func (cg *codeGenerator) ProcessAst(syntaxTree *ast.Ast) {
	cg.ProcessPackage(syntaxTree.Package)
	cg.writer.Flush()
}

func (cg *codeGenerator) ProcessPackage(pkg *ast.Node) {
	cg.WriteIdentation()
	cg.writer.WriteString("package main\n\n")
	cg.ProcessChildren(pkg)
}

func (cg *codeGenerator) ProcessFuncDef(funcDef *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprintf(cg.writer, "func %s() ", funcDef.Properties[0].Value)
	cg.ProcessChildren(funcDef)
}

func (cg *codeGenerator) ProcessFuncDefBody(funcDefBody *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprint(cg.writer, "{\n")
	cg.identation++
	cg.ProcessChildren(funcDefBody)
	cg.identation--
	fmt.Fprint(cg.writer, "}\n")
}

func (cg *codeGenerator) ProcessGoaFuncCall(goaFuncCall *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprint(cg.writer, strings.ToLower(goaFuncCall.Properties[0].Value))
	oldIdentation := cg.identation
	cg.identation = 0
	cg.ProcessChildren(goaFuncCall)
	cg.identation = oldIdentation
	cg.writer.WriteByte('\n')
}

func (cg *codeGenerator) ProcessFuncCallArgs(funcCallArgs *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprintf(cg.writer, "(%s)", strings.Join(funcCallArgs.GetPropertiesValues(), ", "))
}
