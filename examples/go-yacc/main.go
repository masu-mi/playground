package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	l, _ := parse(strings.NewReader(os.Args[1]))
	fmt.Fprintf(os.Stderr, "%#v\n", l.result)
	fmt.Println(eval(l.result))
}

func eval(e Expression) int {
	switch e.(type) {
	case BinOpExpr:
		left := eval(e.(BinOpExpr).left)
		right := eval(e.(BinOpExpr).right)

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
