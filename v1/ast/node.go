package ast

import (
	"bufio"
	"fmt"
	"strings"
)

/* CONSTANTS */
var kind_withValuesMap = map[Kind]bool{
	FuncDef:      true,
	GoaFuncCall:  true,
	FuncCallArgs: true,
}

/* STRUCT */
type Node struct {
	Kind       Kind
	Properties []Node
	Children   []Node
	Value      string
}

/* FUNCTIONS */
func NewNode(kind Kind, properties []Node, children []Node) Node {
	return Node{Kind: kind, Properties: properties, Children: children}
}

/* METHODS */
func (n *Node) String() string {
	if kind_withValuesMap[n.Kind] {
		return fmt.Sprintf("%s(%s)", n.Kind.String(), strings.Join(n.GetPropertiesValues(false), ", "))
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

func (n *Node) GetPropertiesValues(codegen bool) (values []string) {
	for _, property := range n.Properties {
		value := property.Value
		if codegen {
			switch property.Kind {
			}
		}

		values = append(values, value)
	}

	return values
}

/* METHODS - CODEGEN */
func (n *Node) CodeGen(writer *bufio.Writer, identation int) {
	closingString := ""

	for i := 0; i < identation; i++ {
		fmt.Fprint(writer, "\t")
	}

	switch n.Kind {
	case Prog:
		writer.WriteString("package main\n\n")

	case FuncDef:
		fmt.Fprintf(writer, "func %s() ", n.Properties[0].Value)

	case FuncDefBody:
		fmt.Fprint(writer, "{\n")
		closingString = "}\n"
		identation += 1

	case GoaFuncCall:
		fmt.Fprint(writer, strings.ToLower(n.Properties[0].Value))
		identation = 0
		closingString = "\n"

	case FuncCallArgs:
		fmt.Fprintf(writer, "(%s)", strings.Join(n.GetPropertiesValues(true), ", "))

	default:
		fmt.Fprintf(writer, "UNKNOWN: %s\n", n.Kind.String())
	}

	for _, childNode := range n.Children {
		childNode.CodeGen(writer, identation)
	}

	fmt.Fprint(writer, closingString)
}
