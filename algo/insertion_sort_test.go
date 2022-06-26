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
		{
			input:    []int{-8, 7, -1, 2, 4},
			asc:      true,
			expected: []int{-8, -1, 2, 4, 7},
		},
		{
			input:    []int{},
			asc:      true,
			expected: []int{},
		},
		{
			input:    []int{-8, 7, -1, 2, 4},
			asc:      false,
			expected: []int{7, 4, 2, -1, -8},
		},
	}

	for _, test := range tests {
		got := InsertionSort(test.input, test.asc)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}
