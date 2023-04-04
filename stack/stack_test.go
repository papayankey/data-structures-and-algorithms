package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	t.Run("adds element to stack", func(t *testing.T) {
		defer s.Clear() // clean up stack

		// push a single element
		s.Push(3)

		got := s.Size
		want := 1

		if got != want {
			t.Errorf("Push(%v), size %v want %v", 3, got, want)
		}

		// push multiple elements
		s.Push(6, 2, 9)

		got = s.Size
		want = 4

		if got != want {
			t.Errorf("Push(%v, %v, %v), size %v want %v", 3, 2, 9, got, want)
		}
	})

	t.Run("removes element from stack", func(t *testing.T) {
		defer s.Clear() // clean up stack

		s.Push(8)
		s.Push(23)

		got := s.Pop()
		want := 23

		if got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}

		got = s.Size
		want = 1

		if got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})

	t.Run("peek element at the top of stack", func(t *testing.T) {
		defer s.Clear() // clean up stack

		s.Push(65)
		s.Push(14)

		got := s.Peek()
		want := 14

		if got != want {
			t.Errorf("Peek() = %v, want %v", got, want)
		}

		got = s.Size
		want = 2

		if got != want {
			t.Errorf("Peek() = %v, want %v", got, want)
		}
	})

	t.Run("exchange first two items of the stack ", func(t *testing.T) {
		defer s.Clear() // clean up stack

		s.Push(97)
		s.Push(98)
		s.Push(99)

		s.Swap()

		got := s.Peek()
		want := 98

		if got != want {
			t.Errorf("Swap() = %v, want %v", got, want)
		}
	})

	t.Run("rotates first n items on stack to the right", func(t *testing.T) {
		defer s.Clear() // clean up stack

		s.Push(0)
		s.Push(1)
		s.Push(2)
		s.Push(3)

		// Given the rotation, n = 3
		//                          Top
		// Before rotate: [ 0, 1, 2, 3 ]
		// After rotate:  [ 0, 3, 1, 2 ]

		s.RotateRight(3)

		got := s.Pop()
		want := 2

		if got != want {
			t.Errorf("RotateRight() = %v, want %v", got, want)
		}

		got = s.Pop()
		want = 1

		if got != want {
			t.Errorf("RotateRight() = %v, want %v", got, want)
		}

		got = s.Pop()
		want = 3

		if got != want {
			t.Errorf("RotateRight() = %v, want %v", got, want)
		}
	})
}
