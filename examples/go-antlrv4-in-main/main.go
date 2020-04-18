package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func main() {
	input := antlr.NewInputStream("(12+3+8)*4")
	lexer := NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := NewCalcParser(stream)
	tree := p.Prog()

	fmt.Println(tree.ToStringTree([]string{}, p)) // ← 解析木を LISP-style で表示する，最初の文字列配列をなぜ与えるのかはよく分からん
}
