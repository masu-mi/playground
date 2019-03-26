package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	as := make([]int, 3*n)
	for i := 0; i < 3*n; i++ {
		sc.Scan()
		as[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Ints(as)
	var sum int
	for i := n; i < 3*n; i += 2 {
		sum += as[i]
	}
	fmt.Printf("%d\n", sum)
}
