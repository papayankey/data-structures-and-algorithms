package stack

import "testing"

func TestStack(t *testing.T) {

	t.Run("adds element to stack", func(t *testing.T) {
		s := NewStack()

		// push a single element
		s.Push(3)

		got := s.Size
		want := 1

		if got != want {
			t.Errorf("Push(%v), size %v want % v", 3, got, want)
		}

		// push multiple elements
		s.Push(6, 2, 9)

		got = s.Size
		want = 4

		if got != want {
			t.Errorf("Push(%v, %v, %v), size %v want % v", 3, 2, 9, got, want)
		}
	})
}
