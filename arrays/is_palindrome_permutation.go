package arrays

type isPalindromePermutation func(string) bool

func IsPalindromePermutationWithCounter(input string) bool {
	counter := [128]int{}

	for _, r := range input {
		counter[r] += 1
	}

	oddFound := false

	for _, count := range counter {
		if count%2 == 1 {
			// odd count found
			if oddFound {
				// no more than single odd count for palyndrome
				return false
			}

			oddFound = true
		}
	}

	return true
}

func IsPalindromePermutationWithCounterOneWay(input string) bool {
	counter := [128]int{}

	oddNumber := 0

	for _, r := range input {
		counter[r] += 1

		if counter[r]%2 == 0 {
			oddNumber--
		} else {
			oddNumber++
		}
	}

	return oddNumber == 0 || oddNumber == 1
}

func IsPalindromePermutationWithBitArray(input string) bool {
	counter := [2]int64{}

	for _, r := range input {
		ix := int(r) / 64
		offset := int(r) % 64

		counter[ix] ^= (1 << offset)
	}

	if counter[0] == 0 && counter[1] == 0 {
		// all odd
		return true
	}

	onlyOneOdd := false

	for _, c := range counter {
		if c == 0 {
			continue
		}

		if (c & (c - 1)) == 0 {
			if onlyOneOdd {
				return false
			}

			onlyOneOdd = true
		}
	}

	return onlyOneOdd
}
