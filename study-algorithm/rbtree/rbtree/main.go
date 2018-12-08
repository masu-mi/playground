package rbtree

import (
	"errors"
	"fmt"
)

type Dict interface {
	Lookup(k Key) (Value, error)
	Insert(k Key, item Value) (Value, error)
	Delete(k Key) (Value, error)
}

type OrderedDict interface {
	Dict
	Scan(start, end Key, step int) (chan Value, error)
}

type Key interface {
	CompareTo(k Key) int
}

type Value interface{}

var NotFoundErr = errors.New("not found")

type RBTree struct {
	root *Node
}

func (tree *RBTree) Lookup(k Key) (Value, error) {
	_, cur := find(tree.root, k)
	if cur != nil {
		return cur.value, nil
	}
	return nil, NotFoundErr
}

func (tree *RBTree) Insert(k Key, item Value) (Value, error) {
	node := tree.set(k, item)
	tree.balance(node)
	return item, nil
}

func (tree *RBTree) set(k Key, item Value) *Node {
	p, cur := find(tree.root, k)
	if cur != nil {
		cur.value = item
		return cur
	}
	cur = &Node{color: RED, k: k, value: item, p: p}
	if p == nil {
		cur.p = cur
	} else {
		if k.CompareTo(p.k) < 0 {
			p.l = cur
		} else {
			p.r = cur
		}
	}
	return cur
}

func find(n *Node, k Key) (p, cur *Node) {
	cur = n
	for cur != nil {
		if diff := k.CompareTo(cur.k); diff == 0 {
			return p, cur
		} else {
			p = cur
			if diff < 0 {
				cur = cur.l
			} else {
				cur = cur.r
			}
		}
	}
	return p, cur
}

func (tree *RBTree) balance(n *Node) {
	cur := n
	for cur.p.color != BLACK {
		if cur.p == cur {
			cur.color = BLACK
			tree.root = cur
			return
		}
		ggp := cur.p.p
		isLeft := ggp.isLeftChild()
		isRoot := ggp.p == ggp
		cur = rotate(cur)
		setColor(cur)
		if isRoot {
			cur.p = cur
		} else if isLeft {
			cur.p, ggp.p.l = ggp.p, cur
		} else {
			cur.p, ggp.p.r = ggp.p, cur
		}
	}
	tree.root = findRoot(cur)
}

func setColor(n *Node) {
	n.color = RED
	if n.l != nil {
		n.l.color = BLACK
	}
	if n.r != nil {
		n.r.color = BLACK
	}
}

func rotate(n *Node) *Node {
	if n.isLeftChild() {
		if n.p.isLeftChild() {
			return rotateR(n.p.p)
		} else { // !n.p.isLeftChild()
			return rotateRL(n.p.p)
		}
	} else { // !n.isLeftChild()
		if n.p.isLeftChild() {
			return rotateLR(n.p.p)
		} else { // !n.p.isLeftChild()
			return rotateL(n.p.p)
		}
	}
}

type placeType int

const (
	root placeType = 0 + iota
	left
	right
)

type place struct {
	t      placeType
	tree   *RBTree
	parent *Node
}

func (tree *RBTree) Delete(k Key) (Value, error) {
	p, cur := find(tree.root, k)
	if cur == nil {
		return nil, NotFoundErr
	}

	if cur.l == nil && cur.r == nil {
		deletedColor, linkT := tree.replaceWith(p, cur, nil)
		if linkT != root && deletedColor == BLACK {
			tree.recoverRank(p, linkT)
		}
	} else {
		srcP, srcCur := findSubstitue(cur)

		var ssrcCur *Node
		if srcCur != nil {
			ssrcCur = srcCur.l
		}
		deletedColor, linkT := tree.replaceWith(srcP, srcCur, ssrcCur)
		_, _, e := tree.updateValueWith(p, cur, srcCur)
		if e != nil {
			panic("not supported")
		}
		if deletedColor == BLACK {
			if srcP.isRoot() {
				tree.recoverRank(tree.root, linkT)
			} else if srcP == cur {
				tree.recoverRank(srcCur, linkT)
			} else {
				tree.recoverRank(srcP, linkT)
			}
		}
	}
	return cur.value, nil
}

func (tree *RBTree) updateValueWith(p, old, new *Node) (Color, placeType, error) {
	if new == nil {
		return RED, root, fmt.Errorf("invalid input; new attr is nil")
	}
	new.l, new.r = old.l, old.r
	if new.l != nil {
		new.l.p = new
	}
	if new.r != nil {
		new.r.p = new
	}
	c, t := tree.replaceWith(p, old, new)
	return c, t, nil
}

