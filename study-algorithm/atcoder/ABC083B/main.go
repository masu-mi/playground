package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scanf("%d %d %d", &n, &a, &b)
	fmt.Printf("%d\n", count(n, a, b))
}

func count(n, a, b int) (result int) {
	for i := 1; i <= n; i++ {
		if isValidNumber(i, a, b) {
			result += i
		}
	}
	return result
}

func isValidNumber(i, a, b int) bool {
	var sum int
	for ; i > 0; i /= 10 {
		sum += i % 10
	}
	return sum >= a && sum <= b
}
