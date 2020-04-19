package arrays

type oneAway func(string, string) bool

func stringLenDiff(s1, s2 string) int {
	diff := len(s1) - len(s2)
	if diff < 0 {
		return -1 * diff
	}

	return diff
}

func OneAway(s1, s2 string) bool {
	if stringLenDiff(s1, s2) > 1 {
		return false
	}

	edited := false

	ix1 := 0
	ix2 := 0

	for {
		// loop exit condition
		if ix1 == len(s1) || ix2 == len(s2) {
			break
		}

		if s1[ix1] != s2[ix2] {
			// more than one edition discovered
			if edited {
				return false
			}

			edited = true

			if s1[ix1] == s2[ix2+1] {
				// insertion
				ix2++
			} else if s1[ix1+1] == s2[ix2] {
				// removal
				ix1++
			} else if s1[ix1+1] == s2[ix2+1] {
				// replacement
				ix1++
				ix2++
			}
		} else {
			ix1++
			ix2++
		}
	}

	return true
}
