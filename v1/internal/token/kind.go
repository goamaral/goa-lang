package token

/* CONSTANTS */
const (
	UNKNOWN Kind = iota

	// Keywords
	DEF
	DO
	END

	// Symbols
	RPAR
	LPAR
	HASH
	COMMA
	PLUS
	MINUS

	// Datatypes
	BOOL_PTR
	BOOL
	INT_PTR
	INT
	STR_PTR
	STR

	// Untyped constants
	BOOL_LIT
	INT_LIT
	STR_LIT
	NIL

	// Id
	UPPER_ID
	LOWER_ID
)

var kindToString = map[Kind]string{
	UNKNOWN: "UNKNOWN",

	// Keywords
	DEF: "DEF",
	DO:  "DO",
	END: "END",

	// Symbols
	RPAR:  "RPAR",
	LPAR:  "LPAR",
	HASH:  "HASH",
	COMMA: "COMMA",
	PLUS:  "PLUS",
	MINUS: "MINUS",

	// Datatypes
	BOOL_PTR: "BOOL_PTR",
	BOOL:     "BOOL",
	INT_PTR:  "INT_PTR",
	INT:      "INT",
	STR_PTR:  "STR_PTR",
	STR:      "STR",

	// Untyped constants
	BOOL_LIT: "BOOL_LIT",
	INT_LIT:  "INT_LIT",
	STR_LIT:  "STR_LIT",
	NIL:      "NIL",

	// Id
	UPPER_ID: "UPPER_ID",
	LOWER_ID: "LOWER_ID",
}

/* STRUCT */
type Kind int

/* METHODS */
func (k *Kind) String() string {
	return kindToString[*k]
}
