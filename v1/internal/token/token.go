package token

import (
	"fmt"
	"regexp"
)

/* CONSTANTS */
var (
	// Keywords
	DefRegex *regexp.Regexp = regexp.MustCompile(`def`)
	DoRegex                 = regexp.MustCompile(`do`)
	EndRegex                = regexp.MustCompile(`end`)

	// Symbols
	LparRegex  = regexp.MustCompile(`\(`)
	RparRegex  = regexp.MustCompile(`\)`)
	HashRegex  = regexp.MustCompile(`#`)
	CommaRegex = regexp.MustCompile(`,`)
	PlusRegex  = regexp.MustCompile(`\+`)
	MinusRegex = regexp.MustCompile(`-`)

	// Datatypes
	BoolPtrRegex = regexp.MustCompile(`bool\*`)
	BoolRegex    = regexp.MustCompile(`bool`)
	IntPtrRegex  = regexp.MustCompile(`int\*`)
	IntRegex     = regexp.MustCompile(`int`)
	StrPtrRegex  = regexp.MustCompile(`string\*`)
	StrRegex     = regexp.MustCompile(`string`)

	// Untyped constants
	BoolLiteralRegex = regexp.MustCompile(`true|false`)
	IntLiteralRegex  = regexp.MustCompile(`-?\d+`)
	StrLiteralRegex  = regexp.MustCompile(`\"[^\"]*\"`)
	NilRegex         = regexp.MustCompile(`nil`)

	// Id
	UpperIdRegex = regexp.MustCompile(`[A-Z]([a-zA-Z]|_|\d)*`)
	LowerIdRegex = regexp.MustCompile(`[a-z]([a-zA-Z]|_|\d)*`)
)

var RegexToKind = map[*regexp.Regexp]Kind{
	// Keywords
	DefRegex: DEF,
	DoRegex:  DO,
	EndRegex: END,

	// Symbols
	LparRegex:  LPAR,
	RparRegex:  RPAR,
	HashRegex:  HASH,
	CommaRegex: COMMA,
	PlusRegex:  PLUS,
	MinusRegex: MINUS,

	// Datatypes
	BoolPtrRegex: BOOL_PTR,
	BoolRegex:    BOOL,
	IntPtrRegex:  INT_PTR,
	IntRegex:     INT,
	StrPtrRegex:  STR_PTR,
	StrRegex:     STR,

	// Untyped constants
	BoolLiteralRegex: BOOL_LIT,
	IntLiteralRegex:  INT_LIT,
	StrLiteralRegex:  STR_LIT,
	NilRegex:         NIL,

	// Id
	UpperIdRegex: UPPER_ID,
	LowerIdRegex: LOWER_ID,
}

var kindIsDisplayable = map[Kind]bool{
	BOOL_LIT: true,
	INT_LIT:  true,
	STR_LIT:  true,
	UPPER_ID: true,
	LOWER_ID: true,
}

/* STRUCT */
type Token struct {
	Kind            Kind
	Value           []byte
	SourceCodeIndex int
}

/* METHODS */
func (t *Token) String() string {
	if kindIsDisplayable[t.Kind] {
		return fmt.Sprintf("%s(%s)", t.Kind.String(), t.Value)
	} else {
		return t.Kind.String()
	}
}
