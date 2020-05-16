package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"

	"github.com/k0kubun/pp"
)

func main() {
	example()
}

func example() {
	b := newBIT(10)
	b.add(1, 1)
	b.add(2, 2)
	b.add(3, 3)
	b.add(5, 2)
	b.add(6, -1)
	b.add(7, -1)
	b.add(8, -1)
	b.add(9, 20)
	pp.Println(b)
	for i := 1; i <= 10; i++ {
		fmt.Printf(
			"[%d] = %d,	[0..%d] = %d,	[%d..%d] = %d\n",
			i, b.value(i),
			i, b.accum(i),
			i, i-2, b.sum(i-2, i),
		)
	}
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

type BIT struct {
	size int
	list []int64
}

var ErrInvalidIndex = errors.New("BIT index require > 0")

func newBIT(size int) *BIT {
	return &BIT{size: size, list: make([]int64, size+1)}
}

func (b *BIT) add(index int, v int64) {
	if index == 0 {
		panic(ErrInvalidIndex)
	}
	for i := index; i <= b.size; i += (i & -i) {
		b.list[i] += v
	}
}

func (b *BIT) value(i int) int64 {
	return b.accum(i) - b.accum(i-1)
}

func (b *BIT) sum(s, e int) int64 {
	return b.accum(e) - b.accum(s-1)
}

func (b *BIT) accum(index int) (a int64) {
	if index <= 0 {
		return 0
	}
	for i := index; i > 0; i -= (i & -i) {
		a += b.list[i]
	}
	return a
}
