package main

import (
	"testing"
)

func Test_findMaxCalorie(t *testing.T) {
	type testCase struct {
		c        int
		inputX   []int
		inputV   []int
		expected int
	}
	for _, test := range []testCase{
		{
			c:        20,
			inputX:   []int{2, 9, 16},
			inputV:   []int{80, 120, 1},
			expected: 191,
		},
		{
			c:        20,
			inputX:   []int{2, 9, 16},
			inputV:   []int{80, 0, 120},
			expected: 192,
		},
		{
			c:        100000000000000,
			inputX:   []int{50000000000000},
			inputV:   []int{1},
			expected: 0,
		},
	} {
		act := findMaxCalorie(test.c, test.inputX, test.inputV)
		if act != test.expected {
			t.Errorf("[unexpedted max calorie]: expected: %d; act: %d", test.expected, act)
		}
	}
}

func Test_leftSafePoints(t *testing.T) {
	type testCase struct {
		inputC         int
		inputX         []int
		inputV         []int
		expectedSingle []gain
		expectedDouble []gain
	}
	for _, test := range []testCase{
		{
			inputC: 20,
			inputX: []int{2, 9, 16},
			inputV: []int{80, 120, 1},

			expectedSingle: []gain{{2, 78}, {9, 113}},
			expectedDouble: []gain{{2, 76}, {9, 106}},
		},
	} {
		singleIdx, doubleIndex := findLeftSafePoints(test.inputX, test.inputV)
		if !gainListEqual(singleIdx, test.expectedSingle) {
			t.Errorf("[unexpedted singleIdx]: expected: %v; act: %v", test.expectedSingle, singleIdx)
		}
		if !gainListEqual(doubleIndex, test.expectedDouble) {
			t.Errorf("[unexpedted doubleIndex]: expected: %v; act: %v", test.expectedDouble, doubleIndex)
		}
	}
}

func Test_rightSafePoints(t *testing.T) {
	type testCase struct {
		inputC         int
		inputX         []int
		inputV         []int
		expectedSingle []gain
		expectedDouble []gain
	}
	for _, test := range []testCase{
		{
			inputC: 20,
			inputX: []int{2, 9, 16},
			inputV: []int{80, 120, 1},

			expectedSingle: []gain{{9, 110}, {2, 73}},
			expectedDouble: []gain{{9, 99}, {2, 66}},
		},
	} {
		singleIdx, doubleIndex := findRightSafePoints(test.inputC, test.inputX, test.inputV)
		if !gainListEqual(singleIdx, test.expectedSingle) {
			t.Errorf("[unexpedted singleIdx]: expected: %v; act: %v", test.expectedSingle, singleIdx)
		}
		if !gainListEqual(doubleIndex, test.expectedDouble) {
			t.Errorf("[unexpedted doubleIndex]: expected: %v; act: %v", test.expectedDouble, doubleIndex)
		}
	}
}

func gainListEqual(x, y []gain) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if xx, yy := x[i], y[i]; xx.pos != yy.pos || xx.totalGain != yy.totalGain {
			return false
		}
	}
	return true
}
