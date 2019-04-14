package queue

import (
	"errors"
)

type queue struct {
	rIdx, wIdx int
	qs         [][]interface{}
}

func NewQueue() *queue {
	q := &queue{}
	q.qs = make([][]interface{}, 2)
	return q
}

func (q *queue) Push(item interface{}) {
	if q.rIdx == q.wIdx {
		q.wIdx = (q.wIdx + 1) % 2
	}
	q.qs[q.wIdx] = append(q.qs[q.wIdx], item)
}

func (q *queue) Empty() bool {
	return q.rIdx == q.wIdx && len(q.qs[q.rIdx]) == 0
}

func (q *queue) Pop() (item interface{}, e error) {
	if q.rIdx != q.wIdx && len(q.qs[q.rIdx]) == 0 {
		q.qs[q.rIdx] = nil
		q.rIdx = q.wIdx
	}
	if len(q.qs[q.rIdx]) == 0 {
		return nil, errors.New("empty")
	}
	item, q.qs[q.rIdx] = q.qs[q.rIdx][0], q.qs[q.rIdx][1:len(q.qs[q.rIdx])]
	return item, nil
}
