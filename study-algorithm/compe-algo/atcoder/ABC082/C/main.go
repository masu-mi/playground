package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	nums := make(map[int]int)
	var num int
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		nums[num] += 1
	}
	var shoudBeRemoved int
	for num, count := range nums {
		if count < num {
			shoudBeRemoved += count
		} else if count > num {
			shoudBeRemoved += count - num
		}
	}
	fmt.Printf("%d\n", shoudBeRemoved)
}
