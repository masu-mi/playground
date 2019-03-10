package main

import "fmt"

func main() {
	var n, prefixNum, suffixNum int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	nested := 0
	for i := 0; i < n; i++ {
		cur := s[i]
		if cur == '(' {
			nested++
		} else {
			nested--
		}
		if nested < 0 {
			prefixNum++
			nested = 0
		}
	}
	if nested > 0 {
		suffixNum += nested
	}
	for i := 0; i < prefixNum; i++ {
		fmt.Print("(")
	}
	fmt.Printf("%s", s)
	for i := 0; i < suffixNum; i++ {
		fmt.Print(")")
	}
}
