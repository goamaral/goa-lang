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
	token.PlusRegex,
	token.MinusRegex,

	// Datatypes
	token.BoolPtrRegex,
	token.BoolRegex,
	token.IntPtrRegex,
	token.IntRegex,
	token.StrPtrRegex,
	token.StrRegex,

	// Untyped constants
	token.BoolLiteralRegex,
	token.IntLiteralRegex,
	token.StrLiteralRegex,
	token.NilRegex,

	// Id
	token.UpperIdRegex,
	token.LowerIdRegex,
}

/* STRUCT */
type Lexer struct {
	SourceCode     []byte
	NextTokenIndex int
	Lines          []int
	Tokens         []token.Token
}

type Chunk struct {
	Value           []byte
	SourceCodeIndex int
}

/* FUNCTIONS */
func NewLexer(sourceCode []byte) Lexer {
	return Lexer{SourceCode: sourceCode, Lines: []int{0}}
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
	startIndex := 0
	for i := 0; i < len(l.SourceCode); i++ {
		c := l.SourceCode[i]
		if c == '\n' {
			l.Lines = append(l.Lines, i+1)
		}

		if c == ' ' || c == '\t' || c == '\n' {
			l.extractTokens(Chunk{Value: l.SourceCode[startIndex:i], SourceCodeIndex: startIndex})
			startIndex = i + 1
		}
	}

	l.extractTokens(Chunk{Value: l.SourceCode[startIndex:len(l.SourceCode)], SourceCodeIndex: startIndex})
}

func (l *Lexer) Pop(kinds []token.Kind) (token.Token, bool) {
	tk, ok := l.Peek(kinds)
	if !ok {
		return token.Token{}, false
	}
	l.NextTokenIndex++
	return tk, ok
}

func (l *Lexer) Peek(kinds []token.Kind) (token.Token, bool) {
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

		return tk, true
	} else {
		return token.Token{}, false
	}
}

func (l *Lexer) UndoPops(nPops uint) {
	l.NextTokenIndex = l.NextTokenIndex - int(nPops)
}

/* PRIVATE */
// Extract tokens from chunk
func (l *Lexer) extractTokens(chunk Chunk) {
	for len(chunk.Value) > 0 {
		chunk = l.extractToken(chunk)
	}
}

// Extract token from chunk
func (l *Lexer) extractToken(chunk Chunk) Chunk {
	matches := []token.Token{}
	for _, rgx := range regexPriorityOrder {
		loc := rgx.FindStringIndex(string(chunk.Value))

		if loc != nil && loc[0] == 0 {
			matches = append(matches, token.Token{
				Kind:            token.RegexToKind[rgx],
				Value:           chunk.Value[loc[0]:loc[1]],
				SourceCodeIndex: chunk.SourceCodeIndex,
			})
		}
	}

	sort.SliceStable(matches, func(i, j int) bool {
		return len(matches[i].Value) > len(matches[j].Value)
	})

	if len(matches) > 0 {
		l.Tokens = append(l.Tokens, matches[0])
		chunk.Value = chunk.Value[len(matches[0].Value):]
		chunk.SourceCodeIndex += len(matches[0].Value)
		return chunk
	}

	// TODO: Return LexerSyntaxError
	panic(fmt.Sprintf("Couldn't identify chunk '%s'", chunk.Value))
}
