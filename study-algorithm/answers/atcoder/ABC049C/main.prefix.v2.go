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
	for len(input) > 0 {
		switch {
		case match("eraser", input):
			input = consume("eraser", input)
		case match("erase", input):
			input = consume("erase", input)
		case match("dream", input):
			candidate := consume("dream", input)
			switch {
			case candidate == "", match("dream", candidate), match("eraser", candidate), match("erase", candidate):
				input = candidate
			case match("dreamer", input):
				input = consume("dreamer", input)
			default:
				return false
			}
		default:
			return false
		}
	}
	return true
}

func consume(fixed, input string) string {
	return input[len(fixed):len(input)]
}

func match(key, input string) bool {
	return len(input) >= len(key) &&
		input[0:len(key)] == key
}
