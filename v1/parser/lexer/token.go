package lexer

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

/* STRUCT */
type token struct {
	Kind  Kind
	Value string
}

/* METHODS */
func (t *token) String() string {
	switch t.Kind {
	case UPPER_ID:
		return fmt.Sprintf("%s(%s)", t.Kind.String(), t.Value)
	case LOWER_ID:
		return fmt.Sprintf("%s(%s)", t.Kind.String(), t.Value)
	default:
		return t.Kind.String()
	}
}
