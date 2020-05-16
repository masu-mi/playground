package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, int, int, []int, [][]int) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n, m, x := scanInt(sc), scanInt(sc), scanInt(sc)
	cs := make([]int, n)
	as := make([][]int, n)
	for i := 0; i < n; i++ {
		cs[i] = scanInt(sc)
		as[i] = make([]int, m)
		for j := 0; j < m; j++ {
			as[i][j] = scanInt(sc)
		}
	}
	return n, m, x, cs, as
}

func resolve(n, m, x int, cs []int, as [][]int) int {
	ok := false
	min := math.MaxInt64
	for bits := range bitCombinations(n) {
		cost := 0
		skils := make([]int, m)
		for i := 0; i < n; i++ {
			if bits>>uint(i)&1 == 1 {
				cost += cs[i]
				for j := 0; j < m; j++ {
					skils[j] += as[i][j]
				}
			}
		}
		cl := true
		for j := 0; j < m; j++ {
			if skils[j] < x {
				cl = false
			}
		}
		ok = ok || cl
		if cl && min > cost {
			min = cost
		}
	}
	if ok {
		return min
	}
	return -1
}

func bitCombinations(num int) chan uint {
	ch := make(chan uint)
	go func() {
		defer close(ch)
		for i := 0; i < 1<<uint(num); i++ {
			ch <- uint(i)
		}
	}()
	return ch
}

func bitCombinationsOverSubsets(nums ...int) chan uint {
	ch := make(chan uint)
	s := uint(0)
	for _, v := range nums {
		s |= 1 << uint(v)
	}
	go func() {
		defer close(ch)
		for bit := s; ; bit = (bit - 1) & s {
			ch <- uint(bit)
			if bit == 0 {
				break
			}
		}
	}()
	return ch
}

func bitCombinationsWithSize(num, size int) chan uint {
	ch := make(chan uint)
	bit := uint(1<<uint(size) - 1)
	go func() {
		defer close(ch)
		for ; bit < 1<<uint(num); bit = nextBitCombination(uint(bit)) {
			ch <- bit
		}
	}()
	return ch
}

func nextBitCombination(cur uint) uint {
	x := cur & -cur // rightest bit only         '10100' -> '00100'
	y := cur + x    // carry at rightest 1-block '10111' -> '11000'
	return (((cur & ^y) / x) >> 1) | y
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
func scanString(sc *bufio.Scanner) string {
	sc.Scan()
	return sc.Text()
}
