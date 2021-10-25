%{
package main

import (
  "fmt"
  "os"
  "io/ioutil"
)
%}

// fields inside this union end up as the fields in a structure known
// as ${PREFIX}SymType, of which a reference is passed to the lexer.
%union{
  strVal string
}

// any non-terminal which returns a value needs a type, which is
// really a field name in the above union struct
%type <strVal> FunctionDefinition

// same for terminals
%token <strVal> DEF DO END ID

%start FunctionDefinition

%left '|'
%left '&'
%left '+'  '-'
%left '*'  '/'  '%'
%left UMINUS      /*  supplies  precedence  for  unary  minus  */

%%
FunctionDefinition: DEF ID DO END;

%%
func main() {
  var err error
  var bytes []byte

  if (len(os.Args) == 1) {
    bytes, err = ioutil.ReadAll(os.Stdin)
    if err != nil {
      fmt.Println("Error reading stdin")
      return
    }
  } else {
    bytes, err = ioutil.ReadFile(os.Args[1])
    if err != nil {
      fmt.Printf("Error reading %s\n", os.Args[1])
      return
    }
  }

  lexer := yyLex{code: string(bytes)}
  yyParse(&lexer)
}