package lexer

/* CONSTANTS */
const (
	DEF Kind = iota
	DO
	END
	RPAR
	LPAR
	HASH
	UPPER_ID
	LOWER_ID
)

var kind_nameMap = map[Kind]string{
	DEF:      "DEF",
	DO:       "DO",
	END:      "END",
	RPAR:     "RPAR",
	LPAR:     "LPAR",
	HASH:     "HASH",
	UPPER_ID: "UPPER_ID",
	LOWER_ID: "LOWER_ID",
}

/* STRUCT */
type Kind int

/* METHODS */
func (k *Kind) String() string {
	return kind_nameMap[*k]
}
