package queue

import "testing"

func Test_SimpleScinario(t *testing.T) {
	q := NewQueue()
	seq := []int{1, 10, 100, 1000, 10000, 100000}
	for _, v := range seq {
		q.Push(v)
	}
	for _, exp := range seq {
		act, _ := q.Pop()
		if exp != act {
			t.Errorf("act: %d; exp: %d\n", act, exp)
		}
	}
}