func (tree *RBTree) replaceWith(p, old, new *Node) (deleted Color, t placeType) {
	pl := tree.place(p, old)
	if new == nil {
		deleted = pl.Node().color
	} else {
		deleted = new.color
		new.color = pl.Node().color
	}
	pl.setOnPlace(new)
	t = pl.t
	return
}

func (p place) Node() *Node {
	switch p.t {
	case root:
		return p.tree.root
	case left:
		return p.parent.l
	case right:
		return p.parent.r
	}
	return nil
}

func (tree *RBTree) place(p, n *Node) place {
	if p == nil || n.isRoot() {
		return place{t: root, tree: tree, parent: nil}
	} else if n.isLeftChild() {
		return place{t: left, tree: nil, parent: p}
	} else {
		return place{t: right, tree: nil, parent: p}
	}
}

func (p place) setOnPlace(n *Node) {
	switch p.t {
	case root:
		p.tree.root = n
	case left:
		p.parent.l = n
	case right:
		p.parent.r = n
	}
	if n == nil {
		return
	}
	if p.t == root {
		n.p = n
		return
	}
	n.p = p.parent
	return
}

func findSubstitue(n *Node) (p, cur *Node) {
	if n.l != nil {
		if cp, cc := findMax(n.l); cp != nil {
			return cp, cc
		} else {
			return n, cc
		}
	}
	return n, n.r
}

func findMax(n *Node) (p, cur *Node) {
	cur = n
	if n == nil {
		return nil, nil
	}
	for cur.r != nil {
		p = cur
		cur = cur.r
	}
	return p, cur
}

func (tree *RBTree) recoverRank(p *Node, linkT placeType) {
	switch linkT {
	case left:
		tree.recoverRankLeft(p)
	case right:
		tree.recoverRankRight(p)
	case root:
		p.color = BLACK
	default:
		if p != nil {
			panic("invalid attr p is not nil")
		}
	}
	return
}

func (tree *RBTree) recoverRankLeft(p *Node) {
	pp := p.p
	isRoot := p.isRoot()
	isLeft := p.isLeftChild()

	topColor := p.color
	switch p.r.Color() {
	case BLACK:
		if p.r.l.Color() == RED {
			new := rotateRL(p)
			new.l.color, new.color, new.r.color = BLACK, topColor, BLACK
			tree.replaceWith(pp, p, new)
		} else if p.r.r.Color() == RED {
			new := rotateL(p)
			new.l.color, new.color, new.r.color = BLACK, topColor, BLACK
			tree.replaceWith(pp, p, new)
		} else {
			p.color, p.r.color = BLACK, RED
			if topColor == BLACK {
				var t placeType
				if isRoot {
					t = root
				} else if isLeft {
					t = left
				} else {
					t = right
				}
				tree.recoverRank(p.p, t)
				return
			}
		}
	case RED:
		new := rotateL(p)
		new.l.color, new.color = RED, BLACK
		tree.replaceWith(pp, p, new)
		tree.recoverRankLeft(new.l)
		return
	}
	tree.root = findRoot(p)
	return
}

func (tree *RBTree) recoverRankRight(p *Node) {

	pp := p.p
	isRoot := p.isRoot()
	isLeft := p.isLeftChild()

	topColor := p.color
	switch p.l.Color() {
	case BLACK:
		if p.l.r.Color() == RED {
			new := rotateLR(p)
			new.l.color, new.color, new.r.color = BLACK, topColor, BLACK
			tree.replaceWith(pp, p, new)

		} else if p.l.l.Color() == RED {
			new := rotateR(p)
			new.l.color, new.color, new.r.color = BLACK, topColor, BLACK
			tree.replaceWith(pp, p, new)
		} else {
			p.color, p.l.color = BLACK, RED
			if topColor == BLACK {
				var t placeType
				if isRoot {
					t = root
				} else if isLeft {
					t = left
				} else {
					t = right
				}
				tree.recoverRank(p.p, t)
				return
			}
		}
	case RED:
		new := rotateR(p)
		new.r.color, new.color = RED, BLACK
		tree.replaceWith(pp, p, new)
		tree.recoverRankRight(new.r)
		return
	}
	tree.root = findRoot(p)
	return
}
