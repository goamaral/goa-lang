package parser_yacc

import (
	"fmt"
	"strings"
)

var tokenKind_yaccTokenMap = map[Kind]int{
	DEF:      Y_DEF,
	DO:       Y_DO,
	END:      Y_END,
	RPAR:     Y_RPAR,
	LPAR:     Y_LPAR,
	HASH:     Y_HASH,
	COMMA:    Y_COMMA,
	TRUE:     Y_TRUE,
	FALSE:    Y_FALSE,
	INTEGER:  Y_INTEGER,
	UPPER_ID: Y_UPPER_ID,
	LOWER_ID: Y_LOWER_ID,
	STRING:   Y_STRING,
}

type YaccLexer interface {
	yyLexer
	GetToken() (token, bool)
}

type yaccLexer struct {
	lexer        *Lexer
	parsedTokens int
}

func NewYaccLexer(lexer *Lexer) YaccLexer {
	return &yaccLexer{lexer: lexer}
}

func (lf *yaccLexer) Lex(lval *yySymType) int {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		token := lf.lexer.Tokens[lf.parsedTokens]
		lval.value = token.Value
		lf.parsedTokens += 1
		return tokenKind_yaccTokenMap[token.Kind]
	} else {
		return 0
	}
}

func (lf *yaccLexer) Error(err string) {
	fmt.Printf(
		"Syntax error at line %d, column %d: %s\n",
		lf.lexer.LineNumber,
		lf.lexer.ColumnNumber,
		strings.ToLower(err[len("syntax error: "):]),
	)
}

func (lf *yaccLexer) GetToken() (token, bool) {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		token := lf.lexer.Tokens[lf.parsedTokens]
		lf.parsedTokens += 1
		return token, true
	} else {
		return token{}, false
	}
}
