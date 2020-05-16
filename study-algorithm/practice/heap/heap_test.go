package heap

import "testing"

func Test_PushPop(t *testing.T) {

	type testCase struct {
		input    []int
		expected []int
	}

	for idx, test := range []testCase{
		testCase{
			input:    []int{1},
			expected: []int{1},
		},
		testCase{
			input:    []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		testCase{
			input:    []int{1, 3, 2},
			expected: []int{3, 2, 1},
		},
		testCase{
			input:    []int{3, 1, 2},
			expected: []int{3, 2, 1},
		},
		testCase{
			input:    []int{3, 2, 1},
			expected: []int{3, 2, 1},
		},
	} {
		h := NewHeap(len(test.input))
		for _, v := range test.input {
			h.Push(v)
		}
		for _, exp := range test.expected {
			result := h.Pop()
			if result != exp {
				t.Errorf("test_id: %d expected %d but returned %d", idx, exp, result)
			}
		}
	}
}
