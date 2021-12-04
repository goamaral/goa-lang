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
	upperIdRegex                = regexp.MustCompile(`[A-Z]([a-zA-Z]|_|\d)*`)
	lowerIdRegex                = regexp.MustCompile(`[a-z]([a-zA-Z]|_|\d)*`)
)

var regex_KindMap = map[*regexp.Regexp]Kind{
	defRegex:     DEF,
	doRegex:      DO,
	endRegex:     END,
	lparRegex:    LPAR,
	rparRegex:    RPAR,
	hashRegex:    HASH,
	upperIdRegex: UPPER_ID,
	lowerIdRegex: LOWER_ID,
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
