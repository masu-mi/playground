package main

import "fmt"

func main() {
	c := make([][]int, 3)
	for i := 0; i < 3; i++ {
		tmp := make([]int, 3)
		fmt.Scan(&tmp[0], &tmp[1], &tmp[2])
		c[i] = tmp
	}
	if checkRow(diff(c, 1, 0)) && checkCol(diff(c, 0, 1)) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}

func checkRow(diff [][]int) bool {
	for i := 0; i < 3; i++ {
		row := diff[i]
		for j := 1; j < len(row); j++ {
			if diff[i][0] != diff[i][j] {
				return false
			}
		}
	}
	return true
}

func checkCol(diff [][]int) bool {
	for i := 0; i < 3; i++ {
		for j := 1; j < len(diff); j++ {
			if diff[0][i] != diff[j][i] {
				return false
			}
		}
	}
	return true
}

func diff(c [][]int, slideI, slicdJ int) [][]int {
	diffs := make([][]int, 3)
	for i := slideI; i < len(c); i++ {
		diffs[i] = make([]int, 3)
		for j := slicdJ; j < len(c[i]); j++ {
			diffs[i][j] = c[i][j] - c[i-slideI][j-slicdJ]
		}
	}
	return diffs
}
