package queue

// queueList represents queue data structure using slice
type queueList struct {
	data []int
	size int
}

// NewQueueList creates an instance of queue
func NewQueueList() *queueList {
	return &queueList{
		data: []int{},
		size: 0,
	}
}

// Enqueue appends item(s) at the rear of queue
func (q *queueList) Enqueue(items ...int) bool {
	q.data = append(q.data, items...)
	q.size += len(items)
	return true
}

// Dequeue removes an item in front of queue
func (q *queueList) Dequeue() int {
	item := q.data[0]
	q.data = q.data[1:]
	q.size -= 1
	return item
}

// Clear resets and cleans up all items in queue
func (q *queueList) Clear() {
	q.data = q.data[:0]
	q.size = 0
}
