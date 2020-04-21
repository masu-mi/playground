package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	parseProblem(os.Stdin)
}

func parseProblem(r io.Reader) {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	n := scanInt(sc)
	str := scanString(sc)

	bu := bytes.NewBufferString("")
	for i := 0; i < len(str); i++ {
		c := str[i] + byte(n)
		if c > 'Z' {
			c = c - 'Z' + 'A' - 1
		}
		bu.WriteByte(c)
	}
	fmt.Println(bu.String())
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
