package algo

// CombSort is comparison sort.
// It is little improvement over bubble sort.
// Whilst bubble sort compares and swaps adjacent elements
// comb sort swaps from wider distance starting with size of
// list as gap size and shrinks down to 1 using a shrink
// factor of 1.3
func CombSort(s []int, n int) []int {
	sf := 1.3
	for gap, swapped := n, true; swapped; {
		gap = int(float64(gap) / sf)
		if gap < 1 {
			gap = 1
		}
		swapped = false
		for i := 0; i+gap < n; i++ {
			if s[i] > s[i+gap] {
				swap(s, i, i+gap)
				swapped = true
			}
		}
	}
	return s
}
