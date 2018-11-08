package rbtree

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"testing"
)

type item struct {
	k int
}

func (i *item) Key() int {
	return i.k
}

func newNode(k int) *Node {
	return &Node{v: &item{k: k}}
}

func Test_Lookup(t *testing.T) {
	tree := &RBTree{
		root: &Node{
			v: &item{k: 10},
			l: &Node{
				v: &item{k: 5},
				l: &Node{v: &item{k: 2}},
				r: &Node{v: &item{k: 8}},
			},
			r: &Node{v: &item{k: 12}},
		},
	}
	type test struct {
		key  int
		item *item
		err  error
	}
	for _, c := range []test{
		test{key: 2, item: &item{k: 2}},
		test{key: 8, item: &item{k: 8}},
		test{key: 11, err: errors.New("not found")},
		test{key: 12, item: &item{k: 12}},
	} {
		act, actErr := tree.Lookup(c.key)
		if act == nil && c.item != nil {
			t.Errorf("unmatch c.item(%v) act: nil", c.item)
		} else if c.item != nil && act != nil {
			if c.item.k != act.(*item).Key() {
				t.Errorf("unmatch %v != %v", c.item.k, act.(*item).k)
			}
		}
		if expectedErr := c.err; expectedErr != nil && actErr != nil {
			if expectedErr.Error() != actErr.Error() {
				t.Errorf("unmatch %v != %v", expectedErr, actErr)
			}
		}
	}
}

func Test_Insert(t *testing.T) {
	type testCase struct {
		inserts  []int
		expected string
	}
	for _, c := range []testCase{
		{
			inserts: []int{1, 2, 3},
			expected: `node(k: 2, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 2, color: Color(BLACK))
  -[Right]-> node(k: 3, p: 2, color: Color(BLACK))
`,
		},
		{
			inserts: []int{3, 2, 1},
			expected: `node(k: 2, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 2, color: Color(BLACK))
  -[Right]-> node(k: 3, p: 2, color: Color(BLACK))
`,
		},
		{
			inserts: []int{3, 1, 2},
			expected: `node(k: 2, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 2, color: Color(BLACK))
  -[Right]-> node(k: 3, p: 2, color: Color(BLACK))
`,
		},
		{
			inserts: []int{1, 3, 2},
			expected: `node(k: 2, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 2, color: Color(BLACK))
  -[Right]-> node(k: 3, p: 2, color: Color(BLACK))
`,
		},
	} {
		tree := &RBTree{}
		for _, k := range c.inserts {
			tree.Insert(&item{k: k})
		}
		act := treeToString(tree)
		expected := c.expected
		if expected != act {
			t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, expected)
		}
	}
}

func sampleImputTree() *RBTree {
	root := newNode(5)
	cL := newNode(3)
	setLC(root, cL)
	gcLL := newNode(1)
	setLC(cL, gcLL)
	gcLR := newNode(4)
	setRC(cL, gcLR)
	cR := newNode(7)
	setRC(root, cR)
	gcRR := newNode(8)
	setRC(cR, gcRR)
	return &RBTree{root: root}
}

