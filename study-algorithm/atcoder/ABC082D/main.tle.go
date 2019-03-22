package main

import (
	"fmt"
)

func main() {
	var (
		s    string
		x, y int
	)
	fmt.Scan(&s, &x, &y)
	start, xs, ys := parseOps(s)
	if fullSearchOnPlane(x-start, y, xs, ys) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func fullSearchOnPlane(goalX, goalY int, moveXs, moveYs map[int]int) bool {
	return fullSearchOnLine(goalX, moveXs) && fullSearchOnLine(goalY, moveYs)
}

func fullSearchOnLine(goal int, moves map[int]int) bool {
	selection := make([]int, len(moves))
	currentState := make(map[int]int)
	for move, num := range moves {
		selection = append(selection, move)
		currentState[move] = num
	}
	for true {
		// check finish
		allStopped := true
		for _, k := range selection {
			if currentState[k] >= -moves[k] {
				allStopped = false
			}
		}
		if check(goal, currentState) {
			return true
		}
		if allStopped {
			break
		}
		// progress state
		for _, k := range selection {
			currentState[k] -= 2
			if currentState[k] < -moves[k] {
				currentState[k] = moves[k]
			} else {
				break
			}
		}
	}
	return false
}

func check(goal int, current map[int]int) bool {
	sum := 0
	for k, v := range current {
		sum += k * v
	}
	return sum == goal
}

func parseOps(s string) (start int, x, y map[int]int) {
	for len(s) > 0 {
		if s[0] != 'F' {
			break
		}
		start++
		s = s[1:len(s)]
	}
	const (
		Y = -1
		X = 1
	)
	x, y = map[int]int{}, map[int]int{}
	direction := X
	continueLength := 0
	for len(s) > 0 {
		switch s[0] {
		case 'T':
			if direction == X {
				x[continueLength]++
			} else { // Y
				y[continueLength]++
			}
			continueLength = 0
			direction *= -1
		case 'F':
			continueLength++
		}
		s = s[1:len(s)]
	}
	if direction == X {
		x[continueLength]++
	} else { // Y
		y[continueLength]++
	}
	delete(x, 0)
	delete(y, 0)
	return
}
