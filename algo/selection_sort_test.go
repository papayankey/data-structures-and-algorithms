package algo

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	testCases := []struct {
		input    []int
		asc      bool
		expected []int
	}{
		{[]int{3, 8, 2, -1}, true, []int{-1, 2, 3, 8}},
		{[]int{}, true, []int{}},
		{[]int{2, -1, 5, 3, 0}, false, []int{5, 3, 2, 0, -1}},
	}

	for _, test := range testCases {
		got := SelectionSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("SelectionSort(%v, %v) = %v; want %v", test.input, test.asc, got, want)
		}
	}
}
