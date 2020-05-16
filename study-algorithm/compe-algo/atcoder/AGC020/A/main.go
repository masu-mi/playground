package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	var length int
	if diff := a - b; diff > 0 {
		length = diff
	} else {
		length = -diff
	}
	if length%2 == 0 {
		fmt.Println("Alice")
	} else {
		fmt.Println("Borys")
	}
}
