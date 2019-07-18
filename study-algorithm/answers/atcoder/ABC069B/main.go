package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	if len(s) < 3 {
		panic(1)
	}
	fmt.Printf("%c%d%c", s[0], len(s)-2, s[len(s)-1])
}
