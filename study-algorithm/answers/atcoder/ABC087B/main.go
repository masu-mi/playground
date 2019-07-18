package main

import (
	"fmt"
)

func main() {
	var a, b, c, x int
	fmt.Scanf("%d\n%d\n%d\n%d", &a, &b, &c, &x)
	fmt.Printf("%d", count(a, b, c, x))
}

func count(a, b, c, x int) (n int) {

	sum := x / 50
	var rests = make(map[int]struct{}, c)
	for ; c > -1; c-- {
		rests[sum-c] = struct{}{}
	}
	for i := 0; i < a+1; i++ {
		for j := 0; j < b+1; j++ {
			if _, ok := rests[10*i+2*j]; ok {
				n++
			}
		}
	}
	return n
}
