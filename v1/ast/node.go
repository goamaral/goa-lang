package ast

import (
	"bufio"
	"fmt"
)

type Node struct {
	Kind     Kind
	Value    string
	Children []Node
}

func (n *Node) String() string {
	switch n.Kind {
	case FuncDef:
		return fmt.Sprintf("%s(%s)", kind_nameMap[n.Kind], n.Value)
	default:
		return kind_nameMap[n.Kind]
	}
}

func (n *Node) CodeGen(writer *bufio.Writer) {
	switch n.Kind {
	case Prog:
		n.CodeGenChildren(writer)
	case FuncDef:
		fmt.Fprintf(writer, "func %s() {\n}\n", n.Value)
	}
}

func (n *Node) CodeGenChildren(writer *bufio.Writer) {
	for _, childNode := range n.Children {
		childNode.CodeGen(writer)
	}
}

func (n *Node) addChild(child Node) {
	n.Children = append(n.Children, child)
}
