package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// for http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=0516
	// referenced by https://qiita.com/drken/items/56a6b68edef8fc605821#fn1
	var n, k int
	fmt.Scan(&n, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	cumulative := make([]int, n)
	sc.Scan()
	cumulative[0], _ = strconv.Atoi(sc.Text())
	for i := 1; i < n; i++ {
		sc.Scan()
		cumulative[i], _ = strconv.Atoi(sc.Text())
		cumulative[i] += cumulative[i-1]
	}
	sc.Scan() // 0
	sc.Scan() // 0
	var max int
	for i := k; i < n; i++ {
		sum := cumulative[i] - cumulative[i-k]
		if max < sum {
			max = sum
		}
	}
	fmt.Printf("%d\n", max)
}
