package ds

// node represents an element of the linked list.
type Node[E comparable] struct {
	Data E
	Next *Node[E]
}

// linkedlist represents collection of nodes
type Linkedlist[E comparable] struct {
	Head *Node[E]
	Tail *Node[E]
	Len  int
}

// func (l *Linkedlist[E]) String() string {
// 	out := ""
// 	curr := l.Head
// 	for i := 0; i < l.Len; i++ {
// 		out += curr.Data + " -> "
// 		curr = curr.Next
// 	}
// 	out += "nil"
// 	return out
// }

// NewLinkedList constructs an empty linked list
func NewList[E comparable]() *Linkedlist[E] {
	return new(Linkedlist[E])
}

// Add inserts node at the beginning of the list
func (l *Linkedlist[E]) Add(v E) {
	l.AddFirst(v)
}

// addFirst inserts node at the beginning of list
func (l *Linkedlist[E]) AddFirst(v E) {
	n := &Node[E]{Data: v}
	if l.Head == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
	l.Len++
}

// AddLast adds node at the end of list
func (l *Linkedlist[E]) AddLast(v E) {
	if l.Head == nil {
		l.AddFirst(v)
	} else {
		n := &Node[E]{Data: v}
		l.Tail.Next = n
		l.Tail = n
		l.Len++
	}
}

// AddBefore inserts node before specified node in list.
// before node must not be nil.
// List is not modified if before is nil or does not exist.
func (l *Linkedlist[E]) AddBefore(bf *Node[E], v E) {
	if bf == nil {
		return
	}

	if l.Head.Data == bf.Data {
		l.AddFirst(v)
		return
	}

	if l.Contains(bf) {
		p := l.findPrev(bf)
		if p != nil {
			n := &Node[E]{Data: v}
			n.Next = p.Next
			p.Next = n
			l.Len++
		}
	}
}

// AddAfter inserts node after specified node in list
func (l *Linkedlist[E]) AddAfter(af *Node[E], v E) {
	panic("to be implemented")
}

// RemoveFirst deletes node from head of list
func (l *Linkedlist[E]) RemoveFirst() (out E) {
	if l.Head == nil {
		return out
	}
	out = l.Head.Data
	l.Head = l.Head.Next
	l.Len--
	return
}

// RemoveLast deletes node from end of list
func (l *Linkedlist[E]) RemoveLast() (out E) {
	if l.Head == nil {
		return out
	}

	var prev *Node[E]
	curr := l.Head
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	var n *Node[E]
	if prev == nil {
		n = l.Head
		l.Head = nil
		l.Tail = nil
	} else {
		n = prev.Next
		prev.Next = prev.Next.Next
		l.Tail = prev
	}

	l.Len--
	out = n.Data
	return
}

// RemoveAt deletes a node at specified index
func (l *Linkedlist[E]) RemoveAt(index int) E {
	panic("Eo be implemented")
}

// GetAt retrieves and returns node at specified index
func (l *Linkedlist[E]) GetAt(index int) (out E) {
	if index < 0 || index > l.Len {
		return out
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	out = curr.Data
	return
}

// Contains checks if node exists in list
func (l *Linkedlist[E]) Contains(n *Node[E]) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == n.Data {
			return true
		}
		curr = curr.Next
	}
	return false
}

// findPrev returns node which preceeds specified node
func (l *Linkedlist[E]) findPrev(n *Node[E]) *Node[E] {
	var prev *Node[E]
	curr := l.Head
	for curr != nil && curr.Data != n.Data {
		prev = curr
		curr = curr.Next
	}
	return prev
}
