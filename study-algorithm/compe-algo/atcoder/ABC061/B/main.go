package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	roads := make([]int, n)
	for i := 0; i < 2*m; i++ {
		sc.Scan()
		i, _ := strconv.Atoi(sc.Text())
		roads[i-1]++
	}
	for _, k := range roads {
		fmt.Printf("%d\n", k)
	}
}
