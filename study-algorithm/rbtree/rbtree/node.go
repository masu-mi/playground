package rbtree

import "fmt"

type Color int

const (
	RED = 0 + iota
	BLACK
)

type Node struct {
	k       Key
	value   Value
	color   Color
	p, l, r *Node
}

func (n *Node) EqualTo(o *Node) bool {
	if n == nil || o == nil {
		return n == o
	}
	return n.color == o.color && n.k.CompareTo(o.k) == 0
}

func (n *Node) EqualAsSubTree(o *Node) bool {
	if n == nil {
		return n == o
	}
	return n.EqualTo(o) &&
		n.l.EqualAsSubTree(o.l) &&
		n.r.EqualAsSubTree(o.r)
}

func (n *Node) copyFrom(o *Node) {
	n.color = o.color
	n.k = o.k
	n.value = o.value
	n.l, n.r = o.l, o.r
}

func (n *Node) Color() Color {
	if n == nil {
		return BLACK
	}
	return n.color
}

func (n *Node) String() string {
	if n == nil {
		return "<nil>"
	}
	return fmt.Sprintf(
		"<Node(c: %d, k: %v); l: %s, r: %s>",
		n.color,
		n.k,
		n.l,
		n.r,
	)
}

func (n *Node) isLeftChild() bool {
	if n == nil {
		return false
	}
	return n.p.l == n
}
func (n *Node) isRoot() bool {
	if n == nil {
		return false
	}
	return n.p == n
}

func findRoot(n *Node) *Node {
	cur := n
	for cur.p != cur {
		cur = cur.p
	}
	return cur
}

func rotateL(n *Node) *Node {
	p := &Node{
		color: n.r.color,
		k:     n.r.k,
		value: n.r.value,
		l: &Node{
			color: n.color,
			k:     n.k,
			value: n.value,
			l:     n.l,
			r:     n.r.l,
		},
		r: n.r.r,
	}
	n.copyFrom(p)

	n.l.p = n
	if n.r != nil {
		n.r.p = n
	}
	if n.l.l != nil {
		n.l.l.p = n.l
	}
	if p.l.r != nil {
		n.l.r.p = n.l
	}
	return n
}
func rotateR(n *Node) *Node {
	p := &Node{
		color: n.l.color,
		k:     n.l.k,
		value: n.l.value,
		l:     n.l.l,
		r: &Node{
			color: n.color,
			k:     n.k,
			value: n.value,
			l:     n.l.r,
			r:     n.r,
		},
	}
	n.copyFrom(p)

	n.r.p = n
	if n.l != nil {
		n.l.p = n
	}
	if n.r.l != nil {
		n.r.l.p = n.r
	}
	if n.r.r != nil {
		n.r.r.p = n.r
	}
	return n
}
func rotateLR(n *Node) *Node {
	p := rotateL(n.l)
	n.l, p.p = p, n
	return rotateR(n)
}

func rotateRL(n *Node) *Node {
	p := rotateR(n.r)
	n.r, p.p = p, n
	return rotateL(n)
}
