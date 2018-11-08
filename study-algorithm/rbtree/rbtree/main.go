package rbtree

import (
	"errors"
)

type RBTree struct {
	/// TODO change RBType definition to `tyep RBType *Node`$
	root *Node
}

type Item interface {
	Key() int
}

var (
	NotFoundError      = errors.New("not found")
	RotateUnmatchError = errors.New("rotate unmatch pattern")
)

func new() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) Lookup(key int) (Item, error) {
	ref, _ := tree.search(key)
	if *ref == nil {
		return nil, NotFoundError
	}
	return (*ref).v, nil
}
func (tree *RBTree) Insert(item Item) (Item, error) {
	ref, p := tree.search(item.Key())
	if *ref == nil {
		*ref = &Node{color: Red, v: item, p: p}
	} else {
		(*ref).v = item
	}
	balance(*ref)
	return item, nil
}

func balance(n *Node) {
	if n.p == nil {
		n.color = Black
		return
	}
	if n.color == Red && n.p.color == Red {
		pp := n.p.p
		rotate(n)
		l, r := pp.l, pp.r
		l.color, pp.color, r.color = Black, Red, Black
		balance(pp)
	}
}

func rotate(n *Node) error {
	p, pp := n.p, n.p.p
	if p.l == n && pp.l == p {
		return rotateR(pp)
	} else if p.r == n && pp.l == p {
		return rotateLR(pp)
	} else if p.l == n && pp.r == p {
		return rotateRL(pp)
	} else if p.r == n && pp.r == p {
		return rotateL(pp)
	}
	return RotateUnmatchError
}

func (tree *RBTree) Delete(key int) (Item, error) {
	ref, p := tree.search(key)
	if *ref == nil {
		return nil, NotFoundError
	}
	deletedColor, p, t := replace(ref, p)
	if deletedColor == Black {
		recoverRank(p, t)
	}
	return nil, nil
}

func recoverRank(p *Node, n *Node) (err error) {
	if n.isRootNode() {
		n.color = Black
		return
	}
	var isLeftNode bool
	if n.isLeaf() {
		isLeftNode = p.l == nil
	} else {
		isLeftNode = n.isLeftNode()
	}
	if isLeftNode {
		switch {
		case (!p.r.isLeaf() && p.r.isBlack() && p.r.l.isBlack() && p.r.r.isBlack()):
			p.r.color = Red
			fallthrough
		case p.r.isLeaf():
			if p.color == Black {
				recoverRank(p.p, p)
			}
			p.color = Black
			return
		case p.r.isBlack() && !p.r.l.isBlack():
			c := p.color
			e := rotateRL(p)
			if e != nil {
				return e
			}
			p.color, p.r.color = c, Black
			if p.l != nil {
				p.l.color = Black
			}
			return
		case p.r.isBlack() && !p.r.r.isBlack():
			c := p.color
			e := rotateL(p)
			if e != nil {
				return e
			}
			p.color, p.r.color = c, Black
			if p.l != nil {
				p.l.color = Black
			}
			return
		case !p.r.isBlack():
			e := rotateL(p)
			if e != nil {
				return e
			}
			p.l.color, p.color = Red, Black
			recoverRank(p.l, p.l.l)
			return
		}
	} else {
		switch {
		case (!p.l.isLeaf() && p.l.isBlack() && p.l.l.isBlack() && p.l.r.isBlack()):
			p.l.color = Red
			fallthrough
		case p.l.isLeaf():
			if p.color == Black {
				recoverRank(p.p, p)
			}
			p.color = Black
			return
		case p.l.isBlack() && !p.l.r.isBlack():
			c := p.color
			e := rotateLR(p)
			if e != nil {
				return e
			}
			p.l.color, p.color, p.r.color = Black, c, Black
			return
		case p.l.isBlack() && !p.l.l.isBlack():
			c := p.color
			e := rotateR(p)
			if e != nil {
				return e
			}
			p.l.color, p.color, p.r.color = Black, c, Black
			return
		case !p.l.isBlack():
			e := rotateR(p)
			if e != nil {
				return e
			}
			p.l.color, p.color = Red, Black
			return
		}
	}
	return
}

func replace(node **Node, p *Node) (Color, *Node, *Node) {
	c := (**node).color
	if (**node).l == nil {
		if (**node).r == nil {
			*node = nil
			return c, p, *node
		} else {
			deletedColor := (**node).r.color
			(**node).r.color, (**node).r.p = c, p
			*node = (**node).r
			return deletedColor, p, *node
		}
	} else {
		infRef, infP := searchInf(*node)
		(**node).v = (*infRef).v

		if (*infRef).l == nil {
			deletedColor := (*infRef).color
			*infRef = nil
			return deletedColor, infP, *infRef
		} else {
			deletedColor := (*infRef).l.color
			(*infRef).l.color, (*infRef).l.p = (*infRef).color, infP
			*infRef = (*infRef).l
			return deletedColor, infP, *infRef
		}
	}
}

func searchInf(node *Node) (**Node, *Node) {
	if node.l.r == nil {
		return &(node.l), node
	}
	n := node.l
	for n.r != nil {
		n = n.r
	}
	return &(n.p.r), n.p
}

func (tree *RBTree) search(key int) (ref **Node, p *Node) {
	for ref = &(tree.root); *ref != nil; {
		if nodeKey := (*ref).v.Key(); nodeKey > key {
			ref, p = &((*ref).l), *ref
		} else if nodeKey == key {
			return ref, p
		} else {
			ref, p = &((*ref).r), *ref
		}
	}
	return ref, p
}

type Event int

const (
	Ent Event = iota
	PreLeft
	PostLeft
	PreRight
	PostRight
	Exit
)

type handler func(e Event, n Node)

func (tree *RBTree) Traverse(h handler) {
	if ref := &(tree.root); *ref != nil {
		(**ref).traverse(h)
	}
}

func (node Node) traverse(h handler) {
	h(Ent, node)
	if node.l != nil {
		h(PreLeft, node)
		node.l.traverse(h)
		h(PostLeft, node)
	}
	if node.r != nil {
		h(PreRight, node)
		node.r.traverse(h)
		h(PostRight, node)
	}
	h(Exit, node)
}
