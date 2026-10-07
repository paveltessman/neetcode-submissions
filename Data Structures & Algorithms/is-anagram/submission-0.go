func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	freqS := make(map[rune]int)
	freqT := make(map[rune]int)

	for _, char := range s {
		freqS[char] += 1
	}

	for _, char := range t {
		freqT[char] += 1
	}

	for char, countS := range freqS {
		if countT, exists := freqT[char]; !exists || countS != countT {
			return false
		}
	}
	return true
}
