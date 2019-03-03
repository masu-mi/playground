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
		case eraser(input):
			input = consume("eraser", input)
		case erase(input):
			input = consume("erase", input)
		case dream(input):
			candidate := consume("dream", input)
			switch {
			case candidate == "", dream(candidate), eraser(candidate), erase(candidate):
				input = candidate
			case dreamer(input):
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

func dream(input string) bool {
	return len(input) >= len("dream") &&
		input[0:len("dream")] == "dream"
}
func dreamer(input string) bool {
	return len(input) >= len("dreamer") &&
		input[0:len("dreamer")] == "dreamer"
}
func erase(input string) bool {
	return len(input) >= len("erase") &&
		input[0:len("erase")] == "erase"
}
func eraser(input string) bool {
	return len(input) >= len("eraser") &&
		input[0:len("eraser")] == "eraser"
}
