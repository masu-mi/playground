package main

import "fmt"

func main() {
	var n, a int
	fmt.Scan(&n)
	fmt.Scan(&a)
	if a/500 > 1 {
		fmt.Println("Yes")
		return
	}
	if n%500 <= a {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
