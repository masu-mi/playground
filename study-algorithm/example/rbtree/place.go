package rbtree

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

func (p place) setOnPlace(n *Node) {
	var link **Node
	switch p.t {
	case root:
		link = &(p.tree.root)
	case left:
		link = &(p.parent.l)
	case right:
		link = &(p.parent.r)
	}

	if n == nil {
		(*link) = nil
		return
	} else if (*link) == nil {
		*link = n
	} else {
		(*link).copyFrom(n)
	}

	if (*link).l != nil {
		(*link).l.p = (*link)
	}
	if (*link).r != nil {
		(*link).r.p = (*link)
	}

	if p.t == root {
		(*link).p = (*link)
		return
	}
	(*link).p = p.parent
	return
}
