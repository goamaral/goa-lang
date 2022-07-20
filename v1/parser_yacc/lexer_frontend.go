package parser_yacc

import (
	"fmt"
	"strings"
)

/* CONSTANTS */
var tokenKind_yaccTokenMap = map[Kind]int{
	DEF:      Y_DEF,
	DO:       Y_DO,
	END:      Y_END,
	RPAR:     ')',
	LPAR:     '(',
	HASH:     '#',
	COMMA:    ',',
	TRUE:     Y_TRUE,
	FALSE:    Y_FALSE,
	INTEGER:  Y_INTEGER,
	UPPER_ID: Y_UPPER_ID,
	LOWER_ID: Y_LOWER_ID,
	STRING:   Y_STRING,
}

/* STRUCT */
type lexerFrontend struct {
	lexer        *Lexer
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
		"Syntax error at line %d, column %d: %s\n",
		lf.lexer.LineNumber,
		lf.lexer.ColumnNumber,
		strings.ToLower(err[len("syntax error: "):]),
	)
}
