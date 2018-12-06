package rbtree

import "testing"

type Event int

const (
	Entry Event = 0 + iota
	PreLeftEntry
	PostLeftEntry
	PreRightEntry
	PostRightEntry
	Exit
)

func traverse(n *Node, f func(e Event, n *Node)) {
	if n == nil {
		return
	}
	f(Entry, n)
	if n.l != nil {
		f(PreLeftEntry, n)
		traverse(n.l, f)
		f(PostLeftEntry, n)
	}
	if n.r != nil {
		f(PreRightEntry, n)
		traverse(n.r, f)
		f(PostRightEntry, n)
	}
	f(Exit, n)
}

func ranks(n *Node) chan int {
	currentRank := 0
	ch := make(chan int)
	go func() {
		traverse(n, func(e Event, n *Node) {
			switch e {
			case Entry:
				if n.color == BLACK {
					currentRank += 1
				}
				if n.l == nil {
					ch <- currentRank
				}
				if n.r == nil {
					ch <- currentRank
				}
			case Exit:
				if n.color == BLACK {
					currentRank -= 1
				}
			}
		})
		close(ch)
	}()
	return ch
}

func allRanksSame(n *Node) bool {
	if n == nil {
		return true
	}
	rs := map[int]struct{}{}
	for r := range ranks(n) {
		rs[r] = struct{}{}
	}
	return len(rs) == 1
}

func checkNoBrokenLink(n *Node) (bool, *Node) {
	type result struct {
		valid bool
		node  *Node
	}
	ch := make(chan result)
	go func() {
		traverse(n, func(e Event, n *Node) {
			switch e {
			case Entry:
				if n.p != n && n.p != nil {
					ch <- result{
						n.p.l == n || n.p.r == n,
						n,
					}
				}
				if n.l != nil {
					ch <- result{
						n.l.p == n,
						n,
					}
				}
				if n.r != nil {
					ch <- result{
						n.r.p == n,
						n,
					}
				}
			}
		})
		close(ch)
	}()
	for r := range ch {
		if !r.valid {
			return false, r.node
		}
	}
	return true, nil
}

func Test_allRankSame(t *testing.T) {
	for _, node := range []*Node{
		nil,
		simpleNode(BLACK, 0),
		parentNode(
			BLACK, 0,
			simpleNode(BLACK, 1),
			simpleNode(BLACK, 2),
		),
		parentNode(
			BLACK, 0,
			simpleNode(BLACK, 1),
			parentNode(
				BLACK, 2,
				simpleNode(RED, 3),
				simpleNode(RED, 4),
			),
		),
		parentNode(
			BLACK, 0,
			simpleNode(BLACK, 1),
			parentNode(
				RED, 2,
				simpleNode(BLACK, 3),
				simpleNode(BLACK, 4),
			),
		),
	} {
		if !allRanksSame(node) {
			t.Errorf("FAILED[allRanksSame()] (node: %s)", node)
		}
	}
}

func Test_checkNoBrokenLink(t *testing.T) {
	type testCase struct {
		node         *Node
		expected     bool
		expectedNode *Node
	}
	for idx, tc := range []testCase{
		testCase{
			node:         nil,
			expected:     true,
			expectedNode: nil,
		},
		testCase{
			node: &Node{
				k: key(1),
				l: &Node{},
			},
			expected:     false,
			expectedNode: simpleNode(RED, 1),
		},
		testCase{
			node:         parentNode(BLACK, 0, nil, simpleNode(RED, 1)),
			expected:     true,
			expectedNode: nil,
		},
	} {
		act, actErrNode := checkNoBrokenLink(tc.node)
		if tc.expected != act {
			t.Errorf(
				"FAILED[checkNoBrokenLink()] (no: %d, act: %t)",
				idx,
				act,
			)
		}
		if !tc.expectedNode.EqualTo(actErrNode) {
			t.Errorf(
				"FAILED[checkNoBrokenLink()] (no: %d, actNode: %s)",
				idx,
				actErrNode,
			)
		}
	}
}

func simpleNode(c Color, k int) *Node {
	return &Node{color: c, k: key(k)}
}
func parentNode(c Color, k int, l, r *Node) *Node {
	p := &Node{color: c, k: key(k), l: l, r: r}
	if l != nil {
		l.p = p
	}
	if r != nil {
		r.p = p
	}
	return p
}
func rootNode(c Color, k int, l, r *Node) *Node {
	p := parentNode(c, k, l, r)
	p.p = p
	return p
}

func assertNode(t *testing.T, name string, expected, act *Node) {
	if !act.EqualTo(expected) {
		t.Errorf("invalid %s node was found(%v)!!", name, act)
	}
}
