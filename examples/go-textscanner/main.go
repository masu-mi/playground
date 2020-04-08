package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"text/scanner"
	"unicode"
)

const inputString = `
val x := {
  "id": 2500,
  'name' => "object x", // name is string
}

# This  /* comment

*/
val y-z = x


if x.name != "simple str" {
  print(200+30**4)
}
`

type token struct {
	token rune
	text  string
}

func main() {
	type testCase struct {
		mode        uint
		isIdentRune func(ch rune, i int) bool
	}
	c := []testCase{
		testCase{
			mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanComments | scanner.ScanChars,
		},
		testCase{
			mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanComments | scanner.SkipComments | scanner.ScanChars,
		},
		testCase{
			mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanStrings | scanner.ScanChars,
		},
		testCase{
			mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanStrings | scanner.ScanChars | scanner.ScanComments,
		},
		testCase{
			mode: scanner.ScanIdents | scanner.ScanInts | scanner.ScanChars | scanner.ScanComments,
			isIdentRune: func(ch rune, i int) bool {
				if ch == '-' && i > 0 || unicode.IsLetter(ch) || unicode.IsDigit(ch) && i > 0 {
					return true
				}
				return false
			},
		},
	}
	sc := &scanner.Scanner{}
	sc = sc.Init(bytes.NewBufferString(inputString))
	n, _ := strconv.Atoi(os.Args[1])
	sc.Mode = c[n].mode
	sc.IsIdentRune = c[n].isIdentRune
	result := []token{}
	for tok := sc.Scan(); tok != scanner.EOF; tok = sc.Scan() {
		result = append(result, token{token: tok, text: sc.TokenText()})
	}
	fmt.Printf("%v\n", result)
}
