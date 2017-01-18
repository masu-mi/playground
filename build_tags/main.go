package main

import (
	"fmt"

	"./debugaid"
)

func main() {
	input := 10
	debugaid.Assert(input > 0)
	fmt.Printf("input:%d\n", input)
}
