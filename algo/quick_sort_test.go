package algo

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		desc   string
		input  []int
		output []int
	}{
		{"test00", []int{2, 9, 7, 15, 11}, []int{2, 7, 9, 11, 15}},
		{"test01", []int{3, 6, 8, 10, 1, 2, 1}, []int{1, 1, 2, 3, 6, 8, 10}},
		{"test02", []int{0, 2, 1, 3, 5}, []int{0, 1, 2, 3, 5}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := QuickSort(test.input)
			want := test.output

			if !reflect.DeepEqual(got, want) {
				t.Errorf("QuickSort(%v) = %v, want %v", test.input, got, want)
			}
		})
	}
}
