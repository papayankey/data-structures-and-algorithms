package algo

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		asc      bool
		expected []int
	}{
		{[]int{8, -1, 5, 3, -2}, true, []int{-2, -1, 3, 5, 8}},
		{[]int{}, true, []int{}},
		{[]int{1, 3, 2, 4, 5}, false, []int{5, 4, 3, 2, 1}},
		{[]int{3, 12, -1, 0, 8, -11}, true, []int{-11, -1, 0, 3, 8, 12}},
	}

	for _, test := range tests {
		got := BubbleSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BubbleSort(%v, %v) = %v; want %v", test.input, test.asc, got, want)
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("fast performance when data list is small", func(b *testing.B) {
		// Average: 46851207 22.84 ns/op
		for i := 0; i < b.N; i++ {
			BubbleSort([]int{8, 2, 5, -1, 0}, true)
		}
	})

	b.Run("slow performance when data list is large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			// Average: 16273899 75.20 ns/op
			BubbleSort([]int{55, 96, 61, 563, 4, 464, 112, 788, 934, 2}, true)
		}
	})
}
