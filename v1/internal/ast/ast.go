package ast

import (
	"fmt"
)

type Ast struct {
	Package *Node
}

func New() *Ast {
	return &Ast{
		Package: &Node{Kind: Package},
	}
}

func (a *Ast) Print() {
	fmt.Println("===== AST =====")
	a.Package.Print(0)
	fmt.Println()
}
