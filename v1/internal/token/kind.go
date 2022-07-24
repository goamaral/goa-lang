package token

/* CONSTANTS */
const (
	UNKNOWN Kind = iota
	DEF
	DO
	END
	RPAR
	LPAR
	HASH
	COMMA
	TRUE
	FALSE
	INTEGER
	UPPER_ID
	LOWER_ID
	STRING
	NIL
)

var kind_nameMap = map[Kind]string{
	UNKNOWN:  "UNKNOWN",
	DEF:      "DEF",
	DO:       "DO",
	END:      "END",
	RPAR:     "RPAR",
	LPAR:     "LPAR",
	HASH:     "HASH",
	COMMA:    "COMMA",
	TRUE:     "TRUE",
	FALSE:    "FALSE",
	INTEGER:  "INTEGER",
	UPPER_ID: "UPPER_ID",
	LOWER_ID: "LOWER_ID",
	STRING:   "STRING",
	NIL:      "NIL",
}

/* STRUCT */
type Kind int

/* METHODS */
func (k *Kind) String() string {
	return kind_nameMap[*k]
}
