package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	edges := map[int]map[int]struct{}{}
	for i := 1; i <= n; i++ {
		edges[i] = map[int]struct{}{}
	}
	ranks := make([]int, n)
	for i := 0; i < m; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		edges[x][y] = struct{}{}
		edges[y][x] = struct{}{}
		// for search upper limit of clique set
		ranks[x-1]++
		ranks[y-1]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ranks)))
	var limitSize int
	for i := 0; i < n; i++ {
		if ranks[i] < i {
			break
		}
		limitSize = i + 1
	}
	// generate k-subset and check it's clique.
	for k := limitSize; k > 0; k-- {
		if searchKClique(k, 0, n, nil, edges) {
			fmt.Printf("%d\n", k)
			return
		}
	}
}

func searchKClique(k, idx, last int, nodes []int, edges map[int]map[int]struct{}) bool {
	if idx == last+1 || k == 0 {
		return k == 0
	}
	for i := idx; i <= last-k+1; i++ {
		if searchKClique(k-1, i+1, last, append(nodes, i), edges) {
			allConnected := true
			for _, j := range nodes {
				if _, ok := edges[i][j]; !ok {
					allConnected = false
					break
				}
			}
			if allConnected {
				return true
			}
		}
		if searchKClique(k, i+1, last, nodes, edges) {
			return true
		}
	}
	return false
}
