package main

import "fmt"

func main() {
	var i, n int
	fmt.Scanf("%b", &i)
	for ; i != 0; i &= i - 1 {
		n++
	}
	fmt.Printf("%d\n", n)
}
