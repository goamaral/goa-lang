package internal

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Goamaral/goa-lang/v1/internal/ast"
	"github.com/Goamaral/goa-lang/v1/internal/token"
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

func (cg *codeGenerator) ProcessDatatype(dataType *ast.Node) {
	switch dataType.Token.Kind {
	case token.BOOL_PTR:
		cg.writer.WriteString("*bool")
	case token.BOOL:
		cg.writer.WriteString("bool")
	case token.STRING_PTR:
		cg.writer.WriteString("*string")
	case token.STRING:
		cg.writer.WriteString("string")
	default:
		panic(fmt.Sprintf("Data type %s not supported", dataType.String()))
	}
}

func (cg *codeGenerator) ProcessChildren(n *ast.Node, seperator string) {
	for i, childNode := range n.Children {
		cg.ProcessNode(childNode)
		if i != len(n.Children)-1 {
			cg.writer.WriteString(seperator)
		}
	}
}

func (cg *codeGenerator) ProcessNode(n *ast.Node) {
	if n.IsTerminal() {
		cg.writer.WriteString(n.Value)
		return
	}

	switch n.Kind {
	case ast.Package:
		cg.ProcessPackage(n)
	case ast.FuncDef:
		cg.ProcessFuncDef(n)
	case ast.Block:
		cg.ProcessBlock(n)
	case ast.GoaFuncCall:
		cg.ProcessGoaFuncCall(n)
	case ast.FuncCallArgs:
		cg.ProcessFuncCallArgs(n)
	case ast.VarDecl:
		cg.ProcessVarDecl(n)
	default:
		panic(fmt.Sprintf("Node %s not supported", n.Kind.String()))
	}
}

func (cg *codeGenerator) ProcessAst(syntaxTree *ast.Ast) {
	cg.ProcessPackage(syntaxTree.Package)
	cg.writer.Flush()
}

func (cg *codeGenerator) ProcessPackage(pkg *ast.Node) {
	cg.WriteIdentation()
	cg.writer.WriteString("package main\n\n")
	cg.ProcessChildren(pkg, "")
}

func (cg *codeGenerator) ProcessFuncDef(funcDef *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprintf(cg.writer, "func %s() ", funcDef.Value)
	cg.ProcessChildren(funcDef, "")
}

func (cg *codeGenerator) ProcessBlock(block *ast.Node) {
	cg.WriteIdentation()
	cg.writer.WriteString("{\n")
	cg.identation++
	cg.ProcessChildren(block, "")
	cg.identation--
	cg.writer.WriteString("}\n")
}

func (cg *codeGenerator) ProcessGoaFuncCall(goaFuncCall *ast.Node) {
	cg.WriteIdentation()
	cg.writer.WriteString(strings.ToLower(goaFuncCall.Value))
	oldIdentation := cg.identation
	cg.identation = 0
	cg.ProcessChildren(goaFuncCall, "")
	cg.identation = oldIdentation
	cg.writer.WriteByte('\n')
}

func (cg *codeGenerator) ProcessFuncCallArgs(funcCallArgs *ast.Node) {
	cg.WriteIdentation()
	cg.writer.WriteByte('(')
	cg.ProcessChildren(funcCallArgs, ", ")
	cg.writer.WriteByte(')')
}

func (cg *codeGenerator) ProcessVarDecl(varDecl *ast.Node) {
	cg.WriteIdentation()
	fmt.Fprintf(cg.writer, "var %s ", varDecl.Value)
	cg.ProcessDatatype(varDecl.DataType)
	cg.writer.WriteByte('\n')
}
