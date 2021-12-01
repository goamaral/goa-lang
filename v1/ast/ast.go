package ast

import "fmt"

type Ast struct {
	Root Node
}

func New() Ast {
	return Ast{
		Root: Node{Kind: Prog},
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
	newNode := Node{Kind: FuncDef, Value: name}
	a.Root.addChild(newNode)
}
