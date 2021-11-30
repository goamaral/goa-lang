package ast

import "fmt"

type NodeKind int

const (
	// NON TERMINALS
	Prog NodeKind = iota
	FuncDef

	// TERMINALS
	Id
)

var nodeKind_nameMap = map[NodeKind]string{
	Prog:    "PROG",
	FuncDef: "FUNC_DEF",
	Id:      "ID",
}

/* AST */
type Ast struct {
	Root *Node
}

func New() Ast {
	return Ast{
		Root: &Node{Kind: Prog},
	}
}

func (a *Ast) Print() {
	fmt.Println("===== AST =====")
	fmt.Println(a.Root.String())

	for _, node := range a.Root.Children {
		fmt.Println(node.String())
	}

	fmt.Println()
}

func (a *Ast) AddFuncDef(name string) {
	newNode := &Node{Kind: FuncDef, Value: name}
	a.Root.addChild(newNode)
}

/* NODE */
type Node struct {
	Kind     NodeKind
	Value    string
	Children []*Node
}

func (n *Node) String() string {
	switch n.Kind {
	case FuncDef:
		return fmt.Sprintf("%s(%s)", nodeKind_nameMap[n.Kind], n.Value)
	default:
		return nodeKind_nameMap[n.Kind]
	}
}

func (n *Node) addChild(child *Node) {
	n.Children = append(n.Children, child)
}
