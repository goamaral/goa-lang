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
	lexer.COMMA:    COMMA,
	lexer.TRUE:     TRUE,
	lexer.FALSE:    FALSE,
	lexer.STRING:   STRING,
}

/* STRUCT */
type lexerFrontend struct {
	lexer        lexer.Lexer
	parsedTokens int
}

/* METHODS */
func (lf *lexerFrontend) Lex(lval *yySymType) int {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		token := lf.lexer.Tokens[lf.parsedTokens]
		lval.value = token.Value
		lf.parsedTokens += 1
		return tokenKind_yaccTokenMap[token.Kind]
	} else {
		return 0
	}
}

func (lf *lexerFrontend) Error(err string) {
	fmt.Printf(
		"Syntax error at line %d, column %d: unexpected token \"%s\" (%s)\n",
		lf.lexer.LineNumber,
		lf.lexer.ColumnNumber,
		lf.lexer.Tokens[lf.parsedTokens-1].Value,
		err[len("syntax error: "):],
	)
}
