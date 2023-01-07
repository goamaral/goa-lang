package token

import (
	"fmt"
	"regexp"
)

/* CONSTANTS */
var (
	// Keywords
	defRegex *regexp.Regexp = regexp.MustCompile(`def`)
	doRegex                 = regexp.MustCompile(`do`)
	endRegex                = regexp.MustCompile(`end`)

	// Symbols
	lparRegex  = regexp.MustCompile(`\(`)
	rparRegex  = regexp.MustCompile(`\)`)
	hashRegex  = regexp.MustCompile(`#`)
	commaRegex = regexp.MustCompile(`,`)

	// Datatypes
	boolPtrRegex   = regexp.MustCompile(`bool\*`)
	boolRegex      = regexp.MustCompile(`bool`)
	stringPtrRegex = regexp.MustCompile(`string\*`)
	stringRegex    = regexp.MustCompile(`string`)

	// Untyped constants
	trueRegex           = regexp.MustCompile(`true`)
	falseRegex          = regexp.MustCompile(`false`)
	stringLiteralRegex  = regexp.MustCompile(`\"[^\"]*\"`)
	integerLiteralRegex = regexp.MustCompile(`-?\d+`)
	nilRegex            = regexp.MustCompile(`nil`)

	// Id
	upperIdRegex = regexp.MustCompile(`[A-Z]([a-zA-Z]|_|\d)*`)
	lowerIdRegex = regexp.MustCompile(`[a-z]([a-zA-Z]|_|\d)*`)
)

var regex_KindMap = map[*regexp.Regexp]Kind{
	// Keywords
	defRegex: DEF,
	doRegex:  DO,
	endRegex: END,

	// Symbols
	lparRegex:  LPAR,
	rparRegex:  RPAR,
	hashRegex:  HASH,
	commaRegex: COMMA,

	// Datatypes
	boolPtrRegex:   BOOL_PTR,
	boolRegex:      BOOL,
	stringPtrRegex: STRING_PTR,
	stringRegex:    STRING,

	// Untyped constants
	trueRegex:           TRUE,
	falseRegex:          FALSE,
	integerLiteralRegex: INTEGER_LIT,
	stringLiteralRegex:  STRING_LIT,
	nilRegex:            NIL,

	// Id
	upperIdRegex: UPPER_ID,
	lowerIdRegex: LOWER_ID,
}

var kind_ShouldDisplayValueMap = map[Kind]bool{
	INTEGER_LIT: true,
	STRING_LIT:  true,
	UPPER_ID:    true,
	LOWER_ID:    true,
}

/* STRUCT */
type Token struct {
	Kind  Kind
	Value string
}

/* METHODS */
func (t *Token) String() string {
	if kind_ShouldDisplayValueMap[t.Kind] {
		return fmt.Sprintf("%s(%s)", t.Kind.String(), t.Value)
	} else {
		return t.Kind.String()
	}
}
