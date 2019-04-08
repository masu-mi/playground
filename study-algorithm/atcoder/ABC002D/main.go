package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	edgeds := map[int]map[int]struct{}{}
	for i := 1; i <= n; i++ {
		edgeds[i] = map[int]struct{}{}
	}
	ranks := make([]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		edgeds[x][y] = struct{}{}
		edgeds[y][x] = struct{}{}
		ranks[x-1]++
		ranks[y-1]++
	}
	// for cutting branch
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))
	limitSize := 1
	for i := 0; i < n; i++ {
		if ranks[i] < i {
			break
		}
		limitSize = i + 1
	}
	// generate k-subset and check it's clique.
	for i := limitSize; i > 0; i-- {
		if searchClique(i, 0, nil, edgeds) {
			fmt.Printf("%d\n", i)
			return
		}
	}
}

func searchClique(size, min int, nodes []int, edgeds map[int]map[int]struct{}) bool {
	n := len(edgeds)
	if min == n+1 || size == 0 {
		return size == 0
	}
	for i := min; i <= (n+1)-size; i++ {
		if searchClique(size-1, i+1, append(nodes, i), edgeds) {
			allConnected := true
			for _, j := range nodes {
				if _, ok := edgeds[i][j]; !ok {
					allConnected = false
					break
				}
			}
			if allConnected {
				return true
			}
		}
		if searchClique(size, i+1, nodes, edgeds) {
			return true
		}
	}
	return false
}
