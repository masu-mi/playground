package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, c, k int
	fmt.Scan(&n, &c, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	ts := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		ts = append(ts, a)
	}
	sort.Sort(sort.IntSlice(ts))
	limit, load, numOfBus := -math.MaxInt32, 0, 0
	for _, t := range ts {
		if t > limit {
			numOfBus++
			load = 1
			limit = t + k
		} else if load == c {
			numOfBus++
			load = 1
			limit = t + k
		} else {
			load++
		}
	}
	fmt.Printf("%d\n", numOfBus)
}
