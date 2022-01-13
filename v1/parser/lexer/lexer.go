package lexer

import (
	"fmt"
	"regexp"
)

/* CONSTANTS */
var regexList = []*regexp.Regexp{
	defRegex,
	doRegex,
	endRegex,
	lparRegex,
	rparRegex,
	hashRegex,
	commaRegex,
	upperIdRegex,
	lowerIdRegex,
}

/* STRUCT */
type Lexer struct {
	sourceCode string
	Tokens     []token
}

/* FUNCTIONS */
func New(sourceCode string) Lexer {
	return Lexer{sourceCode: sourceCode}
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
	for i := 0; i < len(l.sourceCode); i++ {
		c := l.sourceCode[i]

		if c == ' ' || c == '\t' || c == '\n' {
			l.extractTokens(chunk)
			chunk = ""
		} else {
			chunk = chunk + string(c)
		}
	}

	l.extractTokens(chunk)
}

/* PRIVATE */
// Extract tokens from chunk
func (l *Lexer) extractTokens(chunk string) {
	for chunk = l.extractToken(chunk); len(chunk) > 0; chunk = l.extractToken(chunk) {
	}
}

// Extract token from chunk
func (l *Lexer) extractToken(chunk string) string {
	if len(chunk) == 0 {
		return ""
	}

	for _, rgx := range regexList {
		loc := rgx.FindStringIndex(chunk)

		if loc != nil && loc[0] == 0 {
			l.Tokens = append(l.Tokens, token{Kind: regex_KindMap[rgx], Value: chunk[loc[0]:loc[1]]})
			return chunk[loc[1]:]
		}
	}

	panic(fmt.Sprintf("Couldn't identify chunk '%s'", chunk))
}
