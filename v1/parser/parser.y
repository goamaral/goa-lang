%{
package parser

import "github.com/Goamaral/goa-lang/v1/ast"
import "github.com/Goamaral/goa-lang/v1/parser/lexer"
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	node ast.Node
	nodeList []ast.Node
	value string
	valueList []string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <node> Prog Stmt
%type <node> FuncDef FuncDefBody
%type <node> FuncCall  GoaFuncCall
%type <nodeList> StmtList
%type <value> Terminal UntypedConstant Id Empty FuncCallArg
%type <valueList> FuncCallArgList

// same for terminals
%token <value> DEF DO END UPPER_ID LOWER_ID LPAR RPAR HASH COMMA TRUE FALSE

%start Prog

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
Prog: FuncDef { syntaxTree.Root.AddChild($1) };

/* Function Definition */
FuncDef: DEF Id FuncDefBody { $$ = ast.NewNode(ast.FuncDef, []string{$2}, []ast.Node{$3}) };

FuncDefBody: DO StmtList END { $$ = ast.NewNode(ast.FuncDefBody, nil, $2) };

/* Statements */
StmtList: StmtList Stmt { $$ = append($1, $2) } | Empty { $$ = nil };

Stmt: FuncCall { $$ = $1 };

/* Function Call */
FuncCall: GoaFuncCall { $$ = $1 };

GoaFuncCall: HASH UPPER_ID LPAR FuncCallArgList RPAR { $$ = ast.NewNode(ast.GoaFuncCall, []string{$2}, []ast.Node{ast.NewNode(ast.FuncCallArgs, $4, nil)}) };

FuncCallArgList: FuncCallArg { $$ = append($$, $1) } | FuncCallArgList COMMA FuncCallArg { $$ = append($1, $3) } | Empty { $$ = nil }; // TODO

FuncCallArg: Terminal { $$ = $1 };

/* Terminal */
Terminal: UntypedConstant { $$ = $1 } | Id { $$ = $1 };
UntypedConstant: TRUE { $$ = $1 } | FALSE { $$ = $1 };
Id: UPPER_ID { $$ = $1 } | LOWER_ID { $$ = $1 };
Empty: {};

%%
var syntaxTree ast.Ast

func Parse(lex *lexer.Lexer) (ast.Ast) {
	yyErrorVerbose = true

	syntaxTree = ast.New()
	yyParse(&lexerFrontend {
		lexer: *lex,
	})
	syntaxTree.Print()

	return syntaxTree
}