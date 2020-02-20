package main

import (
	"fmt"
)

func main() {

	acceptable := map[byte]struct{}{
		'A': struct{}{},
		'C': struct{}{},
		'G': struct{}{},
		'T': struct{}{},
	}

	var str string
	fmt.Scanf("%s", &str)

	var max int

	for i := 0; i < len(str); {
		subL := 0
		for j := i; j < len(str); j++ {
			if _, ok := acceptable[str[j]]; !ok {
				break
			}
			subL++
		}
		if max < subL {
			max = subL
		}
		// next of broken point
		i = i + subL + 1
	}
	fmt.Printf("%d\n", max)
}
