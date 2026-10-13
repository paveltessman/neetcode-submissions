func longestConsecutive(nums []int) int {

	bestFor := make(map[int]int)
	for _, val := range nums {
		bestFor[val] = 1
	}

	for key := range bestFor {
		if _, exists := bestFor[key-1]; exists {
			continue
		}
		next := key + 1
		for {
			if _, exists := bestFor[next]; !exists {
				break
			}
			bestFor[key]++
			next++
		}
	}

	best := 0
	for _, score := range bestFor {
		best = max(best, score)
	}
	return best


}
