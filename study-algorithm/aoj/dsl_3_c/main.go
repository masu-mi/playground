package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	values := make([]int, 0, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		values = append(values, a)
	}
	// this problem's feature have monotonically non-increasing.
	// and 10^5 * 500 < 10^10
	// so I decide to use two pointers.
	for i := 0; i < q; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		fmt.Printf("%d\n", countRange(a, values))
	}
}
func countRange(max int, values []int) (count int) {
	var num int
	length := len(values)
	l, r := 0, 0
	current := values[l]
RESEARCH:
	for l < length {
		if l > r {
			r = l
			current += values[r]
		}
		for true {
			if current > max {
				break
			}
			num += r - l + 1
			r++
			if r >= length {
				break RESEARCH
			}
			current += values[r]
		}
		current -= values[l]
		l++
	}
	return num
}
