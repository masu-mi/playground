package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	tmp, denom := n, 0
	for tmp != 0 {
		denom += tmp % 10
		tmp = tmp / 10
	}
	if n%denom == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
