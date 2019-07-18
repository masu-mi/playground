package main

import "fmt"

func main() {
	var a, b, count int
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		if palindromic(i) {
			count++
		}
	}
	fmt.Printf("%d\n", count)
}

func palindromic(input int) bool {
	tmp := input
	digits := []int{}
	for tmp != 0 {
		digits = append(digits, tmp%10)
		tmp /= 10
	}
	for i := 0; i < len(digits)/2; i++ {
		if digits[i] != digits[len(digits)-1-i] {
			return false
		}
	}
	return true
}
