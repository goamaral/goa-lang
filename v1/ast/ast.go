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
	a.Root.Print(0)
	fmt.Println()
}
