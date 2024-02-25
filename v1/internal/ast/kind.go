package ast

const (
	// NON TERMINALS
	Package Kind = iota
	FuncDef
	Block
	Stmt
	GoaFuncCall
	FuncCallArgList
	Expr
	ParExpr
	BinOp
	VarDecl
	Terminal
	UntypedConstant

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
	Package:         "PACKAGE",
	FuncDef:         "FUNC_DEF",
	Block:           "BLOCK",
	GoaFuncCall:     "GOA_FUNC_CALL",
	FuncCallArgList: "FUNC_CALL_ARG_LIST",
	Expr:            "EXPR",
	VarDecl:         "VAR_DECL",

	// TERMINALS
	Id: "ID",
	// TODO: Try to replace types below with Literal Kind
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
