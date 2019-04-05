package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var q int
	fmt.Scan(&q)
	type query struct{ l, r int }
	qs := make([]query, q)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var max int
	for i := 0; i < q; i++ {
		l := nextInt(sc)
		r := nextInt(sc)
		qs[i] = query{l, r}
		if max < r {
			max = r
		}
	}
	checkPrime := isPrime(max)
	culumtive := make([]int, (max+3)/2+1)
	sum := 0
	for i := 0; i <= (max+3)/2; i++ {
		culumtive[i] = sum
		t := i*2 + 1
		if _, ok := checkPrime[t]; ok {
			if _, ok := checkPrime[(t+1)/2]; ok {
				sum++
			}
		}
	}
	for _, q := range qs {
		fmt.Printf(
			"%d\n",
			culumtive[(q.r+3)/2-1]-culumtive[(q.l+1)/2-1],
		)
	}
}

func isPrime(max int) map[int]struct{} {
	m := map[int]struct{}{}
	for i := 2; i <= max; i++ {
		m[i] = struct{}{}
	}
	for i := 2; i <= max; i++ {
		for j := 2; j <= max/i; j++ {
			delete(m, i*j)
		}
	}
	return m
}

func nextInt(sc *bufio.Scanner) int {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	return a
}
