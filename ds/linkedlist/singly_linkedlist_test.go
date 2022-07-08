package ds

import (
	"testing"
)

func TestAddFirst(t *testing.T) {
	l := NewList[string]()
	l.AddFirst("pawpaw")

	got := l.Head.Data
	want := "pawpaw"

	if got != want {
		t.Errorf("AddAtFirst(%v) = %v; want %v", "pawpaw", got, want)
	}

	l.AddFirst("mango")

	got = l.Head.Data
	want = "mango"

	if got != want {
		t.Errorf("AddAtFirst(%v) = %v; want %v", "mango", got, want)
	}
}

func TestAddLast(t *testing.T) {
	l := NewList[string]()

	l.AddLast("pineapple")

	got := l.Tail.Data
	want := "pineapple"

	if got != want {
		t.Errorf("AddAtLast(%v) = %v; want %v", "pineapple", got, want)
	}

	l.AddLast("blackberry")

	got = l.Tail.Data
	want = "blackberry"

	if got != want {
		t.Errorf("AddAtLast(%v) = %v; want %v", "blackberry", got, want)
	}
}

func TestAddBefore(t *testing.T) {
	l := NewList[string]()

	// when before is nil

	l.Add("orange")
	l.AddBefore(nil, "pear")

	got := l.Head.Data
	want := "orange"

	if got != want {
		t.Errorf("AddBefore(%v, %v) = %v; want %v", nil, "blackberry", got, want)
	}

	// when before is not nil but does not exist

	l.AddBefore(&Node[string]{Data: "avocado"}, "pear")

	got = l.Head.Data
	want = "orange"

	if got != want {
		t.Errorf("AddBefore(%v, %v) = %v; want %v", "avocado", "orange", got, want)
	}

	// when before exists and is the head

	l.Add("avocado")
	l.AddBefore(&Node[string]{Data: "avocado"}, "pear")

	got = l.Head.Data
	want = "pear"

	if got != want {
		t.Errorf("AddBefore(%v, %v) = %v; want %v", "avocado", "pear", got, want)
	}

	// when before exists and not the head

	l.AddBefore(&Node[string]{Data: "orange"}, "mango")

	got = l.Head.Next.Next.Data
	want = "mango"

	if got != want {
		t.Errorf("AddBefore(%v, %v) = %v; want %v", "orange", "mango", got, want)
	}
}

func TestRemoveFirst(t *testing.T) {
	l := NewList[string]()

	// when list is empty

	got := l.RemoveFirst()
	want := ""

	if got != want {
		t.Errorf("RemoveFirst() = %v; want %v", got, want)
	}

	// when list contains at least 1 node

	l.Add("lime")
	l.Add("melon")

	got = l.RemoveFirst()
	want = "melon"

	if got != want {
		t.Errorf("RemoveFirst() = %v; want %v", got, want)
	}
}

func TestGet(t *testing.T) {
	l := NewList[string]()

	l.Add("starfruit")
	l.Add("strawberry")
	l.Add("kiwi")

	// when index is the head

	got := l.GetAt(0)
	want := "kiwi"

	if got != want {
		t.Errorf("Get(%v) = %v; want %v", 0, got, want)
	}

	// when index is less than zero

	got = l.GetAt(-1)
	want = ""

	if got != want {
		t.Errorf("Get(%v) = %v; want %v", -1, got, want)
	}

	// when index is greater than list size

	got = l.GetAt(4)
	want = ""

	if got != want {
		t.Errorf("Get(%v) = %v; want %v", 4, got, want)
	}

	// when index is within range of list size

	got = l.GetAt(3)
	want = "starfruit"

	if got != want {
		t.Errorf("Get(%v) = %v; want %v", 3, got, want)
	}
}

func TestRemoveLast(t *testing.T) {
	l := NewList[string]()

	// when	list is empty

	got := l.RemoveLast()
	want := ""

	if got != want {
		t.Errorf("RemoveLast() = %v; want %v", got, want)
	}

	// when list has 1 node

	l.Add("tangerine")

	got = l.RemoveLast()
	want = "tangerine"

	if got != want {
		t.Errorf("RemoveLast() = %v; want %v", got, want)
	}

	// when list has more than 1 node

	l.Add("peach")
	l.Add("sweet apple")
	l.Add("orange")

	got = l.RemoveLast()
	want = "peach"

	if got != want {
		t.Errorf("RemoveLast() = %v; want %v", got, want)
	}
}
