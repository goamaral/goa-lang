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

	// Datatypes
	BOOL_PTR
	BOOL

	// Untyped constants
	TRUE
	FALSE
	INTEGER_LIT
	STRING_LIT
	NIL

	// Id
	UPPER_ID
	LOWER_ID
)

var kind_nameMap = map[Kind]string{
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

	// Datatypes
	BOOL_PTR: "BOOL_PTR",
	BOOL:     "BOOL",

	// Untyped constants
	TRUE:        "TRUE",
	FALSE:       "FALSE",
	INTEGER_LIT: "INTEGER_LIT",
	STRING_LIT:  "STRING_LIT",
	NIL:         "NIL",

	// Id
	UPPER_ID: "UPPER_ID",
	LOWER_ID: "LOWER_ID",
}

/* STRUCT */
type Kind int

/* METHODS */
func (k *Kind) String() string {
	return kind_nameMap[*k]
}
