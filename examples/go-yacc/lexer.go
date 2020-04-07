package main

import (
	"text/scanner"

	"github.com/k0kubun/pp"
)

type Lexer struct {
	scanner.Scanner
	result Expression
}

func (l *Lexer) Lex(lval *yySymType) int {
	token := int(l.Scan())
	if token == scanner.Int {
		token = NUMBER
	} else if token == scanner.Ident && l.TokenText() == "if" {
		token = IF
	}

	if l.TokenText() == "+" {
		pp.Println(scanner.Ident)
		pp.Println(token)
	}
	lval.token = Token{token: token, literal: l.TokenText()}
	return token
}

func (l *Lexer) Error(e string) {
	panic(e)
}
