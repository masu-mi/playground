package main

import "fmt"

func main() {
	var k, s int
	fmt.Scan(&k, &s)
	var numOfSatisfied int
	for x := 0; x <= k; x++ {
		for y := 0; y <= k; y++ {
			rest := s - (x + y)
			if rest >= 0 && rest <= k {
				numOfSatisfied++
			}
		}
	}
	fmt.Printf("%d\n", numOfSatisfied)
}