func Test_replace(t *testing.T) {
	type testCase struct {
		key      int
		expected string
	}

	for _, c := range []testCase{
		testCase{
			key: 1,
			expected: `node(k: 5, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 3, p: 5, color: Color(RED))
    -[Right]-> node(k: 4, p: 3, color: Color(RED))
  -[Right]-> node(k: 7, p: 5, color: Color(RED))
    -[Right]-> node(k: 8, p: 7, color: Color(RED))
`,
		},
		testCase{
			key: 4,
			expected: `node(k: 5, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 3, p: 5, color: Color(RED))
    -[Left ]-> node(k: 1, p: 3, color: Color(RED))
  -[Right]-> node(k: 7, p: 5, color: Color(RED))
    -[Right]-> node(k: 8, p: 7, color: Color(RED))
`,
		},
		testCase{
			key: 3,
			expected: `node(k: 5, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 1, p: 5, color: Color(RED))
    -[Right]-> node(k: 4, p: 1, color: Color(RED))
  -[Right]-> node(k: 7, p: 5, color: Color(RED))
    -[Right]-> node(k: 8, p: 7, color: Color(RED))
`,
		},
		testCase{
			key: 7,
			expected: `node(k: 5, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 3, p: 5, color: Color(RED))
    -[Left ]-> node(k: 1, p: 3, color: Color(RED))
    -[Right]-> node(k: 4, p: 3, color: Color(RED))
  -[Right]-> node(k: 8, p: 5, color: Color(RED))
`,
		},
		testCase{
			key: 5,
			expected: `node(k: 4, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 3, p: 4, color: Color(RED))
    -[Left ]-> node(k: 1, p: 3, color: Color(RED))
  -[Right]-> node(k: 7, p: 4, color: Color(RED))
    -[Right]-> node(k: 8, p: 7, color: Color(RED))
`,
		},
	} {
		tree := sampleImputTree()
		ref, p := tree.search(c.key)
		replace(ref, p)
		act := treeToString(tree)
		if act != c.expected {
			t.Errorf("Delete %d: [ACT]\n%v\n[EXPECTED]\n%v\n", c.key, act, c.expected)
		}
	}
}

func setLC(p *Node, c *Node) {
	p.l, c.p = c, p
}
func setRC(p *Node, c *Node) {
	p.r, c.p = c, p
}

func Test_recoverRank_root(t *testing.T) {
	tree := &RBTree{root: leafColorNode(Red)}
	recoverRank(nil, tree.root)
	if tree.root.color != Black {
		t.Errorf("recoverRank(root_node) set Black")
	}
}

func Test_recoverRank_left(t *testing.T) {
	type testCase struct {
		rightNode *Node
		expected  string
	}

	for idx, c := range []testCase{
		testCase{
			rightNode: nil,
			expected: `node(k: 1, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
`,
		},
		testCase{
			rightNode: colorNode(Black, nil, nil),
			expected: `node(k: 1, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
  -[Right]-> node(k: 1, p: 1, color: Color(RED))
`,
		},
		testCase{
			rightNode: colorNode(Black, leafColorNode(Black), leafColorNode(Black)),
			expected: `node(k: 1, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
  -[Right]-> node(k: 1, p: 1, color: Color(RED))
    -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Right]-> node(k: 1, p: 1, color: Color(BLACK))
`,
		},
		testCase{
			rightNode: colorNode(
				Black,
				colorNode(Red, leafColorNode(Red), leafColorNode(Red)),
				nil,
			),
			expected: `node(k: 1, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Right]-> node(k: 1, p: 1, color: Color(RED))
  -[Right]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(RED))
`,
		},
		testCase{
			rightNode: colorNode(
				Black,
				nil,
				colorNode(Red, leafColorNode(Red), leafColorNode(Red)),
			),
			expected: `node(k: 1, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
  -[Right]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(RED))
    -[Right]-> node(k: 1, p: 1, color: Color(RED))
`,
		},
		testCase{
			rightNode: colorNode(
				Red,
				colorNode(Black, leafColorNode(Black), nil),
				colorNode(Black, leafColorNode(Black), leafColorNode(Red)),
			),
			expected: `node(k: 1, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Right]-> node(k: 1, p: 1, color: Color(RED))
      -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
  -[Right]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Left ]-> node(k: 1, p: 1, color: Color(BLACK))
    -[Right]-> node(k: 1, p: 1, color: Color(RED))
`,
		},
	} {
		targetL := colorNode(Black, nil, nil)
		tree := &RBTree{root: colorNode(Red, targetL, c.rightNode)}
		recoverRank(tree.root, tree.root.l)
		act := treeToString(tree)
		if act != c.expected {
			t.Errorf("test_id: %d, [ACT]\n%v\n[EXPECTED]\n%v\n", idx, act, c.expected)
		}
	}

	// wn := colorNode(Black, leafColorNode(Red), leafColorNode(Black))
	// wn := colorNode(Black, nil, leafColorNode(Red))
	// wn := colorNode(Black, leafColorNode(Red), leafColorNode(Black))
	// wn := colorNode(Black, nil, nil)
	// wn := colorNode(Black, leafColorNode(Black), leafColorNode(Black))
}

