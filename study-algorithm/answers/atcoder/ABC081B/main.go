package main

import "fmt"

func main() {
	var n, a, result int
	fmt.Scanf("%d", &n)
	for ; n > 0; n-- {
		fmt.Scanf("%d", &a)
		result |= a
	}
	for ; result&1 == 0; result >>= 1 {
		n++
	}
	fmt.Printf("%d\n", n)
}
