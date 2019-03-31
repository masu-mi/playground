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
	oddNum := 0
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if a%2 == 1 {
			oddNum++
		}
	}
	if oddNum%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
