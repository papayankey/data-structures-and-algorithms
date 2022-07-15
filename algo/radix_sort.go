package algo

// RadixSort is also a distribution sort with similar
// concept as bucket sort. The difference is that,
// radix sort reduces the memory allocation for bins.
func RadixSort(s []int, radix int) []int {
	max := GetMax(s, len(s))
	bins := make([]*Node, radix)

	for divisor := 1; divisor < max; divisor *= radix {
		for i := 0; i < len(s); i++ {
			loc := (s[i] / divisor) % radix
			insertToBucket(bins, s[i], loc)
		}

		for i, j := 0, 0; i < radix; i++ {
			for bins[i] != nil {
				s[j] = Remove(bins, i)
				j++
			}
		}
	}

	return s
}

// insertToBucket adds slice element as bucket
func insertToBucket(bins []*Node, val, loc int) {
	n := &Node{value: val}
	if bins[loc] == nil {
		bins[loc] = n
	} else {
		curr := bins[loc]
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}
