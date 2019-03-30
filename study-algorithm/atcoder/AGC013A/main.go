package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	const (
		preload = 0 + iota
		initialized
		inc
		dec
	)
	mode := preload
	chunkNum := 0
	var pre int
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if mode == preload {
			chunkNum++
			mode = initialized
		} else if mode == initialized {
			if a > pre {
				mode = inc
			} else if a < pre {
				mode = dec
			}
		} else if mode == inc {
			if a < pre {
				chunkNum++
				mode = initialized
			}
		} else if mode == dec {
			if a > pre {
				chunkNum++
				mode = initialized
			}
		}
		pre = a
	}
	fmt.Printf("%d\n", chunkNum)
}
