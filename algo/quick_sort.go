package algo

func QuickSort(s []int) []int {
	var f func([]int, int, int)
	f = func(s []int, low, high int) {
		if low < high {
			p := partition(s, low, high)
			f(s, low, p)
			f(s, p+1, high)
		}
	}
	f(s, 0, len(s)-1)
	return s
}

// Partition splits array into two parts
func partition(s []int, low, high int) int {
	pivot := s[low]
	i, j := low-1, high+1

	for i < j {
		for {
			i += 1
			if s[i] >= pivot {
				break
			}
		}
		for {
			j -= 1
			if s[j] <= pivot {
				break
			}
		}
		if i < j {
			swap(s, i, j)
		}
	}

	swap(s, low, j)
	return j
}
