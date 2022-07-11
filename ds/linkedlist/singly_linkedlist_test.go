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

func TestGetAt(t *testing.T) {
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

func TestRemoveAt(t *testing.T) {
	l := NewList[string]()

	// when list is empty

	got := l.RemoveAt(1)
	want := ""

	if got != want {
		t.Errorf("RemoveAt() = %v; want %v", got, want)
	}

	// when index is the head

	l.Add("grapes")

	got = l.RemoveAt(1)
	want = "grapes"

	if got != want {
		t.Errorf("RemoveAt() = %v; want %v", got, want)
	}

	// when index is the tail

	l.Add("watermelon")
	l.Add("strawberry")
	l.Add("passion fruit")

	got = l.RemoveAt(3)
	want = "watermelon"

	if got != want {
		t.Errorf("RemoveAt() = %v; want %v", got, want)
	}

	// when index is neither head or tail

	l.Add("watermelon")
	l.Add("strawberry")
	l.Add("passion fruit")
	l.Add("avocado")

	got = l.RemoveAt(3)
	want = "strawberry"

	if got != want {
		t.Errorf("RemoveAt() = %v; want %v", got, want)
	}
}

func TestAddAfter(t *testing.T) {
	l := NewList[string]()

	// when specified node is the tail

	l.Add("manchester united")
	l.Add("real madrid")

	af := &Node[string]{Data: "manchester united"}
	l.AddAfter(af, "chelsea")

	got := l.Head.Next.Next.Data
	want := "chelsea"

	if got != want {
		t.Errorf("AddAfter(%v, %v); got %v want %v", af, "chelsea", got, want)
	}

	// when specified node is the head

	af = &Node[string]{Data: "real madrid"}
	l.AddAfter(af, "ajax")

	got = l.Head.Next.Data
	want = "ajax"

	if got != want {
		t.Errorf("AddAfter(%v, %v); got %v want %v", af, "ajax", got, want)
	}

	// when specified node is not the head or tail

	af = &Node[string]{Data: "ajax"}
	l.AddAfter(af, "liverpool")

	got = l.Head.Next.Next.Data
	want = "liverpool"

	if got != want {
		t.Errorf("AddAfter(%v, %v); got %v want %v", af, "liverpool", got, want)
	}

	// when list is empty

	l = NewList[string]()

	af = &Node[string]{Data: "bayern munich"}
	l.AddAfter(af, "barcelona")

	var actual *Node[string] = l.Head
	var expected *Node[string] = nil

	if actual != expected {
		t.Errorf("AddAfter(%v, %v); got %v want %v", af, "barcelona", actual, expected)
	}
}
