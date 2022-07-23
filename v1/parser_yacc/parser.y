%{
package parser_yacc

import "github.com/Goamaral/goa-lang/v1/ast"
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	value string
	node *ast.Node
	nodeList []*ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <node> Package Stmt
%type <node> FuncDef FuncDefBody
%type <node> FuncCall  GoaFuncCall
%type <nodeList> StmtList FuncCallArgList
%type <node> Terminal UntypedConstant Id Empty FuncCallArg Boolean

// Reserved words
%token <value> Y_DEF Y_DO Y_END

// Operators
%token <value> Y_LPAR Y_RPAR Y_HASH Y_COMMA

// Terminals
%token <value> Y_UPPER_ID Y_LOWER_ID Y_TRUE Y_FALSE Y_STRING Y_INTEGER Y_NIL

%start Package

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
Package: FuncDef { syntaxTree.Package.AddChild($1) };

/* Function Definition */
FuncDef: Y_DEF Id FuncDefBody { $$ = ast.NewComplexNode(ast.FuncDef, []*ast.Node{$2}, []*ast.Node{$3}) };

FuncDefBody: Y_DO StmtList Y_END { $$ = ast.NewComplexNode(ast.FuncDefBody, nil, $2) };

/* Statements */
StmtList: StmtList Stmt { $$ = append($1, $2) } | Empty { $$ = nil };

Stmt: FuncCall { $$ = $1 };

/* Function Call */
FuncCall: GoaFuncCall { $$ = $1 };

GoaFuncCall: Y_HASH Y_UPPER_ID Y_LPAR FuncCallArgList Y_RPAR { $$ = ast.NewComplexNode(ast.GoaFuncCall, []*ast.Node{ast.NewSimpleNode(ast.Id, $2)}, []*ast.Node{ast.NewComplexNode(ast.FuncCallArgs, $4, nil)}) };

FuncCallArgList: FuncCallArg { $$ = append($$, $1) } | FuncCallArgList Y_COMMA FuncCallArg { $$ = append($1, $3) } | Empty { $$ = nil };

FuncCallArg: Terminal { $$ = $1 };

/* Terminals */
Terminal: UntypedConstant { $$ = $1 } | Id { $$ = $1 };
Id: Y_UPPER_ID { $$ = ast.NewSimpleNode(ast.Id, $1) } | Y_LOWER_ID { $$ = ast.NewSimpleNode(ast.Id, $1) };
UntypedConstant: Boolean { $$ = $1 }
							 | Y_STRING { $$ = ast.NewSimpleNode(ast.String, $1) }
							 | Y_INTEGER { $$ = ast.NewSimpleNode(ast.Integer, $1) }
							 | Y_NIL { $$ = ast.NewSimpleNode(ast.Nil, $1) };

Boolean: Y_TRUE { $$ = ast.NewSimpleNode(ast.Boolean, $1) } | Y_FALSE { $$ = ast.NewSimpleNode(ast.Boolean, $1) };

Empty: {};

%%
type YaccDollar yySymType

var syntaxTree *ast.Ast

func BuildAst(lexer YaccLexer, inDebugMode bool) (*ast.Ast, bool) {
	yyErrorVerbose = true

	syntaxTree = ast.New()
	ok := yyParse(lexer) == 0

	return syntaxTree, ok
}