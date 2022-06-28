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
		{
			input:    []int{3, 8, 2, -1},
			asc:      true,
			expected: []int{-1, 2, 3, 8},
		},
		{
			input:    []int{},
			asc:      true,
			expected: []int{},
		}, {
			input:    []int{2, -1, 5, 3, 0},
			asc:      false,
			expected: []int{5, 3, 2, 0, -1},
		},
	}

	for _, test := range testCases {
		got := SelectionSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
