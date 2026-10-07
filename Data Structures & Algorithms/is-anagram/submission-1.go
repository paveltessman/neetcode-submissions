func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mask := make([]int, 26)

	for _, char := range s {
		mask[int(char) - 'a'] += 1
	}

	for _, char := range t {
		mask[int(char) - 'a'] -= 1
	}

	for _, val := range mask {
		if val != 0 {
			return false
		}
	}
	return true
}
