package rbtree

import "testing"

var rotatedResult = `node(k: 2, p: <nil>, color: Color(RED))
  -[Left ]-> node(k: 1, p: 2, color: Color(RED))
  -[Right]-> node(k: 3, p: 2, color: Color(RED))
`

func Test_rotateR(t *testing.T) {
	pp := &Node{v: &item{k: 3}}
	p := &Node{v: &item{k: 2}}
	n := &Node{v: &item{k: 1}}

	tree := &RBTree{root: pp}
	pp.l = p
	p.l, p.p = n, pp
	n.p = p

	rotateR(pp)
	act := treeToString(tree)
	if rotatedResult != act {
		t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, rotatedResult)
	}
}

func Test_rotateRL(t *testing.T) {
	pp := &Node{v: &item{k: 1}}
	p := &Node{v: &item{k: 3}}
	n := &Node{v: &item{k: 2}}

	tree := &RBTree{root: pp}
	pp.r = p
	p.l, p.p = n, pp
	n.p = p

	rotateRL(pp)
	act := treeToString(tree)
	if rotatedResult != act {
		t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, rotatedResult)
	}
}

func Test_rotateLR(t *testing.T) {
	pp := &Node{v: &item{k: 3}}
	p := &Node{v: &item{k: 1}}
	n := &Node{v: &item{k: 2}}

	tree := &RBTree{root: pp}
	pp.l = p
	p.r, p.p = n, pp
	n.p = p

	rotateLR(pp)
	act := treeToString(tree)
	if rotatedResult != act {
		t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, rotatedResult)
	}
}

func Test_rotateL(t *testing.T) {
	pp := &Node{v: &item{k: 1}}
	p := &Node{v: &item{k: 2}}
	n := &Node{v: &item{k: 3}}

	tree := &RBTree{root: pp}
	pp.r = p
	p.r, p.p = n, pp
	n.p = p

	rotateL(pp)
	act := treeToString(tree)
	if rotatedResult != act {
		t.Errorf("[ACT]\n%v\n[EXPECTED]\n%v\n", act, rotatedResult)
	}
}
