package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	kinds := map[int]int{}
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		kinds[a]++
	}
	nums := make([]int, 0, len(kinds))
	for _, n := range kinds {
		nums = append(nums, n)
	}
	sort.Sort(sort.IntSlice(nums))
	sum := 0
	for i := 0; i < len(kinds)-k; i++ {
		sum += nums[i]
	}
	fmt.Printf("%d\n", sum)
}
