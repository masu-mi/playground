package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var evenNum, max int
	vals := make([]int, 3)
	for i := range vals {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		vals[i] = a
		if a%2 == 0 {
			evenNum++
		}
		if max < a {
			max = a
		}
	}
	const (
		even = 0 + iota
		odd
	)
	var lifted, ops int
	if evenNum%2 == 0 {
		lifted = even
		ops += evenNum / 2
	} else {
		lifted = odd
		ops += (len(vals) - evenNum) / 2
	}
	for _, v := range vals {
		ops += (max - v) / 2
		if (max+v)%2 == 1 {
			if max%2 == lifted {
				ops++
			}
		}
	}
	fmt.Printf("%d\n", ops)
}
