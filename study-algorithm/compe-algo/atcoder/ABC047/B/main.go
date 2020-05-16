package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var w, h, n int
	fmt.Scan(&w, &h, &n)
	current := map[int]int{1: 0, 2: w, 3: 0, 4: h}
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		y, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if a == 1 {
			if current[1] < x {
				current[1] = x
			}
		} else if a == 2 {
			if current[2] > x {
				current[2] = x
			}
		} else if a == 3 {
			if current[3] < y {
				current[3] = y
			}
		} else {
			if current[4] > y {
				current[4] = y
			}
		}
	}
	var restHight, restWidth int
	if w := current[2] - current[1]; w > 0 {
		restWidth = w
	}
	if h := current[4] - current[3]; h > 0 {
		restHight = h
	}
	restArea := restWidth * restHight
	fmt.Printf("%d\n", restArea)
}
