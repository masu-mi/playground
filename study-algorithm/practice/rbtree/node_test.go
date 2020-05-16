package rbtree

import (
	"testing"
)

type key int

func (k key) CompareTo(o Key) int {
	return int(k) - int(o.(key))
}

func Test_EqualTo(t *testing.T) {
	type testCase struct {
		l, r   *Node
		expect bool
	}
	for _, c := range []testCase{
		testCase{
			l:      nil,
			r:      nil,
			expect: true,
		},
		testCase{
			l:      nil,
			r:      simpleNode(RED, 0),
			expect: false,
		},
		testCase{
			l:      simpleNode(RED, 0),
			r:      nil,
			expect: false,
		},
		testCase{
			l:      simpleNode(RED, 0),
			r:      simpleNode(RED, 0),
			expect: true,
		},
		testCase{
			l:      simpleNode(BLACK, 0),
			r:      simpleNode(BLACK, 0),
			expect: true,
		},
		testCase{
			l:      simpleNode(RED, 0),
			r:      simpleNode(BLACK, 0),
			expect: false,
		},
		testCase{
			l:      simpleNode(BLACK, 0),
			r:      simpleNode(BLACK, 1),
			expect: false,
		},
		testCase{
			l:      parentNode(BLACK, 0, simpleNode(RED, -1), nil),
			r:      simpleNode(BLACK, 0),
			expect: true,
		},
	} {
		if c.l.EqualTo(c.r) != c.expect {
			t.Errorf("FAILED[EqualTo()] (l: %s, r:%s, expect: %t)", c.l, c.r, c.expect)
		}
	}
}

func Test_EqualAsSubTree(t *testing.T) {
	type testCase struct {
		l, r   *Node
		expect bool
	}
	for _, c := range []testCase{
		testCase{
			l:      nil,
			r:      nil,
			expect: true,
		},
		testCase{
			l:      nil,
			r:      simpleNode(RED, 0),
			expect: false,
		},
		testCase{
			l:      simpleNode(RED, 0),
			r:      nil,
			expect: false,
		},
		testCase{
			l:      simpleNode(RED, 0),
			r:      simpleNode(RED, 0),
			expect: true,
		},
		testCase{
			l:      simpleNode(RED, 1),
			r:      simpleNode(RED, 0),
			expect: false,
		},
		testCase{
			l:      parentNode(RED, 0, nil, simpleNode(BLACK, 1)),
			r:      parentNode(RED, 0, nil, simpleNode(BLACK, 1)),
			expect: true,
		},
		testCase{
			l:      parentNode(BLACK, 0, simpleNode(BLACK, 1), nil),
			r:      parentNode(BLACK, 0, simpleNode(BLACK, 1), nil),
			expect: true,
		},
		testCase{
			l:      parentNode(BLACK, 0, simpleNode(BLACK, 1), nil),
			r:      parentNode(BLACK, 0, nil, nil),
			expect: false,
		},
		testCase{
			l:      parentNode(BLACK, 0, simpleNode(BLACK, 1), nil),
			r:      parentNode(BLACK, 0, simpleNode(RED, 1), nil),
			expect: false,
		},
	} {
		if c.l.EqualAsSubTree(c.r) != c.expect {
			t.Errorf("FAILED[EqualAsSubTree()] (l: %s, r:%s, expect: %t)", c.l, c.r, c.expect)
		}
	}
}

func Test_rotateL(t *testing.T) {
	type testCase struct {
		input    *Node
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			input:    parentNode(BLACK, 0, nil, simpleNode(RED, 10)),
			expected: parentNode(RED, 10, simpleNode(BLACK, 0), nil),
		},
		testCase{
			input: parentNode(
				BLACK, 0,
				simpleNode(RED, 10),
				parentNode(RED, 20,
					simpleNode(RED, 30),
					simpleNode(RED, 40),
				),
			),
			expected: parentNode(
				RED, 20,
				parentNode(BLACK, 0,
					simpleNode(RED, 10),
					simpleNode(RED, 30),
				),
				simpleNode(RED, 40),
			),
		},
	} {
		rotateL(test.input)
		act := test.input
		if valid, invalidNode := checkNoBrokenLink(act); !valid {
			t.Errorf("case: %d; tree broken!!(\nat %s \nof %s\n)", idx, invalidNode, act)
		}
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("case: %d; unexpected result!!(%s)", idx, act)
		}
	}
}

