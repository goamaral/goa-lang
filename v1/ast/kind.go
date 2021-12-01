package ast

const (
	// NON TERMINALS
	Prog Kind = iota
	FuncDef

	// TERMINALS
	Id
)

var kind_nameMap = map[Kind]string{
	Prog:    "PROG",
	FuncDef: "FUNC_DEF",
	Id:      "ID",
}

type Kind int
