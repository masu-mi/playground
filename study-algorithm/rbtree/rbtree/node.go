package rbtree

import (
	"errors"
	"fmt"
)

type Node struct {
	p, l, r *Node
	color   Color
	v       Item
}

type Color int

const (
	Red Color = 0 + iota
	Black
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Color(BLACK)"
	case Red:
		return "Color(RED)"
	}
	return "Color(error)"
}

func (n *Node) String() string {
	if n.p != nil {
		return fmt.Sprintf("node(k: %d, p: %v, color: %v)", n.v.Key(), n.p.v.Key(), n.color)
	} else {
		return fmt.Sprintf("node(k: %d, p: <nil>, color: %v)", n.v.Key(), n.color)
	}
}
func (n *Node) isRootNode() bool {
	return !n.isLeaf() && n.p == nil
}

func (n *Node) isLeftNode() bool {
	return !n.isRootNode() && n.p.l == n
}

func (n *Node) isLeaf() bool {
	return n == nil
}

func (n *Node) isBlack() bool {
	return n.isLeaf() || n.color == Black
}

var RotateError = errors.New("required node is leaf")

func rotateR(pp *Node) error {
	if pp.l.isLeaf() || pp.l.l.isLeaf() {
		return RotateError
	}
	p, n := pp.l, pp.l.l
	pp.v, p.v = p.v, pp.v
	if pp.r != nil {
		pp.r.p, n.p = p, pp
	} else {
		n.p = pp
	}
	p.l, p.r, pp.r, pp.l = p.r, pp.r, p, p.l
	return nil
}

func rotateLR(pp *Node) error {
	if pp.l.isLeaf() || pp.l.r.isLeaf() {
		return RotateError
	}
	p, n := pp.l, pp.l.r

	pp.v, n.v = n.v, pp.v
	if n.l != nil {
		n.l.p, n.p = p, pp
	} else {
		n.p = pp
	}
	if n.l != nil {
		n.l.p = p
	}
	if pp.r != nil {
		pp.r.p = n
	}
	n.l, n.r, pp.r, p.r = n.r, pp.r, n, n.l
	return nil
}

func rotateRL(pp *Node) error {
	if pp.r.isLeaf() || pp.r.l.isLeaf() {
		return RotateError
	}
	p, n := pp.r, pp.r.l

	pp.v, n.v = n.v, pp.v
	if n.r != nil {
		n.r.p, n.p = p, pp
	} else {
		n.p = pp
	}
	if n.r != nil {
		n.r.p = p
	}
	if pp.l != nil {
		pp.l.p = n
	}
	n.l, n.r, pp.l, p.l = pp.l, n.l, n, n.r
	return nil
}

func rotateL(pp *Node) error {
	if pp.r.isLeaf() || pp.r.r.isLeaf() {
		return RotateError
	}
	p, n := pp.r, pp.r.r

	pp.v, p.v = p.v, pp.v
	if pp.l != nil {
		pp.l.p, n.p = p, pp
	} else {
		n.p = pp
	}

	p.l, p.r, pp.r, pp.l = pp.l, p.l, p.r, p
	return nil
}
