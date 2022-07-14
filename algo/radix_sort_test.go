package algo

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	tests := []struct {
		input    []int
		radii    int
		expected []int
	}{
		// decimals
		{[]int{237, 146, 259, 348, 152, 163, 235, 48, 36, 62}, 10, []int{36, 48, 62, 146, 152, 163, 235, 237, 259, 348}},
		{[]int{6, 8, 3, 10, 15, 6, 9, 12, 6, 3}, 10, []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}},
		{[]int{12, 9, 24, 4, 19, 21, 14, 6, 2, 16}, 10, []int{2, 4, 6, 9, 12, 14, 16, 19, 21, 24}},
		// binary
		{[]int{1011, 1101, 0011, 0101, 0110}, 2, []int{0011, 0101, 0110, 1011, 1101}},
	}

	for _, test := range tests {
		got := RadixSort(test.input, test.radii)
		want := test.expected

		if !reflect.DeepEqual(got, want) {
			t.Errorf("RadixSort(%v, %v) = %v, want %v", test.input, test.radii, got, want)
		}
	}
}
