package ds

// node represents an element of the linked list.
type Node struct {
	Data string
	Next *Node
}

// linkedlist represents collection of nodes
type Linkedlist struct {
	Head *Node
	Tail *Node
	Len  int
}

func (l *Linkedlist) String() string {
	out := ""
	curr := l.Head
	for i := 0; i < l.Len; i++ {
		out += curr.Data + " -> "
		curr = curr.Next
	}
	out += "nil"
	return out
}

// NewLinkedList constructs an empty linked list
func NewList() *Linkedlist {
	return new(Linkedlist)
}

// Add inserts node at the beginning of the list
func (l *Linkedlist) Add(v string) {
	l.AddFirst(v)
}

// addFirst inserts node at the beginning of list
func (l *Linkedlist) AddFirst(v string) {
	n := &Node{Data: v}
	if l.Head == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
	l.Len++
}

// AddLast adds node at the end of list
func (l *Linkedlist) AddLast(v string) {
	if l.Head == nil {
		l.AddFirst(v)
	} else {
		n := &Node{Data: v}
		l.Tail.Next = n
		l.Tail = n
		l.Len++
	}
}

// AddBefore inserts node before specified node in list.
// before node must not be nil.
// List is not modified if before is nil or does not exist.
func (l *Linkedlist) AddBefore(bf *Node, v string) {
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
			n := &Node{Data: v}
			n.Next = p.Next
			p.Next = n
			l.Len++
		}
	}
}

// AddAfter inserts node after specified node in list
func (l *Linkedlist) AddAfter(af *Node, v string) {
	panic("To be implemented")
}

// RemoveFirst deletes node from head of list
func (l *Linkedlist) RemoveFirst() (out string) {
	if l.Head == nil {
		return ""
	}
	out = l.Head.Data
	l.Head = l.Head.Next
	l.Len--
	return
}

// RemoveLast deletes node from end of list
func (l *Linkedlist) RemoveLast() string {
	if l.Head == nil {
		return ""
	}

	var prev *Node
	curr := l.Head
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	var out *Node
	if prev == nil {
		out = l.Head
		l.Head = nil
		l.Tail = nil
	} else {
		out = prev.Next
		prev.Next = prev.Next.Next
		l.Tail = prev
	}

	l.Len--
	return out.Data
}

// RemoveAt deletes a node at specified index
func (l *Linkedlist) RemoveAt(index int) string {
	panic("To be implemented")
}

// GetAt retrieves and returns node at specified index
func (l *Linkedlist) GetAt(index int) string {
	if index < 0 || index > l.Len {
		return ""
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	return curr.Data
}

// Contains checks if node exists in list
func (l *Linkedlist) Contains(n *Node) bool {
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
func (l *Linkedlist) findPrev(n *Node) *Node {
	var prev *Node
	curr := l.Head
	for curr != nil && curr.Data != n.Data {
		prev = curr
		curr = curr.Next
	}
	return prev
}
