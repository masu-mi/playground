package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	spaceNum, sum := len(s)-1, 0
	for i := 0; i < 1<<uint(spaceNum); i++ {
		sum += expNum(s, i)
	}
	fmt.Printf("%d\n", sum)
}

func expNum(s string, selection int) (sum int) {
	rest := selection
	buf := bytes.NewBuffer([]byte{})
	for i := 0; i < len(s); i++ {
		buf.Write([]byte{s[i]})
		if rest&1 == 1 {
			a, _ := strconv.Atoi(buf.String())
			buf = bytes.NewBuffer([]byte{})
			sum += a
		}
		rest >>= 1
	}
	if buf.Len() > 0 {
		a, _ := strconv.Atoi(buf.String())
		buf = bytes.NewBuffer([]byte{})
		sum += a
	}
	return sum
}
