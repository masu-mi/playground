package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	nums := map[int]int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		nums[a-1]++
		nums[a]++
		nums[a+1]++
	}
	var max int
	for _, v := range nums {
		if max < v {
			max = v
		}
	}
	fmt.Printf("%d\n", max)
}
