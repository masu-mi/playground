package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	yy := y / 1000
	r := search(n, yy)
	fmt.Printf("%d %d %d\n", r.yukichi, r.higuchi, r.natsume)
}

type combi struct{ yukichi, higuchi, natsume int }

func search(n, yy int) combi {
	for i := 0; i <= yy/10; i++ {
		for j := 0; j <= (yy-10*i)/5; j++ {
			if r := n - i - j; r == yy-10*i-5*j {
				return combi{i, j, r}
			}
		}
	}
	return combi{-1, -1, -1}
}
