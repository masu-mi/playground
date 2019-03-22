package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	start, xs, ys := parseOperations(s)
	var x, y int
	fmt.Scan(&x, &y)
	if isReachablePoint(x-start, y, xs, ys) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func isReachablePoint(x, y int, mxs, mys []int) bool {
	return isReachable(x, mxs) && isReachable(y, mys)
}

// const max = 8000

func isReachable(goal int, diff []int) bool {
	possibilities := make([]map[int]struct{}, 2)
	possibilities[0] = map[int]struct{}{0: struct{}{}}
	for i := 1; i < len(diff)+1; i++ {
		possibilities[i%2] = map[int]struct{}{}
		for k := range possibilities[(i-1)%2] {
			pos := k + diff[i-1]
			possibilities[i%2][pos] = struct{}{}
			pos = k - diff[i-1]
			possibilities[i%2][pos] = struct{}{}
		}
	}
	_, ok := possibilities[len(diff)%2][goal]
	return ok
}

func parseOperations(s string) (start int, xs, ys []int) {
	progress := 0
	for len(s) > 0 {
		if s[0] == 'T' {
			start = progress
			progress = 0
			s = s[1:len(s)]
			break
		}
		progress++
		s = s[1:len(s)]
	}
	if len(s) == 0 {
		start = progress
		return
	}
	mode := 0
	for len(s) > 0 {
		if s[0] == 'T' {
			if mode == 0 {
				if progress > 0 {
					ys = append(ys, progress)
				}
			} else {
				if progress > 0 {
					xs = append(xs, progress)
				}
			}
			mode ^= 1
			progress = 0
		} else {
			progress++
		}
		s = s[1:len(s)]
	}
	if mode == 0 {
		if progress > 0 {
			ys = append(ys, progress)
		}
	} else {
		if progress > 0 {
			xs = append(xs, progress)
		}
	}
	return
}