func Test_rotateR(t *testing.T) {
	type testCase struct {
		input    *Node
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			input:    parentNode(BLACK, 0, simpleNode(RED, 10), nil),
			expected: parentNode(RED, 10, nil, simpleNode(BLACK, 0)),
		},
		testCase{
			input: parentNode(
				BLACK, 0,
				parentNode(RED, 10,
					simpleNode(RED, 20),
					simpleNode(RED, 30),
				),
				simpleNode(RED, 40),
			),
			expected: parentNode(
				RED, 10,
				simpleNode(RED, 20),
				parentNode(BLACK, 0,
					simpleNode(RED, 30),
					simpleNode(RED, 40),
				),
			),
		},
	} {
		rotateR(test.input)
		act := test.input
		if valid, invalidNode := checkNoBrokenLink(act); !valid {
			t.Errorf("case: %d; tree broken!!(\nat %s \nof %s\n)", idx, invalidNode, act)
		}
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("case: %d; unexpected result!!(%s)", idx, act)
		}
	}
}

func Test_rotateLR(t *testing.T) {
	type testCase struct {
		input    *Node
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			input: parentNode(
				BLACK, 0,
				parentNode(RED, 10, nil, simpleNode(RED, 20)),
				nil,
			),
			expected: parentNode(
				RED, 20,
				simpleNode(RED, 10),
				simpleNode(BLACK, 0),
			),
		},
		testCase{
			input: parentNode(
				BLACK, 0,
				parentNode(RED, 10,
					simpleNode(RED, 20),
					simpleNode(RED, 30),
				),
				simpleNode(RED, 40),
			),
			expected: parentNode(
				RED, 30,
				parentNode(RED, 10,
					simpleNode(RED, 20),
					nil,
				),
				parentNode(BLACK, 0,
					nil,
					simpleNode(RED, 40),
				),
			),
		},
	} {
		rotateLR(test.input)
		act := test.input
		if valid, invalidNode := checkNoBrokenLink(act); !valid {
			t.Errorf("case: %d; tree broken!!(\nat %s \nof %s\n)", idx, invalidNode, act)
		}
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("case: %d; unexpected result!!(%s)", idx, act)
		}
	}
}

func Test_rotateRL(t *testing.T) {
	type testCase struct {
		input    *Node
		expected *Node
	}
	for idx, test := range []testCase{
		testCase{
			input: parentNode(
				BLACK, 0,
				nil,
				parentNode(RED, 10, simpleNode(RED, 20), nil),
			),
			expected: parentNode(
				RED, 20,
				simpleNode(BLACK, 0),
				simpleNode(RED, 10),
			),
		},
		testCase{
			input: parentNode(
				BLACK, 0,
				simpleNode(RED, 10),
				parentNode(RED, 20,
					simpleNode(RED, 30),
					simpleNode(RED, 40),
				),
			),
			expected: parentNode(
				RED, 30,
				parentNode(BLACK, 0,
					simpleNode(RED, 10),
					nil,
				),
				parentNode(RED, 20,
					nil,
					simpleNode(RED, 40),
				),
			),
		},
	} {
		rotateRL(test.input)
		act := test.input
		if valid, invalidNode := checkNoBrokenLink(act); !valid {
			t.Errorf("case: %d; tree broken!!(\nat %s \nof %s\n)", idx, invalidNode, act)
		}
		if !act.EqualAsSubTree(test.expected) {
			t.Errorf("case: %d; unexpected result!!(%s)", idx, act)
		}
	}
}

func Test_isLeftChild(t *testing.T) {
	type testCase struct {
		node     *Node
		expected bool
	}

	for idx, test := range []testCase{
		testCase{
			node:     rootNode(RED, 0, nil, nil),
			expected: false,
		},
		testCase{
			node: func() *Node {
				p := parentNode(RED, 0, simpleNode(BLACK, 1), nil)
				return p.l
			}(),
			expected: true,
		},
		testCase{
			node: func() *Node {
				p := parentNode(RED, 0, nil, simpleNode(BLACK, 1))
				return p.r
			}(),
			expected: false,
		},
	} {
		if test.node.isLeftChild() != test.expected {
			t.Errorf("case: %d; unexpected result!!", idx)
		}
	}
}
