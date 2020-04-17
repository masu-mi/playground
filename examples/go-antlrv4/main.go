package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"github.com/masu-mi/playground/examples/go-antlrv4/parser"
)

func main() {
	input := antlr.NewInputStream("12+3*4") // ← 入力

	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcParser(stream)
	tree := p.Prog() // ← ルールのトップに対応する関数が用意されてるのでそれを起動すると解析して解析木を返す

	fmt.Println(tree.ToStringTree([]string{}, p)) // ← 解析木を LISP-style で表示する，最初の文字列配列をなぜ与えるのかはよく分からん
}
