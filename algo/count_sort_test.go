package algo

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	tests := []struct {
		input  []int
		size   int
		output []int
	}{
		{[]int{8, 4, 1, 3, 1, 8, 2}, 7, []int{1, 1, 2, 3, 4, 8, 8}},
		{[]int{6, 8, 3, 10, 15, 6, 9, 12, 6, 3}, 10, []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}},
	}

	for _, test := range tests {
		got := CountSort(test.input, test.size)
		want := test.output

		if !reflect.DeepEqual(got, want) {
			t.Errorf("CountSort(%v, %v) = %v, want %v", test.input, test.size, got, want)
		}
	}

}
