package arrays

import (
	"sort"
)

type isPermutation func(string, string) bool

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

func IsPermutationWithMapRuneCount(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// this makes them stay on stack
	runeCount1 := map[rune]int{}
	runeCount2 := map[rune]int{}

	for _, r := range s1 {
		runeCount1[r] += 1
	}

	for _, r := range s2 {
		runeCount2[r] += 1
	}

	for r, count := range runeCount1 {
		if runeCount2[r] != count {
			return false
		}
	}

	return true
}

func IsPermutationWithArrayRuneCount(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runeCount1 := [128]int{}
	runeCount2 := [128]int{}

	for _, r := range s1 {
		runeCount1[r] += 1
	}

	for _, r := range s2 {
		runeCount2[r] += 1
	}

	for ix, count := range runeCount1 {
		if runeCount2[ix] != count {
			return false
		}
	}

	return true
}

func IsPermutationWithArrayRuneCount2(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	runeCount := [128]int{}

	for _, r := range s1 {
		runeCount[r] += 1
	}

	for _, r := range s2 {
		runeCount[r] -= 1
		if runeCount[r] < 0 {
			return false
		}
	}

	return true
}
