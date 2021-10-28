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
	root Node
}

func NewAst() Ast {
	return Ast{
		root: Node{
			kind: Prog,
		},
	}
}

func (a Ast) Print() {
	fmt.Println("===== AST =====")
	fmt.Println(a.root.String())

	for _, node := range a.root.children {
		fmt.Println(node.String())
	}

	fmt.Println()
}

func (a *Ast) AddFuncDef(name string) {
	newNode := &Node{kind: FuncDef, value: name}
	a.root.addChild(newNode)
}

/* NODE */
type Node struct {
	kind     NodeKind
	value    string
	children []*Node
}

func (n Node) String() string {
	switch n.kind {
	case FuncDef:
		return fmt.Sprintf("%s(%s)", nodeKind_nameMap[n.kind], n.value)
	default:
		return nodeKind_nameMap[n.kind]
	}
}

func (n *Node) addChild(child *Node) {
	n.children = append(n.children, child)
}
