package internal

import (
	"fmt"
	"runtime"

	"github.com/Goamaral/goa-lang/v1/internal/ast"
	"github.com/Goamaral/goa-lang/v1/internal/token"
)

type parser struct {
	lexer       YaccLexer
	inDebugMode bool
}

type parserContext struct {
	nPops    int
	undoPops bool
}

func (p *parser) Log(format string, a ...interface{}) {
	if p.inDebugMode {
		fmt.Printf(format, a...)
		fmt.Println()
	}
}

func (p *parser) Pop(ctx parserContext, kinds ...token.Kind) (token.Token, bool) {
	tk, ok := p.lexer.Pop(kinds...)
	if !ok {
		return token.Token{}, false
	}

	p.Log("Poping %s", tk.String())
	ctx.nPops++
	return tk, true
}

func (p *parser) UndoPops(ctx parserContext) *ast.Node {
	_, _, no, _ := runtime.Caller(1)
	p.Log("Line %d called undo pops", no)
	ctx.undoPops = true
	return nil
}

func (p *parser) Cleanup(ctx parserContext) {
	if ctx.undoPops {
		p.Log("Undoing %d pops", ctx.nPops)
		p.lexer.UndoPops(ctx.nPops)
	}
	ctx.undoPops = false
	ctx.nPops = 0
}

func BuildAst(lexer YaccLexer, inDebugMode bool) (*ast.Ast, bool) {
	p := parser{lexer: lexer, inDebugMode: inDebugMode}
	p.Log("===== BUILDING AST =====")
	defer p.Log("")

	pkg := p.BuildPackage()
	if pkg == nil {
		return nil, false
	}

	return &ast.Ast{Package: pkg}, true
}

// Package: FuncDef
func (p *parser) BuildPackage() *ast.Node {
	// FuncDef
	funcDef := p.BuildFuncDef()
	if funcDef == nil {
		return nil
	}

	p.Log("Built Package")
	return ast.NewNode(ast.Package).AddChildren(funcDef)
}

// FuncDef: DEF Id Block
func (p *parser) BuildFuncDef() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	var id, block *ast.Node
	var ok bool

	// DEF
	_, ok = p.Pop(ctx, token.DEF)
	if !ok {
		return p.UndoPops(ctx)
	}

	// Id
	id = p.BuildId()
	if id == nil {
		return p.UndoPops(ctx)
	}

	// Block
	block = p.BuildBlock()
	if block == nil {
		return p.UndoPops(ctx)
	}

	p.Log("Built FuncDef[%s]", id.Value)
	return ast.NewNode(ast.FuncDef).AddValue(id.Value).AddChildren(block)
}

// Id: UPPER_ID | LOWER_ID
func (p *parser) BuildId() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	var tk token.Token
	var ok bool

	// UPPER_ID | LOWER_ID
	tk, ok = p.Pop(ctx, token.UPPER_ID, token.LOWER_ID)
	if !ok {
		return p.UndoPops(ctx)
	}

	p.Log("Built Id(%s)", tk.Value)
	return ast.NewNode(ast.Id).AddValue(tk.Value)
}

// Block: DO StmtList END
func (p *parser) BuildBlock() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	// DO
	_, ok := p.Pop(ctx, token.DO)
	if !ok {
		return p.UndoPops(ctx)
	}

	// StmtList: Stmt StmtList | Empty
	var stmtList []*ast.Node
	for {
		stmt := p.BuildStmt()
		if stmt == nil {
			break
		}
		stmtList = append(stmtList, stmt)
	}

	// END
	_, ok = p.Pop(ctx, token.END)
	if !ok {
		return p.UndoPops(ctx)
	}

	p.Log("Built Block")
	return ast.NewNode(ast.Block).AddChildren(stmtList...)
}

// Stmt: FuncCall | VarDecl
func (p *parser) BuildStmt() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	funcCall := p.BuildFuncCall()
	if funcCall != nil {
		return funcCall
	}

	varDecl := p.BuildVarDecl()
	if varDecl != nil {
		return varDecl
	}

	return nil
}

// FuncCall: GoaFuncCall
func (p *parser) BuildFuncCall() *ast.Node {
	// GoaFuncCall
	goaFuncCall := p.BuildGoaFuncCall()
	if goaFuncCall != nil {
		return goaFuncCall
	}

	return nil
}

