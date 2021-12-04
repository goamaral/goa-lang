package ast

import (
	"bufio"
	"fmt"
	"strings"
)

/* CONSTANTS */
var kind_withNameMap = map[Kind]bool{
	FuncDef:     true,
	GoaFuncCall: true,
}

/* STRUCT */
type Node struct {
	Kind     Kind
	Value    string
	Children []Node
}

/* FUNCTIONS */
func NewNode(kind Kind, value string, children []Node) Node {
	return Node{Kind: kind, Value: value, Children: children}
}

/* METHODS */
func (n *Node) String() string {
	if kind_withNameMap[n.Kind] {
		return fmt.Sprintf("%s(%s)", n.Kind.String(), n.Value)
	} else {
		return n.Kind.String()
	}
}

func (n *Node) Print(identation int) {
	for i := 0; i < identation; i += 1 {
		fmt.Print("..")
	}

	fmt.Println(n.String())
	identation += 1

	for _, childNode := range n.Children {
		childNode.Print(identation)
	}
}

func (n *Node) AddChild(child Node) {
	n.Children = append(n.Children, child)
}

/* METHODS - CODEGEN */
func (n *Node) CodeGen(writer *bufio.Writer, identation int) {
	closingString := ""

	for i := 0; i < identation; i++ {
		fmt.Fprint(writer, "  ")
	}

	switch n.Kind {
	case Prog:
		writer.WriteString("package main\n\n")

	case FuncDef:
		fmt.Fprintf(writer, "func %s() ", n.Value)

	case FuncDefBody:
		fmt.Fprint(writer, "{\n")
		closingString = "}\n"
		identation += 1

	case GoaFuncCall:
		fmt.Fprintf(writer, "%s()\n", strings.ToLower(n.Value))

	default:
		fmt.Fprintf(writer, "NODE: %s\n", n.Kind.String())
	}

	for _, childNode := range n.Children {
		childNode.CodeGen(writer, identation)
	}

	fmt.Fprint(writer, closingString)
}
