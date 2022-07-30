package stack

import "testing"

func TestStack(t *testing.T) {
	s := NewStack()

	t.Run("adds element to stack", func(t *testing.T) {
		defer s.Clear() // clean up stack

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

	t.Run("removes element from stack", func(t *testing.T) {
		defer s.Clear() // clean up stack

		s.Push(8)
		s.Push(23)

		got := s.Pop()
		want := 23

		if got != want {
			t.Errorf("Pop() = %v, want % v", got, want)
		}

		got = s.Size
		want = 1

		if got != want {
			t.Errorf("Pop() = %v, want % v", got, want)
		}
	})

}
