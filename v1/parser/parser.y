%{
package parser

import "github.com/Goamaral/goa-lang/v1/ast"
import "github.com/Goamaral/goa-lang/v1/parser/lexer"
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
%type <node> FuncCall GoaFuncCall
%type <value> Id Empty
%type <nodeList> StmtList

// same for terminals
%token <value> DEF DO END UPPER_ID LOWER_ID LPAR RPAR HASH

%start Prog

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
Prog: FuncDef { syntaxTree.Root.AddChild($1) };

/* Function Definition */
FuncDef: DEF Id FuncDefBody { $$ = ast.NewNode(ast.FuncDef, $2, []ast.Node{$3}) };

FuncDefBody: DO StmtList END { $$ = ast.NewNode(ast.FuncDefBody, "", $2) };

/* Statements */
StmtList: Stmt StmtList { $$ = append($$, $1) } | Empty { $$ = nil };

Stmt: FuncCall { $$ = $1 };

/* Function Call */
FuncCall: GoaFuncCall { $$ = $1 };

GoaFuncCall: HASH UPPER_ID LPAR RPAR { $$ = ast.NewNode(ast.GoaFuncCall, $2, nil) };

/* Terminal */
Id: UPPER_ID { $$ = $1 } | LOWER_ID { $$ = $1 };
Empty: {};

%%
var syntaxTree ast.Ast

func Parse(lex *lexer.Lexer) (ast.Ast) {
	syntaxTree = ast.New()
	yyParse(&lexerFrontend {
		lexer: *lex,
	})
	syntaxTree.Print()

	return syntaxTree
}