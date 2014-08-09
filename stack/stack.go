package stack

type node struct {
	data interface{}
	next *node
}

type Stack struct {
	head  *node
	count int
}

func New() *Stack {
	s := &Stack{}
	return s
}

func (s *Stack) Len() int {
	return s.count
}

func (s *Stack) Push(item interface{}) {
	n := &node{data: item}
	if s.head == nil {
		s.head = n
	} else {
		n.next = s.head
		s.head = n
	}

	s.count++
}

func (s *Stack) Pop() interface{} {
	var n *node
	if s.head != nil {
		n = s.head
		s.head = n.next
		s.count--
	}

	if n == nil {
		return nil
	}

	return n.data

}

func (s *Stack) Peek() interface{} {
	n := s.head
	if n == nil || n.data == nil {
		return nil
	}
	return n.data
}
