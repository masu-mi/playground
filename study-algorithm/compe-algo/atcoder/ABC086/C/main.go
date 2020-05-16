package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var (
		preT, preX, preY,
		nextT, nextX, nextY int
	)
	disconnected := false
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d\n", &nextT, &nextX, &nextY)
		distance := nextX - preX + nextY - preY
		restTime := nextT - preT
		if distance > restTime {
			disconnected = true
		}
		if (restTime-distance)%2 == 1 {
			disconnected = true
		}
		preT, preX, preY = nextT, nextX, nextY
	}
	if disconnected {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
