package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) (n, m int, mp []int, ms []uint) {
	fmt.Fscan(r, &n, &m)

	ms = make([]uint, m)
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		for j := 0; j < k; j++ {
			sc.Scan()
			ss, _ := strconv.Atoi(sc.Text())
			ms[i] |= 1 << (uint(ss) - 1)
		}
	}
	mp = make([]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		mp[i] = p
	}
	return n, m, mp, ms
}

func resolve(n, m int, mp []int, ms []uint) int {
	var count int
	for i := 0; i < 1<<uint(n); i++ {
		if allOn(uint(i), m, mp, ms) {
			count++
		}
	}
	return count
}
func allOn(l uint, m int, mp []int, ms []uint) bool {
	for i := 0; i < m; i++ {
		if mp[i] != int(onesCount(ms[i]&l)%2) {
			return false
		}
	}
	return true
}

func onesCount(bits uint) (num uint) {
	num = (bits >> 1) & 03333333333
	num = bits - num - ((num >> 1) & 03333333333)
	num = ((num + (num >> 3)) & 0707070707) % 077
	return
}
