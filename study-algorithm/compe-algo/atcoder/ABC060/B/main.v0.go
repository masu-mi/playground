package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	modulos := map[int]struct{}{}
	i := 1
	for true {
		mod := i * a % b
		if mod == c {
			fmt.Println("YES")
			return
		}
		if _, visited := modulos[mod]; visited {
			break
		}
		modulos[mod] = struct{}{}
		i++
	}
	fmt.Println("NO")
	return
}
