package internal

import (
	"fmt"

	"github.com/Goamaral/goa-lang/v1/internal/ast"
	"github.com/Goamaral/goa-lang/v1/internal/token"
)

type ParserSyntaxError struct {
	SourceCodeIndex    int
	Kind               ast.Kind
	ExpectedTokenKinds []token.Kind
	InnerErrors        []*ParserSyntaxError
}

type popContext struct {
	parser      *parser
	nPops       uint
	kind        ast.Kind
	innerErrors []*ParserSyntaxError
}

func NewPopContext(parser *parser, kind ast.Kind) *popContext {
	return &popContext{parser: parser, kind: kind}
}

func (ctx *popContext) Pop(tokenKinds ...token.Kind) (token.Token, *ParserSyntaxError) {
	tk, ok := ctx.parser.lexer.Pop(tokenKinds)
	if !ok {
		return token.Token{}, ctx.BuildError(tokenKinds)
	}

	ctx.parser.Log("Poping %s", tk.String())
	ctx.nPops++
	return tk, nil
}

func (ctx *popContext) Cleanup() {
	ctx.parser.lexer.UndoPops(ctx.nPops)
	ctx.nPops = 0
}

func (ctx *popContext) BuildNode() *ast.Node {
	n := ast.NewNode(ctx.kind, ctx.nPops)
	ctx.nPops = 0
	ctx.parser.Log("Built %s", n.String())
	return n
}

func (ctx *popContext) AddNode(n *ast.Node) {
	ctx.nPops += n.NPops
}

func (ctx *popContext) RemoveNode(n *ast.Node) {
	ctx.parser.lexer.UndoPops(n.NPops)
	ctx.nPops -= n.NPops
}

func (ctx *popContext) AddError(err *ParserSyntaxError) {
	ctx.innerErrors = append(ctx.innerErrors, err)
}

func (ctx *popContext) BuildError(expectedTokenKinds []token.Kind) *ParserSyntaxError {
	return &ParserSyntaxError{
		SourceCodeIndex:    ctx.parser.lexer.NextTokenIndex,
		Kind:               ctx.kind,
		InnerErrors:        ctx.innerErrors,
		ExpectedTokenKinds: expectedTokenKinds,
	}
}

func BuildAst(lexer *Lexer, inDebugMode bool) (*ast.Ast, *ParserSyntaxError) {
	p := parser{lexer: lexer, inDebugMode: inDebugMode}
	p.Log("===== BUILDING AST =====")
	defer p.Log("")

	pkg, err := p.BuildPackage()
	if err != nil {
		return nil, err
	}

	return &ast.Ast{Package: pkg}, nil
}

type parser struct {
	lexer       *Lexer
	inDebugMode bool
}

func (p *parser) Log(format string, a ...interface{}) {
	if p.inDebugMode {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

// Package: FuncDef
func (p *parser) BuildPackage() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Package)
	defer ctx.Cleanup()

	funcDef, err := p.BuildFuncDef()
	if funcDef != nil {
		return nil, err
	}
	ctx.AddNode(funcDef)

	return ctx.BuildNode().AddChildren(funcDef), nil
}

// FuncDef: DEF Id Block
func (p *parser) BuildFuncDef() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.FuncDef)
	defer ctx.Cleanup()

	_, err := ctx.Pop(token.DEF)
	if err != nil {
		return nil, err
	}

	id, err := p.BuildId()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(id)

	block, err := p.BuildBlock()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(block)

	return ctx.BuildNode().AddValue(id.Value).AddChildren(block), nil
}

// Id: UPPER_ID | LOWER_ID
func (p *parser) BuildId() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Id)
	defer ctx.Cleanup()

	tk, err := ctx.Pop(token.UPPER_ID, token.LOWER_ID)
	if err != nil {
		return nil, err
	}

	return ctx.BuildNode().AddValue(tk.Value), nil
}

// Block: DO StmtList END
func (p *parser) BuildBlock() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Block)
	defer ctx.Cleanup()

	_, err := ctx.Pop(token.DO)
	if err != nil {
		return nil, err
	}

	var stmtList []*ast.Node
	for {
		stmt, err := p.BuildStmt()
		if err != nil {
			return nil, err
		}
		if stmt == nil {
			break
		}
		stmtList = append(stmtList, stmt)
		ctx.AddNode(stmt)
	}

	_, err = ctx.Pop(token.END)
	if err != nil {
		return nil, err
	}

	return ctx.BuildNode().AddChildren(stmtList...), nil
}

// Stmt: FuncCall | VarDecl
func (p *parser) BuildStmt() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Stmt)
	defer ctx.Cleanup()

	funcCall, funcCallErr := p.BuildFuncCall()
	if funcCallErr == nil {
		return funcCall, nil
	}
	ctx.AddError(funcCallErr)

	varDecl, varDeclErr := p.BuildVarDecl()
	if varDeclErr == nil {
		return varDecl, nil
	}
	ctx.AddError(varDeclErr)

	return nil, ctx.BuildError(nil)
}

// FuncCall: GoaFuncCall
func (p *parser) BuildFuncCall() (*ast.Node, *ParserSyntaxError) {
	goaFuncCall, err := p.BuildGoaFuncCall()
	if err != nil {
		return nil, err
	}

	return goaFuncCall, nil
}

