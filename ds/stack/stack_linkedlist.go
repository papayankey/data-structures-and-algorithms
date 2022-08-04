package stack

// node represents element of linked list
// data structure used as container for stack
type node struct {
	Data string
	Next *node
}

// stackLinkedList represents a stack implementation
// using linked list as container
type stackLinkedList struct {
	Top  *node
	Size int
}

// NewStackLinkedList constructs and return a pointer
// to a new stack with default values
func NewStackLinkedList() *stackLinkedList {
	return &stackLinkedList{nil, 0}
}

// Push adds element(s) to the stack
func (s *stackLinkedList) Push(items ...string) {
	for _, item := range items {
		node := &node{item, nil}
		if s.Top == nil {
			s.Top = node
		} else {
			node.Next = s.Top
			s.Top = node
		}
	}
	s.Size += len(items)
}

// Pop removes element from the stack. The stack size
// is reduced by 1 every time this method is invoked
func (s *stackLinkedList) Pop() *node {
	if s.Top == nil {
		return nil
	}
	rem := s.Top
	s.Top = s.Top.Next
	s.Size -= 1
	return rem
}

// Peek return the element at the top of the stack,
// but does not remove it from the stack
func (s *stackLinkedList) Peek() *node {
	panic("To be implemented")
}

// Clear removes all elements from the stack.
// Stack size is reset to zero
func (s *stackLinkedList) Clear() *node {
	panic("To be implemented")
}
