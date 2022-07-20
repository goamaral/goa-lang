// Code generated by goyacc -o parser/parser.go -v parser.output parser/parser.y. DO NOT EDIT.

//line parser/parser.y:2
package parser_yacc

import __yyfmt__ "fmt"

//line parser/parser.y:2

import "github.com/Goamaral/goa-lang/v1/ast"

//line parser/parser.y:9
type yySymType struct {
	yys      int
	value    string
	node     ast.Node
	nodeList []ast.Node
}

const Y_DEF = 57346
const Y_DO = 57347
const Y_END = 57348
const Y_UPPER_ID = 57349
const Y_LOWER_ID = 57350
const Y_TRUE = 57351
const Y_FALSE = 57352
const Y_STRING = 57353
const Y_INTEGER = 57354
const Y_NIL = 57355
const UMINUS = 57356

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"Y_DEF",
	"Y_DO",
	"Y_END",
	"'('",
	"')'",
	"'#'",
	"','",
	"Y_UPPER_ID",
	"Y_LOWER_ID",
	"Y_TRUE",
	"Y_FALSE",
	"Y_STRING",
	"Y_INTEGER",
	"Y_NIL",
	"'|'",
	"'&'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UMINUS",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line parser/parser.y:74

var syntaxTree ast.Ast

func Parse(lex *Lexer, inDebugMode bool) (ast.Ast, bool) {
	yyErrorVerbose = true

	syntaxTree = ast.New()
	lexerFrontend := lexerFrontend{lexer: lex}
	ok := yyParse(&lexerFrontend) == 0

	return syntaxTree, ok
}

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 34

var yyAct = [...]int8{
	19, 5, 6, 28, 29, 25, 26, 27, 10, 5,
	6, 30, 16, 31, 11, 8, 17, 15, 3, 23,
	24, 22, 21, 4, 18, 9, 20, 14, 13, 7,
	2, 12, 32, 1,
}

var yyPact = [...]int16{
	14, -1000, -1000, -2, 10, -1000, -1000, -1000, -1000, 8,
	-1000, -1000, -1000, -1000, -1000, 1, 9, -10, 3, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -10, -1000,
}

var yyPgo = [...]int8{
	0, 33, 31, 30, 29, 28, 27, 25, 24, 22,
	21, 19, 8, 0, 20,
}

var yyR1 = [...]int8{
	0, 1, 3, 4, 7, 7, 2, 5, 6, 8,
	8, 8, 13, 9, 9, 11, 11, 10, 10, 10,
	10, 14, 14, 12,
}

var yyR2 = [...]int8{
	0, 1, 3, 3, 2, 1, 1, 1, 5, 1,
	3, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 0,
}

var yyChk = [...]int16{
	-1000, -1, -3, 4, -11, 11, 12, -4, 5, -7,
	-12, 6, -2, -5, -6, 9, 11, 7, -8, -13,
	-12, -9, -10, -11, -14, 15, 16, 17, 13, 14,
	8, 10, -13,
}

var yyDef = [...]int8{
	0, -2, 1, 0, 0, 15, 16, 2, 23, 0,
	5, 3, 4, 6, 7, 0, 0, 23, 0, 9,
	11, 12, 13, 14, 17, 18, 19, 20, 21, 22,
	8, 0, 10,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 9, 3, 24, 19, 3,
	7, 8, 22, 20, 10, 21, 3, 23, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 18,
}

var yyTok2 = [...]int8{
	2, 3, 4, 5, 6, 11, 12, 13, 14, 15,
	16, 17, 25,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:41
		{
			syntaxTree.Root.AddChild(yyDollar[1].node)
		}
	case 2:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.y:44
		{
			yyVAL.node = ast.NewNode(ast.FuncDef, []ast.Node{yyDollar[2].node}, []ast.Node{yyDollar[3].node})
		}
	case 3:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.y:46
		{
			yyVAL.node = ast.NewNode(ast.FuncDefBody, nil, yyDollar[2].nodeList)
		}
	case 4:
		yyDollar = yyS[yypt-2 : yypt+1]
//line parser/parser.y:49
		{
			yyVAL.nodeList = append(yyDollar[1].nodeList, yyDollar[2].node)
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:49
		{
			yyVAL.nodeList = nil
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:51
		{
			yyVAL.node = yyDollar[1].node
		}
	case 7:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:54
		{
			yyVAL.node = yyDollar[1].node
		}
	case 8:
		yyDollar = yyS[yypt-5 : yypt+1]
//line parser/parser.y:56
		{
			yyVAL.node = ast.NewNode(ast.GoaFuncCall, []ast.Node{ast.Node{Kind: ast.Id, Value: yyDollar[2].value}}, []ast.Node{ast.NewNode(ast.FuncCallArgs, yyDollar[4].nodeList, nil)})
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:58
		{
			yyVAL.nodeList = append(yyVAL.nodeList, yyDollar[1].node)
		}
	case 10:
		yyDollar = yyS[yypt-3 : yypt+1]
//line parser/parser.y:58
		{
			yyVAL.nodeList = append(yyDollar[1].nodeList, yyDollar[3].node)
		}
	case 11:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:58
		{
			yyVAL.nodeList = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:60
		{
			yyVAL.node = yyDollar[1].node
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:63
		{
			yyVAL.node = yyDollar[1].node
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:63
		{
			yyVAL.node = yyDollar[1].node
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:64
		{
			yyVAL.node = ast.Node{Kind: ast.Id, Value: yyDollar[1].value}
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:64
		{
			yyVAL.node = ast.Node{Kind: ast.Id, Value: yyDollar[1].value}
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:65
		{
			yyVAL.node = yyDollar[1].node
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:66
		{
			yyVAL.node = ast.Node{Kind: ast.String, Value: yyDollar[1].value}
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:67
		{
			yyVAL.node = ast.Node{Kind: ast.Integer, Value: yyDollar[1].value}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:68
		{
			yyVAL.node = ast.Node{Kind: ast.Nil, Value: yyDollar[1].value}
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:70
		{
			yyVAL.node = ast.Node{Kind: ast.Boolean, Value: yyDollar[1].value}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line parser/parser.y:70
		{
			yyVAL.node = ast.Node{Kind: ast.Boolean, Value: yyDollar[1].value}
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line parser/parser.y:72
		{
		}
	}
	goto yystack /* stack new state and value */
}