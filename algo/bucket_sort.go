package algo

// BucketSort is a distribution sort. The concept is
// to distribute elements into buckets/bins.
// In this implementation, bucket is a slice of node pointers.
func BucketSort(s []int, n int) []int {
	max := getMax(s, n)
	bins := make([]*Node, max+1)
	for i := 0; i < max; i++ {
		bins[i] = nil
	}
	for j := 0; j < n; j++ {
		insert(bins, s, j)
	}
	for i, j := 0, 0; i < max+1; i++ {
		for bins[i] != nil {
			s[j] = bins[i].value
			bins[i] = bins[i].next
			j++
		}
	}
	return s
}

// Node represents element to store in bucket
type Node struct {
	value int
	next  *Node
}

// insert adds slice elements as nodes and store in buckets
//
func insert(bins []*Node, s []int, j int) {
	n := &Node{value: s[j]}
	if bins[s[j]] == nil {
		bins[s[j]] = n
	} else {
		bins[s[j]].next = n
	}
}

// getMax returns the maximum element in the slice
func getMax(s []int, n int) int {
	max := s[0]
	for i := 1; i < n; i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
