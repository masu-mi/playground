package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("%d\n", resolve(parseProblem(os.Stdin)))
}

func parseProblem(r io.Reader) int {
	const (
		initialBufSize = 100000
		maxBufSize     = 1000000
	)
	buf := make([]byte, initialBufSize)

	sc := bufio.NewScanner(r)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords) // bufio.ScanLines

	return scanInt(sc)
}

func resolve(n int) int {
	states := make([]int, n+1)
	for i := 1; i <= n; i++ {
		states[i] = inf
	}
	states[0] = 0

	q := &candidates{}
	heap.Init(q)
	heap.Push(q, node{0, 0})
	for q.Len() > 0 {
		cur := heap.Pop(q).(node)
		if cur.cost > states[cur.id] {
			continue
		}
		for i := 1; cur.id+i <= n; i *= 6 {
			if changeToMin(&(states[cur.id+i]), cur.cost+1) {
				heap.Push(q, node{cur.id + i, cur.cost + 1})
			}
		}
		for i := 1; cur.id+i <= n; i *= 9 {
			if changeToMin(&(states[cur.id+i]), cur.cost+1) {
				heap.Push(q, node{cur.id + i, cur.cost + 1})
			}
		}
	}
	return states[n]
}

const (
	inf = math.MaxInt32
)

func changeToMin(v *int, cand int) (updated bool) {
	if *v > cand {
		*v = cand
		updated = true
	}
	return updated
}

func changeToMax(v *int, cand int) (updated bool) {
	if *v < cand {
		*v = cand
		updated = true
	}
	return updated
}

type candidates []node

type node struct {
	id, cost int
}

func (c *candidates) Len() int {
	return len(*c)
}

func (c *candidates) Less(i, j int) bool {
	nodes := *c
	return nodes[i].cost < nodes[j].cost
}

func (c *candidates) Swap(i, j int) {
	nodes := *c
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func (c *candidates) Push(x interface{}) {
	*c = append(*c, x.(node))
}

func (c *candidates) Pop() interface{} {
	nodes := *c
	l := nodes[len(nodes)-1]
	*c = nodes[0 : len(nodes)-1]
	return l
}

func (c *candidates) top() node {
	return (*c)[c.Len()-1]
}

// snip-scan-funcs
func scanInt(sc *bufio.Scanner) int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return int(i)
}
