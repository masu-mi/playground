package main

import (
	"fmt"
)

type combi struct{ yukichi, higuchi, natsume int }

func main() {
	var n, y int
	fmt.Scanf("%d %d", &n, &y)
	y /= 1000
	c, ok := search(n, y)
	if !ok {
		fmt.Printf("-1 -1 -1")
		return
	}
	fmt.Printf("%d %d %d", c.yukichi, c.higuchi, c.natsume)
	return
}

func search(num, yen int) (combi, bool) {
	combinations := make([]map[int]combi, 10)
	combinations[0] = map[int]combi{0: combi{}}
	for y := 1; y <= yen; y++ {
		tmp := map[int]combi{}
		if y >= 5 {
			for k, v := range combinations[(y-5)%10] {
				if k < num {
					v.higuchi += 1
					tmp[k+1] = v
				}
			}
		}
		if y >= 10 {
			for k, v := range combinations[(y-10)%10] {
				if k < num {
					v.yukichi += 1
					tmp[k+1] = v
				}
			}
		}
		for k, v := range combinations[(y-1)%10] {
			if k < num {
				v.natsume += 1
				tmp[k+1] = v
			}
		}
		combinations[y%10] = tmp
	}
	c, ok := combinations[yen%10][num]
	return c, ok
}
