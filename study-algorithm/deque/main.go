package deque

type Deque struct {
	lq, rq []interface{}
}

func NewDeque(l int) *Deque {
	return &Deque{
		make([]interface{}, 0, l),
		make([]interface{}, 0, l),
	}
}

func (dq *Deque) Empty() bool {
	return len(dq.lq)+len(dq.rq) == 0
}
func (dq *Deque) PushBack(c interface{}) {
	dq.rq = append(dq.rq, c)
}

func (dq *Deque) PushFront(c interface{}) {
	dq.lq = append(dq.lq, c)
}

func (dq *Deque) Pop() (c interface{}) {
	if len(dq.lq) > 0 {
		c = dq.lq[len(dq.lq)-1]
		dq.lq = dq.lq[0 : len(dq.lq)-1]
		return c
	}
	c = dq.rq[0]
	dq.rq = dq.rq[1:len(dq.rq)]
	return c
}
