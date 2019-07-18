package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var q int
	fmt.Scan(&q)
	results := make([]int, 0, q)
	for i := 0; i < q; i++ {
		results = append(results, resolve())
	}
	for _, v := range results {
		fmt.Printf("%d\n", v)
	}
}

func resolve() int {
	var n, max int
	fmt.Scan(&n, &max)
	nums := make([]int, 0, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		nums = append(nums, a)
	}
	var l, r int
	current := 0
	min := math.MaxInt32
	for true {
		for current < max {
			r++
			if r > n {
				break
			}
			current += nums[r-1]
		}
		if current >= max {
			if length := r - l; min > length {
				min = length
			}
		}
		current -= nums[l]
		l++
		if l >= n {
			break
		}
	}
	return min
}
