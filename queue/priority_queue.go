package queue

import "fmt"

type priorityQueue struct {
	data      []int
	size      int
	isMaximum bool
}

// NewPriorityQueue constructs an instance priority queue
func NewPriorityQueue(maxPriority bool) *priorityQueue {
	return &priorityQueue{
		data:      []int{},
		size:      0,
		isMaximum: maxPriority,
	}
}

// Enqueue appends item to end of priority queue
func (pq *priorityQueue) Enqueue(value int) {
	pq.data = append(pq.data, value)

	if !isEmpty(pq) {
		index := len(pq.data) - 1
		if pq.isMaximum {
			maxHeapifyUp(pq, index)
		} else {
			minHeapifyUp(pq, index)
		}
	}

	pq.size += 1
}

// Dequeue removes item from front of priority queue
func (pq *priorityQueue) Dequeue() int {
	item := pq.data[0]
	swap(pq, 0, len(pq.data)-1)
	pq.data = pq.data[:len(pq.data)-1]
	pq.size -= 1

	if pq.isMaximum {
		minHeapifyDown(pq, 0)
	} else {
		maxHeapifyDown(pq, 0)
	}

	return item
}

func minHeapifyDown(pq *priorityQueue, index int) {
	for index < pq.size {
		left := left(index)
		right := right(index)
		var largest int

		if right >= pq.size || pq.data[left] > pq.data[right] {
			largest = left
		} else {
			largest = right
		}

		if largest < pq.size && pq.data[largest] > pq.data[index] {
			swap(pq, largest, index)
			index = largest
		} else {
			break
		}
	}
}

func maxHeapifyDown(pq *priorityQueue, index int) {
	for index < pq.size {
		left := left(index)
		right := right(index)
		var smallest int

		if right >= pq.size || pq.data[left] < pq.data[right] {
			smallest = left
		} else {
			smallest = right
		}

		if smallest < pq.size && pq.data[smallest] < pq.data[index] {
			swap(pq, smallest, index)
			index = smallest
		} else {
			break
		}
	}
}

func minHeapifyUp(pq *priorityQueue, index int) {
	for pq.data[index] < pq.data[parent(index)] {
		swap(pq, parent(index), index)
		index = parent(index)
	}
}

func maxHeapifyUp(pq *priorityQueue, index int) {
	for pq.data[index] > pq.data[parent(index)] {
		swap(pq, parent(index), index)
		index = parent(index)
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

func isEmpty(pq *priorityQueue) bool {
	return pq.size == 0
}

func (pq *priorityQueue) String() string {
	return fmt.Sprintf("%v", pq.data)
}
