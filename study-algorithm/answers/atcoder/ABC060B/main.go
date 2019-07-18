package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	for i := 1; i < b+1; i++ {
		mod := i * a % b
		if mod == c {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
	return
}
