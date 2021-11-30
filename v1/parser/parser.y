%{
package parser

import "github.com/Goamaral/goa-lang/v1/ast"
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
	value string
	node *ast.Node
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <value> FunctionDefinition

// same for terminals
%token <value> DEF DO END ID

%start FunctionDefinition

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies precedence for unary minus */

%%
FunctionDefinition: DEF ID DO END { syntaxTree.AddFuncDef($2) };

%%
var syntaxTree ast.Ast

func Parse(sourceCode string) (ast.Ast) {
	syntaxTree = ast.New()

	lex := lexer{code: sourceCode}
	yyParse(&lex)

	lex.Print()
	syntaxTree.Print()

	return syntaxTree
}