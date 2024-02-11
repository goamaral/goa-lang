package internal

import (
	"fmt"
	"regexp"
	"sort"

	"github.com/Goamaral/goa-lang/v1/internal/token"
)

/* CONSTANTS */
var regexPriorityOrder = []*regexp.Regexp{
	// Keywords
	token.DefRegex,
	token.DoRegex,
	token.EndRegex,

	// Symbols
	token.LparRegex,
	token.RparRegex,
	token.HashRegex,
	token.CommaRegex,

	// Datatypes
	token.BoolPtrRegex,
	token.BoolRegex,
	token.IntPtrRegex,
	token.IntRegex,
	token.StrPtrRegex,
	token.StrRegex,

	// Untyped constants
	token.TrueRegex,
	token.FalseRegex,
	token.IntLiteralRegex,
	token.StrLiteralRegex,
	token.NilRegex,

	// Id
	token.UpperIdRegex,
	token.LowerIdRegex,
}

/* STRUCT */
type Lexer struct {
	SourceCode     string
	NextTokenIndex int
	LineNumber     int
	ColumnNumber   int
	Tokens         []token.Token
}

/* FUNCTIONS */
func NewLexer(sourceCode string) Lexer {
	return Lexer{SourceCode: sourceCode, LineNumber: 1, ColumnNumber: 1}
}

/* METHODS */
/* PUBLIC */
func (l *Lexer) Print() {
	fmt.Println("===== LEX =====")
	for _, tk := range l.Tokens {
		fmt.Println(tk.String())
	}
	fmt.Println()
}

// Parse tokens from sourceCode
func (l *Lexer) Parse() {
	chunk := ""
	for i := 0; i < len(l.SourceCode); i++ {
		c := l.SourceCode[i]
		if c == '\n' {
			l.LineNumber++
			l.ColumnNumber = 1
		} else if c == '\t' {
			l.ColumnNumber += 2
		} else {
			l.ColumnNumber++
		}

		if c == ' ' || c == '\t' || c == '\n' {
			l.extractTokens(chunk)
			chunk = ""
		} else {
			chunk = chunk + string(c)
		}
	}

	l.extractTokens(chunk)
}

func (l *Lexer) Pop(kinds ...token.Kind) (token.Token, bool) {
	if l.NextTokenIndex < len(l.Tokens) {
		tk := l.Tokens[l.NextTokenIndex]

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

		l.NextTokenIndex++
		return tk, true
	} else {
		return token.Token{}, false
	}
}

func (l *Lexer) UndoPops(nPops int) {
	l.NextTokenIndex = l.NextTokenIndex - nPops
}

/* PRIVATE */
// Extract tokens from chunk
func (l *Lexer) extractTokens(chunk string) {
	for chunk = l.extractToken(chunk); len(chunk) > 0; chunk = l.extractToken(chunk) {
	}
}

// Extract token from chunk
type extractTokenMatch struct {
	score int
	token token.Token
}

func (l *Lexer) extractToken(chunk string) string {
	if len(chunk) == 0 {
		return ""
	}

	matches := []extractTokenMatch{}
	for _, rgx := range regexPriorityOrder {
		loc := rgx.FindStringIndex(chunk)

		if loc != nil && loc[0] == 0 {
			matches = append(matches, extractTokenMatch{score: loc[1], token: token.Token{Kind: token.RegexToKind[rgx], Value: chunk[loc[0]:loc[1]]}})
		}
	}

	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].score > matches[j].score
	})

	if len(matches) > 0 {
		l.Tokens = append(l.Tokens, matches[0].token)
		return chunk[matches[0].score:]
	}

	panic(fmt.Sprintf("Couldn't identify chunk '%s'", chunk))
}
