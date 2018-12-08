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
	created := tree.set(k, item)
	tree.recoverBalance(created)
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
		tree.root = cur
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

func (tree *RBTree) recoverBalance(n *Node) {
	cur := n
	for cur.p.color != BLACK {
		if cur.p == cur {
			cur.color = BLACK
			return
		}
		cur = balance(cur)
	}
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

func balance(n *Node) *Node {
	gp := n.p.p
	if n.isLeftChild() {
		if n.p.isLeftChild() {
			rotateR(gp)
		} else { // !n.p.isLeftChild()
			rotateRL(gp)
		}
	} else { // !n.isLeftChild()
		if n.p.isLeftChild() {
			rotateLR(gp)
		} else { // !n.p.isLeftChild()
			rotateL(gp)
		}
	}
	setColor(gp)
	return gp
}

func (tree *RBTree) Delete(k Key) (Value, error) {
	p, cur := find(tree.root, k)
	if cur == nil {
		return nil, NotFoundErr
	}

	if cur.l == nil && cur.r == nil {
		deletedColor, pl := tree.replaceWith(p, cur, nil)
		if deletedColor == BLACK {
			tree.recoverRank(pl)
		}
	} else {
		srcP, srcCur := findSubstitue(cur)

		var ssrcCur *Node
		if srcCur != nil {
			ssrcCur = srcCur.l
		}
		deletedColor, pl := tree.replaceWith(srcP, srcCur, ssrcCur)
		e := tree.updateValueWith(p, cur, srcCur)
		if e != nil {
			panic("not supported")
		}
		if deletedColor == BLACK {
			tree.recoverRank(pl)
		}
	}
	return cur.value, nil
}

func (tree *RBTree) updateValueWith(p, old, new *Node) error {
	if new == nil {
		return fmt.Errorf("invalid input; new attr is nil")
	}
	new.l, new.r = old.l, old.r
	if new.l != nil {
		new.l.p = new
	}
	if new.r != nil {
		new.r.p = new
	}
	_, _ = tree.replaceWith(p, old, new)
	return nil
}

func (tree *RBTree) replaceWith(p, old, new *Node) (lost Color, pl place) {
	pl = tree.place(p, old)
	if new == nil {
		lost = pl.Node().color
	} else {
		lost = new.color
		new.color = pl.Node().color
	}
	pl.setOnPlace(new)
	return
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

func (tree *RBTree) recoverRank(pl place) {
	switch pl.t {
	case left:
		tree.recoverRankLeft(pl.parent)
	case right:
		tree.recoverRankRight(pl.parent)
	case root:
		pl.tree.root.color = BLACK
	default:
		panic("invalid attr place's t")
	}
	return
}

func (tree *RBTree) recoverRankLeft(p *Node) {
	pp := p.p
	pl := tree.place(pp, p)

	topColor := p.color
	switch p.r.Color() {
	case BLACK:
		if p.r.l.Color() == RED {
			rotateRL(p)
			p.l.color, p.color, p.r.color = BLACK, topColor, BLACK
		} else if p.r.r.Color() == RED {
			rotateL(p)
			p.l.color, p.color, p.r.color = BLACK, topColor, BLACK
		} else {
			p.color, p.r.color = BLACK, RED
			if topColor == BLACK {
				tree.recoverRank(pl)
				return
			}
		}
	case RED:
		rotateL(p)
		p.l.color, p.color = RED, BLACK
		tree.recoverRankLeft(p.l)
		return
	}
	return
}

func (tree *RBTree) recoverRankRight(p *Node) {

	pp := p.p
	pl := tree.place(pp, p)

	topColor := p.color
	switch p.l.Color() {
	case BLACK:
		if p.l.r.Color() == RED {
			rotateLR(p)
			p.l.color, p.color, p.r.color = BLACK, topColor, BLACK
		} else if p.l.l.Color() == RED {
			rotateR(p)
			p.l.color, p.color, p.r.color = BLACK, topColor, BLACK
		} else {
			p.color, p.l.color = BLACK, RED
			if topColor == BLACK {
				tree.recoverRank(pl)
				return
			}
		}
	case RED:
		rotateR(p)
		p.r.color, p.color = RED, BLACK
		tree.recoverRankRight(p.r)
		return
	}
	return
}
