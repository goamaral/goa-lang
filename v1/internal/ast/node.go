package ast

import (
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

func (n *Node) String() string {
	str := n.Kind.String()

	if n.Properties != nil {
		str += fmt.Sprintf("(%s)", strings.Join(n.GetPropertiesValues(), ", "))
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

func (n *Node) GetPropertiesValues() (values []string) {
	for _, property := range n.Properties {
		values = append(values, property.Value)
	}

	return values
}
