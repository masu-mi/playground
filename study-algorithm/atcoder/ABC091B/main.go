package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	scores := map[string]int{}
	var s string
	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		scores[s]++
	}
	var m int
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&s)
		scores[s]--
	}

	var maxScore int
	for _, score := range scores {
		if maxScore < score {
			maxScore = score
		}
	}
	fmt.Printf("%d\n", maxScore)
}