func colorNode(c Color, left, right *Node) *Node {
	p := &Node{
		color: c,
		v:     &item{k: 1},
		l:     left,
		r:     right,
	}
	if left != nil {
		left.p = p
	}
	if right != nil {
		right.p = p
	}
	return p
}

func leafColorNode(c Color) *Node {
	return &Node{color: c, v: &item{k: 1}}
}

func Test_Delete(t *testing.T) {
	type testCase struct {
		delete   int
		expected string
	}
	inserts := []int{10, 20, 30, 90, 40, 50, 60, 55, 70, 80}
	for _, c := range []testCase{
		{
			delete: 10,
			expected: `node(k: 60, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 40, p: 60, color: Color(BLACK))
    -[Left ]-> node(k: 20, p: 40, color: Color(BLACK))
      -[Right]-> node(k: 30, p: 20, color: Color(RED))
    -[Right]-> node(k: 50, p: 40, color: Color(BLACK))
      -[Right]-> node(k: 55, p: 50, color: Color(RED))
  -[Right]-> node(k: 80, p: 60, color: Color(BLACK))
    -[Left ]-> node(k: 70, p: 80, color: Color(BLACK))
    -[Right]-> node(k: 90, p: 80, color: Color(BLACK))
`,
		},
		{
			delete: 55,
			expected: `node(k: 40, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 20, p: 40, color: Color(BLACK))
    -[Left ]-> node(k: 10, p: 20, color: Color(BLACK))
    -[Right]-> node(k: 30, p: 20, color: Color(BLACK))
  -[Right]-> node(k: 60, p: 40, color: Color(BLACK))
    -[Left ]-> node(k: 50, p: 60, color: Color(BLACK))
    -[Right]-> node(k: 80, p: 60, color: Color(RED))
      -[Left ]-> node(k: 70, p: 80, color: Color(BLACK))
      -[Right]-> node(k: 90, p: 80, color: Color(BLACK))
`,
		},
		{
			delete: 80,
			expected: `node(k: 40, p: <nil>, color: Color(BLACK))
  -[Left ]-> node(k: 20, p: 40, color: Color(BLACK))
    -[Left ]-> node(k: 10, p: 20, color: Color(BLACK))
    -[Right]-> node(k: 30, p: 20, color: Color(BLACK))
  -[Right]-> node(k: 60, p: 40, color: Color(BLACK))
    -[Left ]-> node(k: 50, p: 60, color: Color(BLACK))
      -[Right]-> node(k: 55, p: 50, color: Color(RED))
    -[Right]-> node(k: 70, p: 60, color: Color(BLACK))
      -[Right]-> node(k: 90, p: 70, color: Color(RED))
`,
		},
	} {
		tree := &RBTree{}
		for _, k := range inserts {
			tree.Insert(&item{k: k})
		}
		tree.Delete(c.delete)
		act := treeToString(tree)
		expected := c.expected
		if expected != act {
			t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, expected)
		}
	}
}

func nodeToString(node *Node) string {
	return treeToString(&RBTree{root: node})
}

func treeToString(tree *RBTree) string {
	w := bytes.NewBuffer([]byte{})
	var offset int
	tree.Traverse(func(e Event, node Node) {
		switch e {
		case Ent:
			fmt.Fprintf(w, "%v\n", &node)
			offset += 2
		case PreLeft:
			fmt.Fprintf(w, "%s-[Left ]-> ", strings.Repeat(" ", offset))
		case PreRight:
			fmt.Fprintf(w, "%s-[Right]-> ", strings.Repeat(" ", offset))
		case Exit:
			offset -= 2
		}
	})
	return w.String()
}
