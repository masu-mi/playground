package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func main() {
	if resolve(parseProblem(os.Stdin)) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func parseProblem(r io.Reader) []block {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	bs := []block{}
	for i := 0; i < n; i++ {
		t := scanString(sc)
		b := toBlock(t)
		bs = append(bs, b)
	}
	sort.Sort(blockSlice(bs))
	return bs
}

type blockSlice []block

func (b blockSlice) Len() int {
	return len(b)
}

func (b blockSlice) Less(i, j int) bool {
	iP := b[i].r - b[i].l
	jP := b[j].r - b[j].l
	if (iP >= 0 && jP < 0) || (iP < 0 && jP >= 0) {
		// other phase
		return iP > jP
	} else if iP >= 0 {
		// 1st
		return b[i].l < b[j].l
	}
	return b[i].r > b[j].r
}

func (b blockSlice) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

type block struct {
	l, r int
}

func toBlock(s string) block {
	// l ')' num
	// r '(' num
	lc := 0
	rc := 0
	for i := 0; i < len(s); i++ {
		if rc == 0 {
			if s[i] == ')' {
				lc++
			} else {
				rc++
			}
		} else {
			if s[i] == ')' {
				rc--
			} else {
				rc++
			}
		}
	}
	return block{l: lc, r: rc}
}

func resolve(bs []block) bool {
	depth := 0
	for i := 0; i < len(bs); i++ {
		depth -= bs[i].l
		if depth < 0 {
			return false
		}
		depth += bs[i].r
	}
	return depth == 0
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
