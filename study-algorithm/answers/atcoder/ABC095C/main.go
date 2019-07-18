package main

import "fmt"

func main() {
	var a, b, c, x, y, sum int
	fmt.Scan(&a, &b, &c, &x, &y)
	if a+b < 2*c {
		sum = a*x + b*y
	} else if x <= y {
		if 2*c <= b {
			sum = 2 * c * y
		} else {
			sum = 2*c*x + (y-x)*b
		}
	} else {
		if 2*c <= a {
			sum = 2 * c * x
		} else {
			sum = 2*c*y + (x-y)*a
		}
	}
	fmt.Printf("%d\n", sum)
}
