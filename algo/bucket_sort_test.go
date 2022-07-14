package algo

import (
	"reflect"
	"testing"
)

func TestBucketSort(t *testing.T) {
	tests := []struct {
		input     []int
		inputSize int
		expected  []int
	}{
		{[]int{6, 8, 3, 10, 15, 6, 9, 12, 6, 3}, 10, []int{3, 3, 6, 6, 6, 8, 9, 10, 12, 15}},
		{[]int{12, 9, 24, 4, 19, 21, 14, 6, 2, 16}, 10, []int{2, 4, 6, 9, 12, 14, 16, 19, 21, 24}},
	}

	for _, test := range tests {
		got := BucketSort(test.input, test.inputSize)
		want := test.expected
		if !reflect.DeepEqual(got, want) {
			t.Errorf("BucketSort(%v, %v) = %v, want %v", test.input, test.inputSize, got, want)
		}
	}
}
