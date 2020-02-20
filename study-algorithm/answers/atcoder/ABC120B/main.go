package main

import "fmt"

func main() {
	var a, b, k int
	fmt.Scanf("%d %d %d", &a, &b, &k)
	fmt.Printf("%d\n", findInt(a, b, k))
}

func findInt(a, b, k int) int {
	lim := min(a, b)
	count := 0
	for i := lim; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			count++
		}
		if count == k {
			return i
		}
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
