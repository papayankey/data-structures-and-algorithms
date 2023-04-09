package queue

import "testing"

func TestEnqueue(t *testing.T) {
	queue := NewQueueList()

	got := queue.Enqueue(1)
	want := true

	if got != want {
		t.Errorf("Enqueue(1) = %t but want %t", got, want)
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueueList()

	queue.Enqueue(1, 2, 3, 4, 5)

	got := queue.Dequeue()
	want := 1

	if got != want {
		t.Errorf("Dequeue() = %d but want %d", got, want)
	}
}

func TestClear(t *testing.T) {
	queue := NewQueueList()

	queue.Enqueue(1, 2, 3, 4, 5)
	queue.Clear()

	got := queue.size
	want := 0

	if got != want {
		t.Errorf("Clear() = %d but want %d", got, want)
	}
}
