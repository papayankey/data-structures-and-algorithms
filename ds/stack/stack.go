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
	v := s.Data[s.Size-1]
	s.Data = s.Data[:s.Size-1]
	s.Size -= 1

	return *v
}

// Peek returns element at the top of stack,
// but does not remove elemtn from stack.
func (s *Stack) Peek() int {
	return *s.Data[s.Size-1]
}

// Clear removes all element from stack
// The size is also reset to zero
func (s *Stack) Clear() {
	s.Data = s.Data[:0]
	s.Size = 0
}

// Swap exchanges first two items on stack
func (s *Stack) Swap() {
	i, j := s.Pop(), s.Pop()
	s.Push(i)
	s.Push(j)
}

// RotateRight rotates n items on the stack
// to the right.
func (s *Stack) RotateRight(n int) {
	for i := 1; i < n; i++ {
		c := s.Size - i
		p := s.Size - i - 1
		s.Data[c], s.Data[p] = s.Data[p], s.Data[c]
	}
}
