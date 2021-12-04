package parser

import (
	"fmt"

	"github.com/Goamaral/goa-lang/v1/parser/lexer"
)

/* CONSTANTS */
var tokenKind_yaccTokenMap = map[lexer.Kind]int{
	lexer.DEF:      DEF,
	lexer.DO:       DO,
	lexer.END:      END,
	lexer.UPPER_ID: UPPER_ID,
	lexer.LOWER_ID: LOWER_ID,
	lexer.RPAR:     RPAR,
	lexer.LPAR:     LPAR,
	lexer.HASH:     HASH,
}

/* STRUCT */
type lexerFrontend struct {
	lexer      lexer.Lexer
	tokenIndex int
}

/* METHODS */
func (lf *lexerFrontend) Lex(lval *yySymType) int {
	if lf.tokenIndex < len(lf.lexer.Tokens) {
		token := lf.lexer.Tokens[lf.tokenIndex]
		lval.value = token.Value
		lf.tokenIndex += 1
		return tokenKind_yaccTokenMap[token.Kind]
	} else {
		return 0
	}
}

func (lf *lexerFrontend) Error(s string) {
	fmt.Printf("Syntax error at line X, column X\n")
	/* fmt.Printf("Syntax error at line %d, column %d\n", lf.lineNumber, lf.columnNumber) */
}
