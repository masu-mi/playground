package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var tmp, step int
	exists := map[int]struct{}{}
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		if _, ok := exists[tmp]; !ok {
			step++
			exists[tmp] = struct{}{}
		}
	}
	fmt.Printf("%d\n", step)
}
