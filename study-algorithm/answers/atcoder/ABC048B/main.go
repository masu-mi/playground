package main

import (
	"fmt"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)
	num := b/x - a/x
	if a%x == 0 {
		num++
	}
	fmt.Printf("%d\n", num)
}
