package algo

// BucketSort is a distribution sort. The concept is
// to distribute elements into buckets or bins.
// This implementation stores node pointers as buckets.
func BucketSort(s []int, n int) []int {
	max := GetMax(s, n)
	bins := make([]*Node, max+1)

	for i := 0; i < n; i++ {
		insert(bins, s[i])
	}

	for i, j := 0, 0; i < max+1; i++ {
		for bins[i] != nil {
			s[j] = Remove(bins, i)
			j++
		}
	}

	return s
}

// Node represents bucket or bin
type Node struct {
	value int
	next  *Node
}

// remove deletes bucket and returns bucket value
// Bucket index is replaced with nil or next bucket
func Remove(b []*Node, i int) int {
	val := b[i].value
	b[i] = b[i].next
	return val
}

// insert adds slice element as bucket in buckets
// Each bucket index corresponds to slice element value
func insert(bins []*Node, i int) {
	n := &Node{value: i}
	if bins[i] == nil {
		bins[i] = n
	} else {
		curr := bins[i]
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = n
	}
}

// GetMax returns the maximum element in the slice
func GetMax(s []int, n int) int {
	max := s[0]
	for i := 1; i < n; i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
