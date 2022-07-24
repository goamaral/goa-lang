package token

import (
	"fmt"
	"regexp"
)

/* CONSTANTS */
var (
	defRegex     *regexp.Regexp = regexp.MustCompile(`def`)
	doRegex                     = regexp.MustCompile(`do`)
	endRegex                    = regexp.MustCompile(`end`)
	lparRegex                   = regexp.MustCompile(`\(`)
	rparRegex                   = regexp.MustCompile(`\)`)
	hashRegex                   = regexp.MustCompile(`#`)
	commaRegex                  = regexp.MustCompile(`,`)
	trueRegex                   = regexp.MustCompile(`true`)
	falseRegex                  = regexp.MustCompile(`false`)
	nilRegex                    = regexp.MustCompile(`nil`)
	integerRegex                = regexp.MustCompile(`-?\d+`)
	upperIdRegex                = regexp.MustCompile(`[A-Z]([a-zA-Z]|_|\d)*`)
	lowerIdRegex                = regexp.MustCompile(`[a-z]([a-zA-Z]|_|\d)*`)
	stringRegex                 = regexp.MustCompile(`\"[^\"]*\"`)
)

var regex_KindMap = map[*regexp.Regexp]Kind{
	defRegex:     DEF,
	doRegex:      DO,
	endRegex:     END,
	lparRegex:    LPAR,
	rparRegex:    RPAR,
	hashRegex:    HASH,
	commaRegex:   COMMA,
	trueRegex:    TRUE,
	falseRegex:   FALSE,
	nilRegex:     NIL,
	integerRegex: INTEGER,
	upperIdRegex: UPPER_ID,
	lowerIdRegex: LOWER_ID,
	stringRegex:  STRING,
}

var kind_ShouldDisplayValueMap = map[Kind]bool{
	INTEGER:  true,
	UPPER_ID: true,
	LOWER_ID: true,
	STRING:   true,
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
