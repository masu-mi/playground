package deque

import "testing"

func Test_Secinario(t *testing.T) {
	const (
		front = 0 + iota
		back
	)
	type ops struct {
		side int
		v    int
	}
	type testCase struct {
		inputs  []ops
		outputs []int
	}
	for _, test := range []testCase{
		{
			inputs:  []ops{{back, 0}, {back, 1}, {back, 2}},
			outputs: []int{0, 1, 2},
		},
		{
			inputs:  []ops{{back, 0}, {front, 1}, {back, 2}},
			outputs: []int{1, 0, 2},
		},
		{
			inputs:  []ops{{front, 0}, {front, 1}, {front, 2}},
			outputs: []int{2, 1, 0},
		},
		{
			inputs:  []ops{{back, 0}, {front, 1}, {front, 2}},
			outputs: []int{2, 1, 0},
		},
		{
			inputs:  []ops{{front, 0}, {front, 1}, {back, 2}},
			outputs: []int{1, 0, 2},
		},
		{
			inputs:  []ops{{front, 0}, {back, 1}, {front, 2}},
			outputs: []int{2, 0, 1},
		},
	} {
		dq := NewDeque(10)
		for _, op := range test.inputs {
			if op.side == back {
				dq.PushBack(op.v)
			} else {
				dq.PushFront(op.v)
			}
		}
		for _, exp := range test.outputs {
			act := dq.Pop()
			if act != exp {
				t.Errorf("exp: %d; act: %d\n", exp, act)
			}
		}
	}
}