// GoaFuncCall: HASH UPPER_ID LPAR FuncCallArgList RPAR
func (p *parser) BuildGoaFuncCall() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	var upperId token.Token

	// HASH
	_, ok := p.Pop(ctx, token.HASH)
	if !ok {
		return p.UndoPops(ctx)
	}

	// UPPER_ID
	upperId, ok = p.Pop(ctx, token.UPPER_ID)
	if !ok {
		return p.UndoPops(ctx)
	}

	// LPAR
	_, ok = p.Pop(ctx, token.LPAR)
	if !ok {
		return p.UndoPops(ctx)
	}

	// FuncCallArgList
	funcCallArgList := p.BuildFuncCallArgList()
	if funcCallArgList == nil {
		return p.UndoPops(ctx)
	}

	// RPAR
	_, ok = p.Pop(ctx, token.RPAR)
	if !ok {
		return p.UndoPops(ctx)
	}

	p.Log("Built GoaFuncCall(%s)", upperId.Value)
	return ast.NewNode(ast.GoaFuncCall).AddValue(upperId.Value).AddChildren(funcCallArgList)
}

// FuncCallArgList: FuncCallArg | FuncCallArg COMMA FuncCallArgList | Empty
func (p *parser) BuildFuncCallArgList() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	var funcCallArgList []*ast.Node
	for {
		funcCallArg := p.BuildFuncCallArg()
		if funcCallArg == nil {
			break
		}
		funcCallArgList = append(funcCallArgList, funcCallArg)

		// COMMA
		_, ok := p.Pop(ctx, token.COMMA)
		if !ok {
			break
		}
	}

	p.Log("Built FuncCallArgList")
	return ast.NewNode(ast.FuncCallArgs).AddChildren(funcCallArgList...)
}

// FuncCallArg: Terminal
func (p *parser) BuildFuncCallArg() *ast.Node {
	// Terminal
	terminal := p.BuildTerminal()
	if terminal != nil {
		return terminal
	}

	return nil
}

// VarDecl: DataType Id
func (p *parser) BuildVarDecl() *ast.Node {
	datatypeTkKind := p.BuildDataType()
	if datatypeTkKind == token.UNKNOWN {
		return nil
	}

	id := p.BuildId()
	if id == nil {
		return nil
	}

	p.Log("Built VarDecl(%s) - %s", id.Value, datatypeTkKind.String())
	return ast.NewNode(ast.VarDecl).AddValue(id.Value).AddDataType(datatypeTkKind)
}

// DataType: BOOL_PTR | BOOL
func (p *parser) BuildDataType() token.Kind {
	var ctx parserContext
	defer p.Cleanup(ctx)

	datatypeTk, ok := p.Pop(ctx, token.BOOL_PTR, token.BOOL)
	if !ok {
		p.UndoPops(ctx)
		return token.UNKNOWN
	}

	return datatypeTk.Kind
}

// Terminal: UntypedConstant | Id
func (p *parser) BuildTerminal() *ast.Node {
	var terminal *ast.Node

	// UntypedConstant
	terminal = p.BuildUntypedConstant()
	if terminal != nil {
		return terminal
	}

	// Id
	terminal = p.BuildId()
	if terminal != nil {
		return terminal
	}

	return nil
}

// UntypedConstant: Boolean | STRING_LIT | INTEGER_LIT | NIL
func (p *parser) BuildUntypedConstant() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	// Boolean
	boolean := p.BuildBoolean()
	if boolean != nil {
		return boolean
	}

	// STRING_LIT | INTEGER_LIT | NIL
	tk, ok := p.Pop(ctx, token.STRING_LIT, token.INTEGER_LIT, token.NIL)
	if !ok {
		return p.UndoPops(ctx)
	}

	var kind ast.Kind
	switch tk.Kind {
	case token.STRING_LIT:
		kind = ast.String
	case token.INTEGER_LIT:
		kind = ast.Integer
	case token.NIL:
		kind = ast.Nil
	default:
		panic("unreachable")
	}

	p.Log("Built %s(%s)", tk.Kind.String(), tk.Value)
	return ast.NewNode(kind).AddValue(tk.Value)
}

// Boolean: TRUE | FALSE
func (p *parser) BuildBoolean() *ast.Node {
	var ctx parserContext
	defer p.Cleanup(ctx)

	// TRUE | FALSE
	tk, ok := p.Pop(ctx, token.TRUE, token.FALSE)
	if !ok {
		return p.UndoPops(ctx)
	}

	p.Log("Built Boolean(%s)", tk.Value)
	return ast.NewNode(ast.Boolean).AddValue(tk.Value)
}
