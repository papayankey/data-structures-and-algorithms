package ds

// node is an element of the linked list.
// It has both data and memory address of next node
type node struct {
	Data string
	Next *node
}

type linkedlist struct {
	First *node
	Last  *node
	Size  int
}

// NewLinkedList constructs and returns a linked list
func NewLinkedList() *linkedlist {
	return &linkedlist{}
}

// Add inserts node at the beginning of the list
// It delegates its operation to {@code AddFirst}
func (l *linkedlist) Add(val string) bool {
	return l.AddFirst(val)
}

// addFirst adds node at the beginning of the list
func (l *linkedlist) AddFirst(val string) bool {
	n := &node{val, nil}
	if l.First == nil {
		l.Last = n
	} else {
		n.Next = l.First
	}
	l.First = n

	l.Size++
	return true
}

// AddLast adds node at the end of the list
func (l *linkedlist) AddLast(val string) bool {
	if l.First == nil {
		l.AddFirst(val)
	} else {
		n := &node{val, nil}
		l.Last.Next = n
		l.Last = n
	}

	l.Size++
	return true
}

// AddBefore adds node before another node in the list
func (l *linkedlist) AddBefore(before, val string) bool {
	if l.First == nil {
		return false
	}

	if l.First.Data == before {
		return l.AddFirst(val)
	}

	p, ok := l.findPrev(before)
	if !ok {
		return false
	}

	n := &node{val, nil}
	n.Next = p.Next
	p.Next = n

	l.Size++
	return true
}

func (l *linkedlist) AddAfter(after, val string) bool {
	if l.First == nil {
		return false
	}

	n := &node{val, nil}

	if l.First.Data == after {
		n.Next = l.First.Next
		l.First.Next = n
		l.Size++
		return true
	}

	if l.Last.Data == after {
		return l.AddLast(val)
	}

	t := l.First.Next
	for t != nil && t.Data != after {
		t = t.Next
	}

	if t != nil && t.Data == after {
		n.Next = t.Next
		t.Next = n
		l.Size++
		return true
	}

	return false
}

// findPrevious returns preceeding node before given value
func (l *linkedlist) findPrev(val string) (*node, bool) {
	var p *node
	n := l.First

	for n != nil && n.Data != val {
		p = n
		n = n.Next
	}

	if p != nil && p.Next != nil && p.Next.Data == val {
		return p, true
	}

	return nil, false
}
