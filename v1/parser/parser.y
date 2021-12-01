%{
package parser

import "github.com/Goamaral/goa-lang/v1/ast"
import "github.com/Goamaral/goa-lang/v1/parser/lexer"
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	value string
	node *ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <value> FunctionDefinition FunctionBody FunctionCall

// same for terminals
%token <value> DEF DO END ID LPAR RPAR HASH

%start FunctionDefinition

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
FunctionDefinition: DEF ID DO FunctionCall END { syntaxTree.AddFuncDef($2) };

FunctionBody: FunctionCall { };

FunctionCall: HASH ID LPAR RPAR { };

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