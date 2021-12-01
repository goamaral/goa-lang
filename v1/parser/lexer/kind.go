package lexer

/* CONSTANTS */
const (
	DEF Kind = iota
	DO
	END
	ID
	RPAR
	LPAR
	HASH
)

var kind_nameMap = map[Kind]string{
	DEF:  "DEF",
	DO:   "DO",
	END:  "END",
	ID:   "ID",
	RPAR: "RPAR",
	LPAR: "LPAR",
	HASH: "HASH",
}

/* STRUCT */
type Kind int
