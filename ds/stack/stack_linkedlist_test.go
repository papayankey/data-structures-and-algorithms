package stack

import (
	"strconv"
	"testing"
)

func TestStackLinkedList(t *testing.T) {
	s := NewStackLinkedList()

	t.Run("adds element to stack", func(t *testing.T) {
		defer s.Clear()
		s.Push("ghana", "nigeria")

		got := s.Top.Data
		want := "nigeria"

		if got != want {
			t.Errorf("Push(%v, %v); got %v, want %v", "ghana", "nigeria", got, want)
		}

		got = strconv.Itoa(s.Size)
		want = "2"

		if got != want {
			t.Errorf("Push(%v, %v); got %v, want %v", "ghana", "nigeria", got, want)
		}
	})

	t.Run("removes element from stack", func(t *testing.T) {
		defer s.Clear()
		s.Push("congo", "benin")

		got := (*s.Pop()).Data
		want := "benin"

		if got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}

		got = strconv.Itoa(s.Size)
		want = "1"

		if got != want {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})
}
