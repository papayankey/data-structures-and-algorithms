package queue

import "fmt"

type priorityQueue struct {
	data        []int
	size        int
	maxPriority bool
}

// NewPriorityQueue constructs an instance of priority queue
func NewPriorityQueue(maxPriority bool) *priorityQueue {
	return &priorityQueue{
		data:        []int{},
		size:        0,
		maxPriority: maxPriority,
	}
}

// Enqueue appends item to end of priority queue
func (pq *priorityQueue) Enqueue(value int) {
	pq.data = append(pq.data, value)
	index := len(pq.data) - 1

	heapifyUp(pq, index, pq.maxPriority)

	pq.size += 1
}

// Dequeue removes item from front of priority queue
func (pq *priorityQueue) Dequeue() int {
	item := pq.data[0]
	swap(pq, 0, len(pq.data)-1)
	pq.data = pq.data[:len(pq.data)-1]
	pq.size -= 1

	heapifyDown(pq, 0, pq.maxPriority)

	return item
}

func (pq *priorityQueue) String() string {
	return fmt.Sprintf("%v", pq.data)
}

func heapifyDown(pq *priorityQueue, index int, maxPriority bool) {
	for {
		if index >= pq.size {
			break
		}

		left := left(index)
		right := right(index)
		current := left

		if right < pq.size {
			if maxPriority && pq.data[right] > pq.data[left] {
				current = right
			} else if !maxPriority && pq.data[right] < pq.data[left] {
				current = right
			}
		}

		if current >= pq.size {
			break
		}

		if maxPriority && !(pq.data[current] > pq.data[index]) {
			break
		} else if !maxPriority && !(pq.data[current] < pq.data[index]) {
			break
		}

		swap(pq, current, index)
		index = current
	}
}

func heapifyUp(pq *priorityQueue, index int, maxPriority bool) {
	for {
		parent := parent(index)

		if maxPriority && !(pq.data[index] > pq.data[parent]) {
			break
		} else if !maxPriority && !(pq.data[index] < pq.data[parent]) {
			break
		}

		swap(pq, parent, index)
		index = parent
	}
}

func swap(pq *priorityQueue, i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func parent(index int) int {
	return (index - 1) / 2
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}
