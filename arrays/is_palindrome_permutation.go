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
