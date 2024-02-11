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
	Bool
	Str
	Int
	Nil
	DataType
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
	Id:       "ID",
	Bool:     "BOOL",
	Str:      "STR",
	Int:      "INT",
	Nil:      "NIL",
	DataType: "DATA_TYPE",
}

var terminalKinds = []Kind{Id, Bool, Str, Int, Nil, DataType}

type Kind int

func (k *Kind) String() string {
	return kind_nameMap[*k]
}
