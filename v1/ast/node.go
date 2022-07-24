package ast

import (
	"bufio"
	"fmt"
	"strings"
)

type Node struct {
	Kind       Kind
	Properties []*Node
	Children   []*Node
	Value      string
}

func NewNode(kind Kind) *Node {
	return &Node{Kind: kind}
}

func NewSimpleNode(kind Kind, value string) *Node {
	return &Node{Kind: kind, Value: value}
}

func NewComplexNode(kind Kind, properties []*Node, children []*Node) *Node {
	return &Node{Kind: kind, Properties: properties, Children: children}
}

func (n *Node) String() string {
	str := n.Kind.String()

	if n.Properties != nil {
		str += fmt.Sprintf("(%s)", strings.Join(n.GetPropertiesValues(false), ", "))
	} else if n.Value != "" {
		str += fmt.Sprintf("(%s)", n.Value)
	}

	return str
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

func (n *Node) AddValue(value string) *Node {
	n.Value = value
	return n
}

func (n *Node) AddChildren(children ...*Node) *Node {
	n.Children = append(n.Children, children...)
	return n
}

func (n *Node) AddProperties(properties ...*Node) *Node {
	n.Properties = append(n.Properties, properties...)
	return n
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

func (n *Node) CodeGen(writer *bufio.Writer, identation int) {
	closingString := ""

	for i := 0; i < identation; i++ {
		fmt.Fprint(writer, "\t")
	}

	switch n.Kind {
	case Package:
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