// GoaFuncCall: HASH UPPER_ID FuncCallArgList
func (p *parser) BuildGoaFuncCall() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.GoaFuncCall)
	defer ctx.Cleanup()

	_, err := ctx.Pop(token.HASH)
	if err != nil {
		return nil, err
	}

	upperId, err := ctx.Pop(token.UPPER_ID)
	if err != nil {
		return nil, err
	}

	funcCallArgList, err := p.BuildFuncCallArgList()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(funcCallArgList)

	return ctx.BuildNode().AddValue(upperId.Value).AddChildren(funcCallArgList), nil
}

// FuncCallArgList: LPAR (FuncCallArg | FuncCallArg COMMA FuncCallArgList | Empty) RPAR
func (p *parser) BuildFuncCallArgList() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.FuncCallArgList)
	defer ctx.Cleanup()

	_, err := ctx.Pop(token.LPAR)
	if err != nil {
		return nil, err
	}

	var funcCallArgList []*ast.Node
	for {
		funcCallArg, err := p.BuildFuncCallArg()
		if err != nil {
			break
		}
		funcCallArgList = append(funcCallArgList, funcCallArg)
		ctx.AddNode(funcCallArg)

		_, err = ctx.Pop(token.COMMA)
		if err != nil {
			break
		}
	}

	_, err = ctx.Pop(token.RPAR)
	if err != nil {
		return nil, err
	}

	return ctx.BuildNode().AddChildren(funcCallArgList...), nil
}

// FuncCallArg: Expr
func (p *parser) BuildFuncCallArg() (*ast.Node, *ParserSyntaxError) {
	expr, err := p.BuildExpr()
	if err != nil {
		return nil, err
	}

	return expr, nil
}

// Expr: BinOp | ParExpr | Terminal
func (p *parser) BuildExpr() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Expr)
	defer ctx.Cleanup()

	terminal, terminalErr := p.BuildTerminal()
	if terminalErr == nil {
		binOp, binOpErr := p.BuildBinOp(terminal)
		if binOpErr == nil {
			return binOp, nil
		}
		ctx.AddError(binOpErr)
	} else {
		ctx.AddError(terminalErr)
	}

	parExpr, parExprErr := p.BuildParExpr()
	if parExprErr == nil {
		return parExpr, nil
	}

	return nil, ctx.BuildError(nil)
}

// ParExpr: LPAR Expr RPAR
func (p *parser) BuildParExpr() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.ParExpr)
	defer ctx.Cleanup()

	_, err := ctx.Pop(token.LPAR)
	if err != nil {
		return nil, err
	}

	expr, err := p.BuildExpr()
	if err != nil {
		return nil, err
	}

	_, err = ctx.Pop(token.RPAR)
	if err != nil {
		return nil, err
	}

	return expr, nil
}

// BinOp: Expr (PLUS | MINUS) Expr
func (p *parser) BuildBinOp(lExpr *ast.Node) (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.BinOp)
	defer ctx.Cleanup()

	tk, err := ctx.Pop(token.PLUS, token.MINUS)
	if err != nil {
		return nil, err
	}

	rExpr, err := p.BuildExpr()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(rExpr)

	ctx.AddNode(lExpr)
	return ctx.BuildNode().AddChildren(lExpr, rExpr).AddToken(tk), nil
}

// VarDecl: DataType Id
func (p *parser) BuildVarDecl() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.VarDecl)
	defer ctx.Cleanup()

	datatype, err := p.BuildDataType()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(datatype)

	id, err := p.BuildId()
	if err != nil {
		return nil, err
	}
	ctx.AddNode(id)

	return ctx.BuildNode().AddValue(id.Value).AddDataType(datatype), nil
}

// DataType: BOOL_PTR | BOOL | INT_PTR | INT | STRING_STR | STRING
func (p *parser) BuildDataType() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.DataType)
	defer ctx.Cleanup()

	tk, err := ctx.Pop(token.BOOL_PTR, token.BOOL, token.INT_PTR, token.INT, token.STR_PTR, token.STR)
	if err != nil {
		return nil, err
	}

	return ctx.BuildNode().AddToken(tk), nil
}

// Terminal: UntypedConstant | Id
func (p *parser) BuildTerminal() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.Terminal)
	defer ctx.Cleanup()

	uc, err := p.BuildUntypedConstant()
	if err == nil {
		return uc, nil
	}
	ctx.AddError(err)

	id, err := p.BuildId()
	if err == nil {
		return id, nil
	}
	ctx.AddError(err)

	return nil, ctx.BuildError(nil)
}

// UntypedConstant: BOOL_LIT | INT_LIT | STR_LIT | NIL
func (p *parser) BuildUntypedConstant() (*ast.Node, *ParserSyntaxError) {
	ctx := NewPopContext(p, ast.UntypedConstant)
	tk, err := ctx.Pop(token.BOOL_LIT, token.INT_LIT, token.STR_LIT, token.NIL)
	if err != nil {
		return nil, err
	}

	switch tk.Kind {
	case token.INT_LIT:
		ctx.kind = ast.Int
	case token.STR_LIT:
		ctx.kind = ast.Str
	case token.NIL:
		ctx.kind = ast.Nil
	case token.BOOL_LIT:
		ctx.kind = ast.Bool
	default:
		panic("unreachable")
	}

	return ctx.BuildNode().AddValue(tk.Value), nil
}
