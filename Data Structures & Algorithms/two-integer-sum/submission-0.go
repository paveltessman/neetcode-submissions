func twoSum(nums []int, target int) []int {
    values := make(map[int]int)

	for i, val := range nums {
		values[val] = i
	}

	for i, first := range nums {
		second := target - first
		if j, exists := values[second]; exists && i != j {
			if i < j {
				return []int{i, j}
			}
			return []int{j, i}
		}
	}
	panic("not reachable")
}