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
	// NON TERMINALS
	Package:      "PACKAGE",
	FuncDef:      "FUNC_DEF",
	FuncDefBody:  "FUNC_DEF_BODY",
	GoaFuncCall:  "GOA_FUNC_CALL",
	FuncCallArgs: "FUNC_CALL_ARGS",

	// TERMINALS
	Id:      "ID",
	Boolean: "BOOLEAN",
	String:  "STRING",
	Integer: "INTEGER",
	Nil:     "NIL",
}

type Kind int

func (k *Kind) String() string {
	return kind_nameMap[*k]
}
