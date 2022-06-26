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
		{
			input:    []int{8, -1, 5, 3, -2},
			asc:      true,
			expected: []int{-2, -1, 3, 5, 8},
		},
		{
			input:    []int{},
			asc:      true,
			expected: []int{},
		},
		{
			input:    []int{1, 3, 2, 4, 5},
			asc:      false,
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			input:    []int{3, 12, -1, 0, 8, -11},
			asc:      true,
			expected: []int{-11, -1, 0, 3, 8, 12},
		},
	}

	for _, test := range tests {
		got := BubbleSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
}
