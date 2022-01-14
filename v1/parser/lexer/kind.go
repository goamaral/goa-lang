package lexer

/* CONSTANTS */
const (
	DEF Kind = iota
	DO
	END
	RPAR
	LPAR
	HASH
	COMMA
	UPPER_ID
	LOWER_ID
	TRUE
	FALSE
	STRING
)

var kind_nameMap = map[Kind]string{
	DEF:      "DEF",
	DO:       "DO",
	END:      "END",
	RPAR:     "RPAR",
	LPAR:     "LPAR",
	HASH:     "HASH",
	COMMA:    "COMMA",
	UPPER_ID: "UPPER_ID",
	LOWER_ID: "LOWER_ID",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	STRING:   "STRING",
}

/* STRUCT */
type Kind int

/* METHODS */
func (k *Kind) String() string {
	return kind_nameMap[*k]
}
