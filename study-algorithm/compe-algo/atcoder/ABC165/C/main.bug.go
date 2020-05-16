package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (int, int, []check) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	n := scanInt(sc)
	m := scanInt(sc)
	q := scanInt(sc)
	scores := make([]check, q)
	for i := 0; i < q; i++ {
		scores[i] = check{
			a: scanInt(sc),
			b: scanInt(sc),
			c: scanInt(sc),
			d: scanInt(sc),
		}
	}
	return n, m, scores
}

type check struct {
	a, b, c, d int
}

func resolve(n, m int, checks []check) int {
	result := 0
	pp.Println(checks)
	for v := range bitCombinationsWithSize(m, n) {
		fmt.Printf("B%06b\n", v)
		sum := 0
		for _, c := range checks {
			count := 0
			av := 0
			bv := 0
			aCounted := false
			for i := 0; i < m; i++ {
				if v&uint(1<<uint(i)) != 0 {
					count++
				}
				if !aCounted && count == c.a {
					aCounted = true
					av = i + 1
				}
				if aCounted && count == c.b {
					bv = i + 1
					break
				}
			}
			if bv-av == c.c {
				sum += c.d
			}
		}
		if result < sum {
			result = sum
		}
	}
	return result
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
