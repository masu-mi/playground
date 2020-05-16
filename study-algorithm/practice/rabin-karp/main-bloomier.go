package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	subs, text := getProblem()
	c := newChecker(subs)
	for _, idx := range findPatternInText(c, text) {
		fmt.Printf("ids: %d: (%s)\n", idx, text[idx:idx+c.length])
	}
}

func findPatternInText(c *checker, text string) (indexes []int) {
	// calc hash of pattern
	textHash := calcRollingHash(text[0:c.length])
	headPart := calcHeadCharPart(c.length)
	for i := c.length; i <= len(text); i++ {
		if c.check(textHash) {
			if c.findPart(textHash, text[i-c.length:i]) != "" {
				indexes = append(indexes, i-c.length)
			}
		}
		if i >= len(text) {
			break
		}
		textHash = updateRollingHash(textHash, headPart, text[i-c.length], text[i])
	}
	return indexes
}

type checker struct {
	// filter size 64 bit
	bloomierFilter int64
	length         int
	items          map[int64][]string
}

func newChecker(subs []string) *checker {
	c := &checker{items: make(map[int64][]string)}
	c.length = len(subs[0])
	for _, str := range subs {
		if len(str) != c.length {
			fmt.Fprintln(os.Stderr, "require same length")
			os.Exit(1)
		}
		// use hash family(64) which is defined as spliting source hash to each bit
		hash := calcRollingHash(str)
		c.bloomierFilter |= hash
		c.items[hash] = append(c.items[hash], str)
	}
	return c
}

func (c checker) check(hash int64) bool {
	return c.bloomierFilter&hash == hash
}

func (c checker) findPart(hash int64, input string) string {
	for _, candidate := range c.items[hash] {
		if candidate == input[0:c.length] {
			return candidate
		}
	}
	return ""
}

const base = 8861
const divisor = 1 << 10

func calcRollingHash(text string) int64 {
	var result int64 = 0
	for i := 0; i < len(text); i++ {
		result = result * base % divisor
		result += (int64)(text[i])
	}
	return result % divisor
}

func calcHeadCharPart(strLen int) int64 {
	part := int64(1)
	for i := strLen - 1; i > 0; i-- {
		part = part * base % divisor
	}
	return part
}

func updateRollingHash(hash int64, headPart int64, headChar, tailChar byte) int64 {
	b := (hash - headPart*int64(headChar)%divisor)
	if b < 0 {
		b += divisor
	}
	return (b*base + int64(tailChar)) % divisor
}

func getProblem() (subs []string, text string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	for i := 0; i < n; i++ {
		sc.Scan()
		sub := sc.Text()
		subs = append(subs, sub)
	}
	sc.Scan()
	text = sc.Text()
	return
}
