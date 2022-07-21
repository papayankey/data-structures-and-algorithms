package algo

func CountSort(s []int, n int) []int {
	max := GetMax(s, n)
	count := make([]int, max+1)

	for i := 0; i < n; i++ {
		count[s[i]] += 1
	}

	for j, k := 0, 0; j < max+1; j++ {
		for count[j] != 0 {
			s[k] = j
			count[j] -= 1
			k += 1
		}
	}

	return s
}
