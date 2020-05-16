package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	patterns, text := parseProblem()
	p := initPMA(patterns)
	for m := range p.search(text) {
		fmt.Printf("index:%d, pattern: %s\n", m.index, m.pattern)
	}
}

func parseProblem() (patterns []string, text string) {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)
	for i := 0; i < n; i++ {
		sc.Scan()
		pattern := sc.Text()
		patterns = append(patterns, pattern)
	}
	sc.Scan()
	text = sc.Text()
	return
}

var (
	ErrUnmatch = errors.New("unmatch")
	ErrEmpty   = errors.New("empty")
)

type Node int

type PMA struct {
	num     int
	nexts   map[Node]map[rune]Node
	failure map[Node]Node
	output  map[Node][]string
}

func newPMA() *PMA {
	return &PMA{
		nexts:   make(map[Node]map[rune]Node),
		failure: make(map[Node]Node),
		output:  make(map[Node][]string),
	}
}

func (p *PMA) setupGoto(patterns []string) {
	current := Node(0)
	for _, pattern := range patterns {
		for _, b := range pattern {
			if _, ok := p.nexts[current]; !ok {
				p.nexts[current] = make(map[rune]Node)
			}
			if _, ok := p.nexts[current][b]; !ok {
				p.num++
				next := Node(p.num)
				p.nexts[current][b] = next
			}
			current = p.nexts[current][b]
		}
		p.output[current] = append(p.output[current], pattern)
		current = Node(0)
	}
}

func (p *PMA) goTo(cur Node, b rune) (n Node, e error) {
	n, ok := p.nexts[cur][b]
	if !ok {
		return -1, ErrUnmatch
	}
	return n, nil
}

func (p *PMA) setupFailure() {
	root := Node(0)
	p.failure[root] = root
	q := &queue{}
	q.Enqueue(root)
	for q.Size() > 0 {
		cur, _ := q.Dequeue()
		for b, n := range p.nexts[cur] {
			q.Enqueue(n)
			_, ok := p.failure[cur]
			if !ok {
				p.failure[cur] = root
			}

			f := p.failure[cur]
			for true {
				if f == root {
					break
				}
				_, e := p.goTo(f, b)
				if e == nil {
					break
				}
				f = p.failure[f]
			}
			if cur == root {
				continue
			}
			fn, e := p.goTo(f, b)
			if e == nil {
				p.failure[n] = fn
				p.output[n] = append(p.output[n], p.output[fn]...)
			} else {
				p.failure[n] = root
				p.output[n] = append(p.output[n], p.output[root]...)
			}
		}
	}
}

type matchPoint struct {
	index   int
	pattern string
}

func (p *PMA) search(text string) chan matchPoint {
	resultCh := make(chan matchPoint)
	go func() {
		root := Node(0)
		cur := root
		idx := 0
		for _, b := range text {
			_, e := p.goTo(cur, b)
			if e != nil {
				cur = p.failure[cur]
			}
			cur, e = p.goTo(cur, b)
			if e != nil {
				cur = root
			}
			for _, result := range p.output[cur] {
				resultCh <- matchPoint{
					index:   idx - len(result) + 1,
					pattern: result,
				}
			}
			idx++
		}
		close(resultCh)
	}()
	return resultCh
}

type queue struct {
	wIdx, rIdx int
	queues     [2][]Node
}

func (q *queue) Size() int {
	return len(q.queues[0]) + len(q.queues[1])
}

func (q *queue) Enqueue(n Node) {
	if q.wIdx == q.rIdx {
		q.wIdx = (q.wIdx + 1) % 2
	}
	q.queues[q.wIdx] = append(q.queues[q.wIdx], n)
}
func (q *queue) Dequeue() (Node, error) {
	var v Node
	if len(q.queues[q.rIdx]) == 0 {
		q.queues[q.rIdx] = nil
		q.rIdx = q.wIdx
	}
	if len(q.queues[q.rIdx]) == 0 {
		return -1, ErrEmpty
	}
	v, q.queues[q.rIdx] = q.queues[q.rIdx][0], q.queues[q.rIdx][1:]
	return v, nil
}

func initPMA(patterns []string) *PMA {
	pma := newPMA()
	pma.setupGoto(patterns)
	pma.setupFailure()
	return pma
}
