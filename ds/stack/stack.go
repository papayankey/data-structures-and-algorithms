package stack

// Stack represent data structure in which
// elements are stored in last-in-first-out (LIFO)
// principle. Stack stores pointers to element.
type Stack struct {
	Data []*int
	Size int
}

// NewStack constructs and returns
// new pointer to stack
func NewStack() *Stack {
	return new(Stack)
}

// Push adds element(s) to stack
func (s *Stack) Push(items ...int) {
	for _, item := range items {
		s.Data = append(s.Data, &item)
	}
	s.Size += len(items)
}

// Pop removes an element from stack
func (s *Stack) Pop() int {
	panic("To be implemented")
}

// Peek returns element at the top of stack,
// but does not remove elemtn from stack.
func (s *Stack) Peek() int {
	panic("To be implemented")
}
