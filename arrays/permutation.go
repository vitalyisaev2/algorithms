package arrays

import (
	"sort"
)

func isPermutation(string, string) bool

func sortString(input string) string {
	rcv := []rune(input)
	sort.Slice(rcv, func(i, j int) bool { return rcv[i] < rcv[j] })

	return string(rcv)
}

func IsPermutationWithSorting(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1 = sortString(s1)
	s2 = sortString(s2)

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

func makeRuneFrequencyMap(input string) map[rune]int {
	rcv := make(map[rune]int)

	for _, r := range input {
		rcv[r] += 1
	}

	return rcv
}

func IsPermutationWithRuneCount(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runeCount1 := makeRuneFrequencyMap(s1)
	runeCount2 := makeRuneFrequencyMap(s2)

	for r, count := range runeCount1 {
		if runeCount2[r] != count {
			return false
		}
	}

	return false
}
