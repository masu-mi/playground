package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l := new(Lexer)
	l.Init(strings.NewReader(os.Args[1]))
	//go:generate goyacc -o parser.go parser.go.y
	yyParse(l)
	fmt.Fprintf(os.Stderr, "%#v\n", l.result)
	fmt.Println(Eval(l.result))
}

func Eval(e Expression) int {
	switch e.(type) {
	case BinOpExpr:
		left := Eval(e.(BinOpExpr).left)
		right := Eval(e.(BinOpExpr).right)

		switch e.(BinOpExpr).operator {
		case '+':
			return left + right
		case '-':
			return left - right
		case '*':
			return left * right
		case '/':
			return left / right
		}
	case NumExpr:
		num, _ := strconv.Atoi(e.(NumExpr).literal)
		return num
	}
	return 0
}
