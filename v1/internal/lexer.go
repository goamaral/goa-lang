package internal

import (
	"github.com/Goamaral/goa-lang/v1/internal/token"
)

type YaccLexer interface {
	Pop(kinds ...token.Kind) (token.Token, bool)
	UndoPops(nPops int)
}

type yaccLexer struct {
	lexer           *token.Lexer
	parsedTokens    int
	maxParsedTokens int
}

func NewYaccLexer(lexer *token.Lexer) YaccLexer {
	return &yaccLexer{lexer: lexer}
}

func (lf *yaccLexer) Pop(kinds ...token.Kind) (token.Token, bool) {
	if lf.parsedTokens < len(lf.lexer.Tokens) {
		tk := lf.lexer.Tokens[lf.parsedTokens]

		if len(kinds) > 0 {
			allowed := false
			for _, kind := range kinds {
				if tk.Kind == kind {
					allowed = true
					break
				}
			}
			if !allowed {
				return token.Token{}, false
			}
		}

		if lf.parsedTokens > lf.maxParsedTokens {
			lf.maxParsedTokens = lf.parsedTokens
		}

		lf.parsedTokens++
		return tk, true
	} else {
		return token.Token{}, false
	}
}

func (lf *yaccLexer) UndoPops(nPops int) {
	lf.parsedTokens = lf.parsedTokens - nPops
}
