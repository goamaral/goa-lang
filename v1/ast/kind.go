package ast

const (
	// NON TERMINALS
	Prog Kind = iota
	FuncDef
	FuncDefBody
	GoaFuncCall
	FuncCallArgs

	// TERMINALS
	Id
	Boolean
	String
	Integer
)

var kind_nameMap = map[Kind]string{
	Prog:         "PROG",
	FuncDef:      "FUNC_DEF",
	FuncDefBody:  "FUNC_DEF_BODY",
	GoaFuncCall:  "GOA_FUNC_CALL",
	FuncCallArgs: "FUNC_CALL_ARGS",
}

type Kind int

func (k *Kind) String() string {
	return kind_nameMap[*k]
}
