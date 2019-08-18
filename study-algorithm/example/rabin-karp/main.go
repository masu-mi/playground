package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sub, text := getProblem()
	for _, idx := range findPatternInText(sub, text) {
		fmt.Printf("ids: %d: (%s)\n", idx, text[idx:idx+len(sub)])
	}
}

func findPatternInText(pat, text string) (indexes []int) {
	// calc hash of pattern
	patHash := calcRollingHash(pat)
	textHash := calcRollingHash(text[0:len(pat)])

	headPart := calcHeadCharPart(len(pat))
	for i := len(pat); i <= len(text); i++ {
		if patHash == textHash {
			if pat == text[i-len(pat):i] {
				indexes = append(indexes, i-len(pat))
			}
		}
		if i >= len(text) {
			break
		}
		textHash = updateRollingHash(textHash, headPart, text[i-len(pat)], text[i])
	}
	return indexes
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

func getProblem() (sub string, text string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	sc.Scan()
	sub = sc.Text()
	sc.Scan()
	text = sc.Text()
	return
}
