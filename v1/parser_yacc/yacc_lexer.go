package parser_yacc

import (
	"fmt"
	"strings"
)

type YaccLexer interface {
	Pop(kinds ...Kind) (Token, bool)
	UndoPops(nPops int)
	PeekTokenString() string
}

type yaccLexer struct {
	lexer           *Lexer
	parsedTokens    int
	maxParsedTokens int
}

func NewYaccLexer(lexer *Lexer) YaccLexer {
	return &yaccLexer{lexer: lexer}
}

func (lf *yaccLexer) Error(err string) {
	fmt.Printf(
		"Syntax error at line %d, column %d: %s\n",
		lf.lexer.LineNumber,
		lf.lexer.ColumnNumber,
		strings.ToLower(err[len("syntax error: "):]),
	)
}

func (lf *yaccLexer) Pop(kinds ...Kind) (Token, bool) {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		token := lf.lexer.Tokens[lf.parsedTokens]

		if len(kinds) > 0 {
			allowed := false
			for _, kind := range kinds {
				if token.Kind == kind {
					allowed = true
					break
				}
			}
			if !allowed {
				return Token{}, false
			}
		}

		if lf.parsedTokens > lf.maxParsedTokens {
			lf.maxParsedTokens = lf.parsedTokens
		}

		lf.parsedTokens++
		return token, true
	} else {
		return Token{}, false
	}
}

func (lf *yaccLexer) UndoPops(nPops int) {
	lf.parsedTokens = lf.parsedTokens - nPops
}

func (lf *yaccLexer) PeekTokenString() string {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		return fmt.Sprintf("%s@%d", lf.lexer.Tokens[lf.parsedTokens].String(), lf.parsedTokens)
	} else {
		tk := Token{}
		return tk.String()
	}
}
