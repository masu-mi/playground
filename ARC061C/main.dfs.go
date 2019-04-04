package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Printf("%d\n", dfs(0, s))
}

func dfs(sum int, s string) int {
	result := sum
	for i := 1; i < len(s); i++ {
		a, _ := strconv.Atoi(s[0:i])
		result += dfs(sum+a, s[i:len(s)])
	}
	a, _ := strconv.Atoi(s)
	return result + a
}
