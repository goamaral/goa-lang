package parser

import "fmt"

type lexer struct {
	code   string
	pos    int
	tokens []*token
}

type token struct {
	kind  int
	value string
}

var tokenKind_nameMap = map[int]string{
	DEF: "DEF",
	DO:  "DO",
	END: "END",
	ID:  "ID",
}

func (tok token) String() string {
	switch tok.kind {
	case ID:
		return fmt.Sprintf("%s(%s)", tokenKind_nameMap[tok.kind], tok.value)
	default:
		return tokenKind_nameMap[tok.kind]
	}
}

func (lex *lexer) Lex(lval *yySymType) int {
	// Ignore spaces
	// TODO: Ignore tabs and new lines
	value := ""
	collectingToken := false
	for ; lex.pos < len(lex.code); lex.pos++ {
		c := lex.code[lex.pos]

		if c == ' ' || c == '\t' || c == '\n' {
			if collectingToken {
				lval.value = value
				kind := lex.identifyToken(value)
				lex.tokens = append(lex.tokens, &token{kind, value})
				return kind
			} else {
				continue
			}
		} else {
			collectingToken = true
			value = value + string(c)
		}
	}

	if collectingToken {
		lval.value = value
		kind := lex.identifyToken(value)
		lex.tokens = append(lex.tokens, &token{kind, value})
		return kind
	}

	return 0
}

func (lexer) identifyToken(token string) int {
	switch token {
	case "def":
		return DEF
	case "do":
		return DO
	case "end":
		return END
	default:
		return ID
	}
}

func (lex *lexer) Error(s string) {
	fmt.Printf("Syntax error: %s\n", s)
}

func (lex lexer) Print() {
	fmt.Println("===== LEX =====")
	for _, tk := range lex.tokens {
		fmt.Println(tk.String())
	}
	fmt.Println()
}
