package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

func main() {
	resolve(parseProblem(os.Stdin))
}

func parseProblem(r io.Reader) (string, chan query) {
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanWords) // bufio.ScanLines
	const (
		initialBufSize = 100000
		maxBufSize     = 100000
	)
	buf := make([]byte, initialBufSize)
	sc.Buffer(buf, maxBufSize)

	sc.Scan()
	s := sc.Text()
	n := scanint(sc)

	ch := make(chan query)
	go func() {
		for i := 0; i < n; i++ {
			ch <- parseQuery(sc)
		}
		close(ch)
	}()
	return s, ch
}

const (
	Switch = 1 + iota
	Append
)

type query struct {
	cmd int
	sub int
	c   string
}

func parseQuery(sc *bufio.Scanner) query {
	cmd := scanint(sc)
	if cmd == Switch {
		return query{cmd, 0, ""}
	}
	sub := scanint(sc)
	sc.Scan()
	t := sc.Text()
	return query{Append, sub, t}
}

func scanint(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}

func resolve(base string, queries chan query) {
	state := 0
	fixies := [2][]string{}
	for q := range queries {
		if q.cmd == Switch {
			state += 1
		} else {
			// state 0; 1 -> 1(prefix), 2 -> 0 (suffix)
			fixies[(state+q.sub)%2] = append(fixies[(state+q.sub)%2], q.c)
		}
	}

	normal := state%2 == 0
	prefix := fixies[(state+1)%2]
	suffix := fixies[(state+2)%2]

	var middle string
	if normal {
		middle = base
	} else {
		middle = reverseString(base)
	}
	writeFlattenStrings(
		os.Stdout,
		reverseStrings(prefix...),
		[]string{middle},
		suffix,
	)
}

func reverseString(str string) string {
	buf := bytes.NewBufferString("")
	writeReverse(buf, str)
	return buf.String()
}

func reverseStrings(strs ...string) (rev []string) {
	for i := len(strs); i > 0; i-- {
		rev = append(rev, strs[i-1])
	}
	return
}

func writeStrings(w io.Writer, strs ...string) {
	ch := make(chan string)
	go func() {
		for _, s := range strs {
			ch <- s
		}
		close(ch)
	}()
	writeWithChanString(w, ch)
}

func writeFlattenStrings(r io.Writer, strss ...[]string) {
	ch := make(chan string)
	go func() {
		for _, strs := range strss {
			for _, s := range strs {
				ch <- s
			}
		}
		close(ch)
	}()
	writeWithChanString(r, ch)
}

func writeWithChanString(r io.Writer, strs chan string) {
	for s := range strs {
		r.Write(([]byte)(s))
	}
}

func writeReverse(r io.Writer, str string) {
	runes := []rune(str)
	for i := len(runes); i > 0; i-- {
		r.Write(([]byte)(string(runes[i-1])))
	}
}
