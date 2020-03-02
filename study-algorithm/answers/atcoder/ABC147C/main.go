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

func parseProblem(r io.Reader) (int, []testimony) {
	var n int
	fmt.Fscanf(r, "%d", &n)

	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	testimonies := []testimony{}
	for i := 0; i < n; i++ {
		testimonies = append(testimonies, parseTestimony(sc))
	}
	return n, testimonies
}

type testimony struct{ h, l uint32 }

func parseTestimony(sc *bufio.Scanner) (t testimony) {
	n := scanInt(sc)

	t.l, t.h = math.MaxUint32, 0
	for i := 0; i < n; i++ {
		idx := scanInt(sc)
		horneset := scanInt(sc)
		if horneset == 0 {
			t.l ^= 1 << uint(idx-1)
		} else {
			t.h |= 1 << uint(idx-1)
		}
	}
	return
}

func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func resolve(n int, testimonies []testimony) int {
	max := uint32(0)
	for i := 0; i < 1<<uint32(n); i++ {
		if checkAllHornest(n, uint32(i), testimonies) {
			numH := onesCount(uint64(i))
			if max < uint32(numH) {
				max = uint32(numH)
			}
		}
	}
	return int(max)
}

func checkAllHornest(n int, i uint32, testimonies []testimony) bool {
	flag := 1
	idx := 0
	for idx < n {
		if i&uint32(flag) == uint32(flag) {
			if i&testimonies[idx].l != i {
				return false
			}
			if i|testimonies[idx].h != i {
				return false
			}
		}
		flag <<= 1
		idx++
	}
	return true
}

func onesCount(x uint64) (num int) {
	const m0 = 0x5555555555555555 // 01010101 ...
	const m1 = 0x3333333333333333 // 00110011 ...
	const m2 = 0x0f0f0f0f0f0f0f0f // 00001111 ...

	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}
