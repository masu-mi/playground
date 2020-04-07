%{
package main

import (
    "github.com/k0kubun/pp"
)

type Expression interface{}
type Token struct {
    token   int
    literal string
}

type NumExpr struct {
    literal string
    token   int
}
type BinOpExpr struct {
    left     Expression
    operator rune
    right    Expression
}
%}

%union{
    token Token
    expr  Expression
}

%type<expr> program
%type<expr> expr
%type<expr> if_state
%token<token> NUMBER
%token<token> IF

%left '+' '-'
%left '*' '/' '%'
%left '(' ')' '<' '>'

%%

program
    : expr
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }
    | if_state
    {
        $$ = $1
        yylex.(*Lexer).result = $$
    }

if_state
    : IF '(' expr ')'
    {
       pp.Println($1)
        $$ = $3
    }

expr
    : NUMBER
    {
        $$ = NumExpr{literal: $1.literal, token: $1.token}
        pp.Println($$)
    }
    | '<' expr '>'
    { $$ = $2 }
    | '(' expr ')'
    { $$ = $2 }
    | expr '+' expr
    { $$ = BinOpExpr{left: $1, operator: '+', right: $3} }
    | expr '-' expr
    { $$ = BinOpExpr{left: $1, operator: '-', right: $3} }
    | expr '*' expr
    { $$ = BinOpExpr{left: $1, operator: '*', right: $3} }
    | expr '/' expr
    { $$ = BinOpExpr{left: $1, operator: '/', right: $3} }

%%
