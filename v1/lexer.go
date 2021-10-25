package main

import "fmt"

type yyLex struct {
	code string
	pos  int
}

func (lex *yyLex) Lex(lval *yySymType) int {
	// Ignore spaces
	// TODO: Ignore tabs and new lines
	token := ""
	collectingToken := false
	for ; lex.pos < len(lex.code); lex.pos++ {
		c := lex.code[lex.pos]

		if c == ' ' || c == '\t' || c == '\n' {
			if collectingToken {
				lval.strVal = token
				return lex.identifyToken(token)
			} else {
				continue
			}
		} else {
			collectingToken = true
			token = token + string(c)
		}
	}

	if collectingToken {
		lval.strVal = token
		return lex.identifyToken(token)
	}

	return 0
}

func (yyLex) identifyToken(token string) int {
	switch token {
	case "def":
		fmt.Print("DEF ")
		return DEF
	case "do":
		fmt.Println("DO")
		return DO
	case "end":
		fmt.Println("END")
		return END
	default:
		fmt.Printf("ID(%s) ", token)
		return ID
	}
}

func (l *yyLex) Error(s string) {
	fmt.Printf("Syntax error: %s\n", s)
}
