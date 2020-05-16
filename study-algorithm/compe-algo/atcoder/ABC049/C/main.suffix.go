package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s", &input)
	if generated(input) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func generated(input string) bool {
MATCHING:
	for len(input) > 0 {
		for _, p := range []string{"eraser", "erase", "dream", "dreamer"} {
			if match(p, input) {
				input = consume(p, input)
				continue MATCHING
			}
		}
		return false
	}
	return true
}

func consume(fixed, input string) string {
	return input[0 : len(input)-len(fixed)]
}

func match(key, input string) bool {
	return len(input) >= len(key) &&
		input[len(input)-len(key):len(input)] == key
}
