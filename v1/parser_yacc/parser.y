%{
package parser

import "github.com/Goamaral/goa-lang/v1/ast"
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	value string
	node ast.Node
	nodeList []ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <node> Prog Stmt
%type <node> FuncDef FuncDefBody
%type <node> FuncCall  GoaFuncCall
%type <nodeList> StmtList FuncCallArgList
%type <node> Terminal UntypedConstant Id Empty FuncCallArg Boolean

// Reserved words
%token <value> Y_DEF Y_DO Y_END

// Operators
%token <value> '(' ')' '#' ','

// Terminals
%token <value> Y_UPPER_ID Y_LOWER_ID Y_TRUE Y_FALSE Y_STRING Y_INTEGER Y_NIL

%start Prog

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
Prog: FuncDef { syntaxTree.Root.AddChild($1) };

/* Function Definition */
FuncDef: Y_DEF Id FuncDefBody { $$ = ast.NewNode(ast.FuncDef, []ast.Node{$2}, []ast.Node{$3}) };

FuncDefBody: Y_DO StmtList Y_END { $$ = ast.NewNode(ast.FuncDefBody, nil, $2) };

/* Statements */
StmtList: StmtList Stmt { $$ = append($1, $2) } | Empty { $$ = nil };

Stmt: FuncCall { $$ = $1 };

/* Function Call */
FuncCall: GoaFuncCall { $$ = $1 };

GoaFuncCall: '#' Y_UPPER_ID '(' FuncCallArgList ')' { $$ = ast.NewNode(ast.GoaFuncCall, []ast.Node{ast.Node{Kind: ast.Id, Value: $2}}, []ast.Node{ast.NewNode(ast.FuncCallArgs, $4, nil)}) };

FuncCallArgList: FuncCallArg { $$ = append($$, $1) } | FuncCallArgList ',' FuncCallArg { $$ = append($1, $3) } | Empty { $$ = nil };

FuncCallArg: Terminal { $$ = $1 };

/* Terminals */
Terminal: UntypedConstant { $$ = $1 } | Id { $$ = $1 };
Id: Y_UPPER_ID { $$ = ast.Node{Kind: ast.Id, Value: $1} } | Y_LOWER_ID { $$ = ast.Node{Kind: ast.Id, Value: $1} };
UntypedConstant: Boolean { $$ = $1 }
							 | Y_STRING { $$ = ast.Node{Kind: ast.String, Value: $1} }
							 | Y_INTEGER { $$ = ast.Node{Kind: ast.Integer, Value: $1} }
							 | Y_NIL { $$ = ast.Node{Kind: ast.Nil, Value: $1} };

Boolean: Y_TRUE { $$ = ast.Node{Kind: ast.Boolean, Value: $1} } | Y_FALSE { $$ = ast.Node{Kind: ast.Boolean, Value: $1} };

Empty: {};

%%
var syntaxTree ast.Ast

func Parse(lex *Lexer, inDebugMode bool) (ast.Ast, bool) {
	yyErrorVerbose = true

	syntaxTree = ast.New()
	lexerFrontend := lexerFrontend{lexer: lex}
	ok := yyParse(&lexerFrontend) == 0

	return syntaxTree, ok
}