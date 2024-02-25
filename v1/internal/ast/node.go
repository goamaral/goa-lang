package ast

import (
	"fmt"
	"slices"
	"strings"

	"github.com/Goamaral/goa-lang/v1/internal/token"
)

type Node struct {
	Kind     Kind
	NPops    uint
	Value    []byte
	DataType *Node

	Children []*Node
	Token    token.Token
}

func NewNode(kind Kind, nPops uint) *Node {
	return &Node{Kind: kind, NPops: nPops}
}

func (n *Node) String() string {
	if n.Kind == DataType {
		return n.Token.String()
	}

	var sb strings.Builder
	sb.WriteString(n.Kind.String())

	if len(n.Value) > 0 {
		sb.WriteString(fmt.Sprintf("(%s)", n.Value))
	}

	if n.DataType != nil {
		sb.WriteString(fmt.Sprintf("[%s]", n.DataType.String()))
	}

	return sb.String()
}

func (n *Node) Print(identation int) {
	for i := 0; i < identation; i += 1 {
		fmt.Print("  ")
	}

	fmt.Println(n.String())
	identation += 1

	for _, childNode := range n.Children {
		childNode.Print(identation)
	}
}

func (n *Node) AddValue(value []byte) *Node {
	n.Value = value
	return n
}

func (n *Node) AddDataType(dataType *Node) *Node {
	n.DataType = dataType
	return n
}

func (n *Node) AddChildren(children ...*Node) *Node {
	n.Children = append(n.Children, children...)
	return n
}

func (n *Node) AddToken(tk token.Token) *Node {
	n.Token = tk
	return n
}

func (n *Node) IsTerminal() bool {
	return slices.Contains(terminalKinds, n.Kind)
}
