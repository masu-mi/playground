package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	fmt.Printf("%d\n", numOfOddDigitNum(n))
}

func numOfOddDigitNum(n int) (count int) {
	for i := 1; i <= n; i++ {
		if digit(i)%2 == 1 {
			count++
		}
	}
	return count
}
func digit(n int) int {
	if n == 0 {
		return 0
	}
	digit := 0
	for i := 1; i <= n; i *= 10 {
		digit++
	}
	return digit
}
