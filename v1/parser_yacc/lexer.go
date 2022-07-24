package parser_yacc

import (
	"fmt"
	"regexp"
	"sort"
)

/* CONSTANTS */
var regexPriorityList = []*regexp.Regexp{
	defRegex,
	doRegex,
	endRegex,
	lparRegex,
	rparRegex,
	hashRegex,
	commaRegex,
	trueRegex,
	falseRegex,
	nilRegex,
	integerRegex,
	upperIdRegex,
	lowerIdRegex,
	stringRegex,
}

/* STRUCT */
type Lexer struct {
	sourceCode   string
	Tokens       []Token
	LineNumber   int
	ColumnNumber int
}

/* FUNCTIONS */
func NewLexer(sourceCode string) *Lexer {
	return &Lexer{sourceCode: sourceCode, LineNumber: 1, ColumnNumber: 1}
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

/* PRIVATE */
// Extract tokens from chunk
func (l *Lexer) extractTokens(chunk string) {
	for chunk = l.extractToken(chunk); len(chunk) > 0; chunk = l.extractToken(chunk) {
	}
}

// Extract token from chunk
type extractTokenMatch struct {
	score int
	token Token
}

func (l *Lexer) extractToken(chunk string) string {
	if len(chunk) == 0 {
		return ""
	}

	matches := []extractTokenMatch{}
	for _, rgx := range regexPriorityList {
		loc := rgx.FindStringIndex(chunk)

		if loc != nil && loc[0] == 0 {
			matches = append(matches, extractTokenMatch{score: loc[1], token: Token{Kind: regex_KindMap[rgx], Value: chunk[loc[0]:loc[1]]}})
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
