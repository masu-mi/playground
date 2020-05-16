package main

import "fmt"

func main() {
	// https://imoz.jp/algorithms/imos_method.html
	var t, c int
	fmt.Scan(&t, &c)
	nums := make([]int, t)
	for i := 0; i < c; i++ {
		var s, e int
		fmt.Scan(&s, &e)
		nums[s]++
		nums[e]--
	}
	// sim
	max := 0
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i-1] + nums[i]
		if max < nums[i] {
			max = nums[i]
		}
	}
	fmt.Printf("%d\n", max)
}
