package main

import "fmt"

const divisor = 1000000000 + 7

func main() {
	var n int
	fmt.Scan(&n)
	level := 1
	for i := 0; i < n; i++ {
		level = level * (i + 1) % divisor
	}
	fmt.Printf("%d\n", level)
}
