package main

import "fmt"

func main() {
	var topping string
	fmt.Scan(&topping)
	sum := 700
	for i := 0; i < len(topping); i++ {
		if topping[i] == 'o' {
			sum += 100
		}
	}
	fmt.Printf("%d\n", sum)
}
