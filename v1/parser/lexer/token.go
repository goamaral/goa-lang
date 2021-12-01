package lexer

import (
	"fmt"
	"regexp"
)

/* CONSTANTS */
var (
	defRegex  *regexp.Regexp = regexp.MustCompile(`def`)
	doRegex                  = regexp.MustCompile(`do`)
	endRegex                 = regexp.MustCompile(`end`)
	lparRegex                = regexp.MustCompile(`\(`)
	rparRegex                = regexp.MustCompile(`\)`)
	hashRegex                = regexp.MustCompile(`#`)
	idRegex                  = regexp.MustCompile(`[a-zA-Z]([a-zA-Z]|_|\d)*`)
)

var regex_KindMap = map[*regexp.Regexp]Kind{
	defRegex:  DEF,
	doRegex:   DO,
	endRegex:  END,
	lparRegex: LPAR,
	rparRegex: RPAR,
	hashRegex: HASH,
	idRegex:   ID,
}

/* STRUCT */
type token struct {
	Kind  Kind
	Value string
}

/* METHODS */
func (t *token) String() string {
	switch t.Kind {
	case ID:
		return fmt.Sprintf("%s(%s)", kind_nameMap[t.Kind], t.Value)
	default:
		return kind_nameMap[t.Kind]
	}
}
