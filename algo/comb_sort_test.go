package algo

import (
	"reflect"
	"testing"
)

func TestCombSort(t *testing.T) {
	tests := []struct {
		desc          string
		size          int
		input, output []int
	}{
		{"test00", 5, []int{5, 3, 8, 2, 0}, []int{0, 2, 3, 5, 8}},
		{"test01", 7, []int{49, 11, 24, 44, 29, 27, 2}, []int{2, 11, 24, 27, 29, 44, 49}},
		{"test02", 7, []int{48, 10, 23, 43, 28, 26, 1}, []int{1, 10, 23, 26, 28, 43, 48}},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := CombSort(test.input, test.size)
			want := test.output

			if !reflect.DeepEqual(got, want) {
				t.Errorf("CombSort(%v, %v) = %v, want %v", test.input, test.size, got, want)
			}
		})
	}
}
