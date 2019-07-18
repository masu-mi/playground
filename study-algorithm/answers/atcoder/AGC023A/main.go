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
	sum, patternNum := 0, 0
	sums := make(map[int]int)
	for i := 0; i < n; i++ {
		sums[sum]++

		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sum += a
		patternNum += sums[sum]
	}
	fmt.Printf("%d\n", patternNum)
}
