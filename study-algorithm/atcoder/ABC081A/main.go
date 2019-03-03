package main

import "fmt"

var toNum = map[int]int{
	0: 0, 1: 1, 2: 1, 3: 2,
	4: 1, 5: 2, 6: 2, 7: 3,
}

func main() {
	var i int
	fmt.Scanf("%b", &i)
	fmt.Printf("%d\n", toNum[i])
}
