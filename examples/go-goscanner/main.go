package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	src := []byte(`
  val x = 100.0

  cos(x) + 1i*sin(x) // Euler
  `)

	x := 1i*(3+4i) + 2i
	fmt.Println(x)

	// Initialize the scanner.
	var s scanner.Scanner
	fset := token.NewFileSet()                      // positions are relative to fset
	file := fset.AddFile("", fset.Base(), len(src)) // register input "file"
	s.Init(file, src, nil /* no error handler */, scanner.ScanComments)

	// Repeated calls to Scan yield the token sequence found in the input.
	for {
		pos, tok, lit := s.Scan()
		if tok == token.EOF {
			break
		}
		fmt.Printf("%s\t%s\t%q\n", fset.Position(pos), tok, lit)
	}
}
