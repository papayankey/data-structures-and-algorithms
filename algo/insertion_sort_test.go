package algo

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		input    []int
		asc      bool
		expected []int
	}{
		{[]int{-8, 7, -1, 2, 4}, true, []int{-8, -1, 2, 4, 7}},
		{[]int{}, true, []int{}},
		{[]int{-8, 7, -1, 2, 4}, false, []int{7, 4, 2, -1, -8}},
	}

	for _, test := range tests {
		got := InsertionSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("InsertionSort(%v, %v) = %v; want %v", test.input, test.asc, got, want)
		}
	}
}
