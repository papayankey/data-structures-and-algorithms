package algo

// BubbleSort is a basic sorting algorithm. It is
// also known as sinking sort. It repeatedly steps through
// a slice or list, compares adjacent elements, and swaps if
// they are in wrong order. It is much less efficient on
// large lists.
//
// Time Complexities:
// - Best Case: O(n)	if already sorted
// - Average Case: O(n^2)
// - Worst Case: O(n^2)
func BubbleSort(s []int, asc bool) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1; j++ {
			if asc {
				if s[j] > s[j+1] {
					swap(s, j, j+1)
				}
			} else {
				if s[j] < s[j+1] {
					swap(s, j, j+1)
				}
			}
		}
	}
	return s
}
