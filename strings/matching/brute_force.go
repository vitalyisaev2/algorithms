package matching

// Time complexity: O(mn)
// Space complexity: O(1)
func bruteForce(pattern, text string) bool {
	limit := len(text) - len(pattern)

	for start := 0; start < limit; start++ {
		i := 0

		for {
			if !(i < len(pattern) && pattern[i] == text[start+i]) {
				break
			}
			i++
		}

		if i >= len(pattern) {
			return true
		}
	}

	return false
}
