package hlist

type head struct {
	first *Element
}

type Element struct {
	next  *Element
	pprev **Element
	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

type Hlist struct {
	root head
	size int
}

func (l *Hlist) init() *Hlist {
	l.root.first = nil
	l.size = 0
	return l
}

func New() *Hlist { return new(Hlist).init() }

func (l *Hlist) Len() int { return l.size }

func (l *Hlist) Front() *Element {
	return l.root.first
}

func (l *Hlist) PushFront(v interface{}) *Element {
	first := l.root.first
	n := &Element{Value: v}
	n.next = first
	if first != nil {
		first.pprev = &n.next
	}
	l.root.first = n
	n.pprev = &l.root.first
	l.size++
	return n
}

func (l *Hlist) Remove(e *Element) interface{} {
	next := e.next
	pprev := e.pprev
	*pprev = next
	if next != nil {
		next.pprev = pprev
	}
	l.size--
	e.next = nil
	e.pprev = nil
	return e.Value
}
