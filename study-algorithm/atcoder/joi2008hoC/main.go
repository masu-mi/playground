package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	areas := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		areas[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.IntSlice(areas))

	twosScores := map[int]struct{}{}
	for i := 0; i < n; i++ {
		v := areas[i]
		if v > m {
			break
		}
		twosScores[v] = struct{}{}
	}
	for i := 0; i < n; i++ {
		vi := areas[i]
		if vi > m {
			break
		}
		for j := 0; j < n; j++ {
			vj := areas[j]
			if vi+vj > m {
				break
			}
			twosScores[vi+vj] = struct{}{}
		}
	}
	twoScores := make([]int, n*n)
	for v := range twosScores {
		twoScores = append(twoScores, v)
	}
	max := 0
	sort.Sort(sort.Reverse(sort.IntSlice(twoScores)))
	for _, fv := range twoScores {
		if max < fv {
			max = fv
		}
		idx := sort.Search(len(twoScores), func(i int) bool { return twoScores[i] <= m-fv })
		if idx < len(twoScores) {
			v := fv + twoScores[idx]
			if max < v {
				max = v
			}
		}
	}
	fmt.Printf("%d\n", max)
}
