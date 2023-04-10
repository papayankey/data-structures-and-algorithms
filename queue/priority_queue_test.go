package queue

import (
	"reflect"
	"testing"
)

func TestMaximumPriorityEnqueue(t *testing.T) {
	pq := NewPriorityQueue(true)

	tests := []struct {
		desc  string
		input int
		out   string
	}{
		{"test00", 3, "[3]"},
		{"test01", 6, "[6 3]"},
		{"test02", 8, "[8 3 6]"},
		{"test03", 2, "[8 3 6 2]"},
		{"test04", 0, "[8 3 6 2 0]"},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			pq.Enqueue(test.input)

			got := pq.String()
			want := test.out

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Enqueue(%d) = want %s but got %s", test.input, want, got)
			}
		})
	}
}

func TestMaximumPriorityDequeue(t *testing.T) {
	pq := NewPriorityQueue(true)

	inputs := []int{3, 6, 8, 2, 0}
	for _, input := range inputs {
		pq.Enqueue(input)
	}

	tests := []struct {
		desc string
		out  string
	}{
		{"test00", "[6 3 0 2]"},
		{"test01", "[3 2 0]"},
		{"test02", "[2 0]"},
		{"test03", "[0]"},
		{"test04", "[]"},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			pq.Dequeue()

			got := pq.String()
			want := test.out

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Dequeue() = want %s but got %s", want, got)
			}
		})
	}
}

func TestMinimumPriorityEnqueue(t *testing.T) {
	pqueue := NewPriorityQueue(false)

	tests := []struct {
		desc  string
		input int
		out   string
	}{
		{"test00", 3, "[3]"},
		{"test01", 6, "[3 6]"},
		{"test02", 8, "[3 6 8]"},
		{"test03", 2, "[2 3 8 6]"},
		{"test04", 0, "[0 2 8 6 3]"},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			pqueue.Enqueue(test.input)

			got := pqueue.String()
			want := test.out

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Enqueue(%d) = want %s but got %s", test.input, want, got)
			}
		})
	}
}

func TestMinimumPriorityDequeue(t *testing.T) {
	pq := NewPriorityQueue(false)

	inputs := []int{3, 6, 8, 2, 0}
	for _, input := range inputs {
		pq.Enqueue(input)
	}

	tests := []struct {
		desc string
		out  string
	}{
		{"test00", "[2 3 8 6]"},
		{"test01", "[3 6 8]"},
		{"test02", "[6 8]"},
		{"test03", "[8]"},
		{"test04", "[]"},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			pq.Dequeue()

			got := pq.String()
			want := test.out

			if !reflect.DeepEqual(got, want) {
				t.Errorf("Dequeue() = want %s but got %s", want, got)
			}
		})
	}
}
