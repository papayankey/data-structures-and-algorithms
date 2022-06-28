package algo

// SelectionSort is a basic sorting algorithms.
// It is much less efficient on large lists.
// It is also much less efficient than insertion sort.
// Elements are build up either at left or right depending
// on the sorting order.
//
// Time Complexities:
// Best Case: O(n^2)
// Average Case: O(n^2)
// Worst Case: O(n^2)
func SelectionSort(s []int, asc bool) []int {
	for i := 0; i < len(s); i++ {
		k := i
		for j := i + 1; j < len(s); j++ {
			if asc {
				if s[j] < s[k] {
					k = j
				}
			} else {
				if s[j] > s[k] {
					k = j
				}
			}

		}
		swap(s, i, k)
	}
	return s
}
