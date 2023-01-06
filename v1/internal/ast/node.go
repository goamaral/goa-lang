package ast

import (
	"fmt"
	"strings"

	"github.com/Goamaral/goa-lang/v1/internal/token"
	"golang.org/x/exp/slices"
)

type Node struct {
	Kind     Kind
	Value    string
	DataType token.Kind
	Children []*Node
}

func NewNode(kind Kind) *Node {
	return &Node{Kind: kind}
}

func (n *Node) String() string {
	var sb strings.Builder
	sb.WriteString(n.Kind.String())

	if n.Value != "" {
		sb.WriteString(fmt.Sprintf("(%s)", n.Value))
	}

	if n.DataType != token.UNKNOWN {
		sb.WriteString(fmt.Sprintf(" -> %s", n.DataType.String()))
	}

	return sb.String()
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

func (n *Node) AddDataType(dataType token.Kind) *Node {
	n.DataType = dataType
	return n
}

func (n *Node) AddChildren(children ...*Node) *Node {
	n.Children = append(n.Children, children...)
	return n
}

func (n *Node) IsTerminal() bool {
	return slices.Contains(terminalKinds, n.Kind)
}
