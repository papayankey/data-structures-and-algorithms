package ds

import (
	"reflect"
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := NewLinkedList()
	f, o := "Pawpaw", "Orange"

	l.Add(o)
	l.AddFirst(f)

	got := l.First
	want := &node{f, &node{o, nil}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddAtFirst(%v) = %v; want %v", f, got, want)
	}
}

func TestAddLast(t *testing.T) {
	l := NewLinkedList()
	f, g := "Pineapple", "Guava"

	l.Add(f)
	l.AddLast(g)

	got := l.Last
	want := &node{g, nil}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddAtLast(%v) = %v; want %v", f, got, want)
	}
}

func TestAddBefore(t *testing.T) {
	t.Run("given that list is empty Returns false", func(t *testing.T) {
		l := NewLinkedList()

		b, v := "Pineapple", "Guava"
		got := l.AddBefore(b, v)
		want := false

		if got != want {
			t.Errorf("AddBefore(%v, %v) = %v; want %v", b, v, got, want)
		}
	})

	t.Run("given that before does not exist Returns false", func(t *testing.T) {
		l := NewLinkedList()
		l.Add("Pear")

		b, v := "Pineapple", "Guava"
		got := l.AddBefore(b, v)
		want := false

		if got != want {
			t.Errorf("AddBefore(%v, %v) = %v; want %v", b, v, got, want)
		}
	})

	t.Run("given that before exists New node is inserted before given node", func(t *testing.T) {
		l := NewLinkedList()
		b, v := "Tangerine", "Apple"

		l.Add(b)
		l.Add("Watermelon")
		l.AddBefore(b, v)

		got := l.First.Next
		want := &node{v, &node{b, nil}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddBefore(%v, %v) = %v; want %v", b, v, got, want)
		}
	})

	t.Run("given that before node is the first New node is inserted at first", func(t *testing.T) {
		l := NewLinkedList()
		b, v := "Pineapple", "Guava"

		l.Add(b)
		l.AddBefore(b, v)

		got := l.First
		want := &node{v, &node{b, nil}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddBefore(%v, %v) = %v; want %v", b, v, got, want)
		}
	})
}

func TestAddAfter(t *testing.T) {
	t.Run("given that list is empty Returns false", func(t *testing.T) {
		l := NewLinkedList()

		a, v := "Bitter Gourd", "Mango"
		got := l.AddAfter(a, v)
		want := false

		if got != want {
			t.Errorf("AddAfter(%v, %v) = %v; want %v", a, v, got, want)
		}
	})

	t.Run("given that after is the first New node is inserted after first", func(t *testing.T) {
		l := NewLinkedList()
		l.Add("Bitter Gourd")

		a, v := "Bitter Gourd", "Mango"
		l.AddAfter(a, v)

		got := l.First.Next
		want := &node{v, nil}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddAfter(%v, %v) = %v; want %v", a, v, got, want)
		}
	})

	t.Run("given that after is the last New node is inserted at last", func(t *testing.T) {
		l := NewLinkedList()
		l.Add("Sweet Apple")
		l.Add("Bitter Gourd")

		a, v := "Sweet Apple", "Mango"
		l.AddAfter(a, v)

		got := l.Last
		want := &node{v, nil}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("AddAfter(%v, %v) = %v; want %v", a, v, got, want)
		}
	})
}
