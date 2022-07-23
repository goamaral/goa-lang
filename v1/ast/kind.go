package ast

const (
	// NON TERMINALS
	Package Kind = iota
	FuncDef
	FuncDefBody
	GoaFuncCall
	FuncCallArgs

	// TERMINALS
	Id
	Boolean
	String
	Integer
	Nil
)

var kind_nameMap = map[Kind]string{
	Package:      "PACKAGE",
	FuncDef:      "FUNC_DEF",
	FuncDefBody:  "FUNC_DEF_BODY",
	GoaFuncCall:  "GOA_FUNC_CALL",
	FuncCallArgs: "FUNC_CALL_ARGS",
}

type Kind int

func (k *Kind) String() string {
	return kind_nameMap[*k]
}
