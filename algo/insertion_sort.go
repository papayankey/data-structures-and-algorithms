package algo

// InsertionSort is basic sorting algorithm. It builds
// the final sorted slice (or list) one at a time.
// It is much less efficient on large lists.
//
// Time Complexities:
// - Best Case: O(n) 		if already sorted
// - Average Case: O(kn)	if partially sorted
// - Worst Case: O(n^2)
//
// Advantages:
// - Adaptive: efficient for data sets that are already substantially sorted
// - Stable: does not change the relative order of elements with equal keys
// - In-place: only requires a constant amount of memory space
// - Online: can sort a list as it receives it
func InsertionSort(s []int, asc bool) []int {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if asc {
				if s[j-1] > s[j] {
					swap(s, j, j-1)
				}
			} else {
				if s[j-1] < s[j] {
					swap(s, j, j-1)
				}
			}
		}
	}
	return s
}

// swap exchanges values of given positions
func swap(src []int, i, j int) {
	src[i], src[j] = src[j], src[i]
}
