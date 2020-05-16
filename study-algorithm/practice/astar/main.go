package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	p := parseGridProblem(os.Stdin)
	for _, pos := range findPath(p) {
		fmt.Printf("(%d, %d)\n", pos.i, pos.j)
	}
}

func findPath(p *problem) []pos {
	visited := make([][]bool, p.h+2)
	for i := 0; i < p.h+2; i++ {
		visited[i] = make([]bool, p.w+2)
	}
	parents := make([][]pos, p.h+2)
	for i := 0; i < p.h+2; i++ {
		parents[i] = make([]pos, p.w+2)
	}

	start := pos{i: 1, j: 1}
	goal := pos{p.h, p.w}
	visited[start.i][start.j] = true
	cs := candidateHeap([]candidate{})
	heap.Push(
		&cs,
		candidate{
			cost:  0,
			score: score(start, goal, 0),
			pos:   start,
		},
	)
	for (&cs).Len() > 0 {
		cur := heap.Pop(&cs).(candidate)
		if cur.pos == goal {
			break
		}
		for _, d := range []pos{
			{i: 0, j: 1},
			{i: 0, j: -1},
			{i: -1, j: 0},
			{i: 1, j: 0},
		} {
			next := pos{i: cur.i + d.i, j: cur.j + d.j}
			if p.field[next.i][next.j] != '#' && !visited[next.i][next.j] {
				heap.Push(&cs, candidate{
					cost:  cur.cost + movingCost(p.field[next.i][next.j]),
					score: score(next, goal, cur.cost),
					pos:   next,
				})
				visited[next.i][next.j] = true
				parents[next.i][next.j] = cur.pos
			}
		}
	}
	cPos := goal
	result := []pos{cPos}
	for cPos != start {
		cPos = parents[cPos.i][cPos.j]
		result = append(result, cPos)
	}
	return result
}

func movingCost(f byte) int {
	switch f {
	case 'w':
		return 2
	case 'P':
		return 4
	case '.':
		return 1
	}
	os.Exit(1)
	return 100
}

func score(p, t pos, cost int) int {
	return t.i - p.i + t.j - p.j + cost
}

type candidate struct {
	cost, score int
	pos
}

type candidateHeap []candidate

func (h *candidateHeap) Len() int {
	return len(*h)
}
func (h *candidateHeap) Less(i, j int) bool {
	return (*h)[i].score < (*h)[j].score || ((*h)[i].score < (*h)[j].score && (*h)[i].cost < (*h)[j].cost)
}
func (h *candidateHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *candidateHeap) Push(x interface{}) {
	*h = append(*h, x.(candidate))
}
func (h *candidateHeap) Pop() (i interface{}) {
	i, *h = (*h)[len(*h)-1], (*h)[0:len(*h)-1]
	return i
}

type pos struct{ i, j int }
type problem struct {
	h, w  int
	field []string
}

func parseGridProblem(r io.Reader) *problem {
	p := &problem{}
	fmt.Scan(&p.h, &p.w)
	p.field = append(p.field, makeWall(p.w))
	sc := bufio.NewScanner(r)
	sc.Split(bufio.ScanLines)
	for i := 0; i < p.h; i++ {
		sc.Scan()
		p.field = append(p.field, "#"+sc.Text()+"#")
	}
	p.field = append(p.field, makeWall(p.w))
	return p
}

func makeWall(w int) string {
	return strings.Repeat("#", w+2)
}
