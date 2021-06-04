package stack

type (
	Stack struct {
		top    *node
		length uint
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// return the length of stack
func (s *Stack) Len() uint {
	return s.length
}

// return the top item of stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// pop the top item of the stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}
	item := s.top.value
	s.top = s.top.prev
	s.length--
	return item
}

// push the item into the stack
func (s *Stack) Push(item interface{}) {
	s.top = &node{value: item, prev: s.top}
	s.length++
}
