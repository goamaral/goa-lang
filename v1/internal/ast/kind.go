package ast

const (
	// NON TERMINALS
	Package Kind = iota
	FuncDef
	Block
	GoaFuncCall
	FuncCallArgs
	VarDecl

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
	Block:        "BLOCK",
	GoaFuncCall:  "GOA_FUNC_CALL",
	FuncCallArgs: "FUNC_CALL_ARGS",
	VarDecl:      "VAR_DECL",

	// TERMINALS
	Id:      "ID",
	Boolean: "BOOLEAN",
	String:  "STRING",
	Integer: "INTEGER",
	Nil:     "NIL",
}

var terminalKinds = []Kind{Id, Boolean, String, Integer, Nil}

type Kind int

func (k *Kind) String() string {
	return kind_nameMap[*k]
}