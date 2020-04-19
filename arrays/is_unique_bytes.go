package arrays

import (
	"sort"
)

type isUniqueBytes func([]byte) bool

func IsUniqueBytesWithBoolArray(input []byte) bool {
	if len(input) > 128 {
		return false
	}

	visited := make([]bool, 128)

	for _, char := range input {
		if visited[int(char)] {
			return false
		}

		visited[int(char)] = true
	}

	return true
}

func IsUniqueBytesWithBitArray(input []byte) bool {
	if len(input) > 128 {
		return false
	}

	bitset := [2]int64{0, 0}

	for _, char := range input {
		bitsetItem := int(char) / 64
		bitsetOffset := int(char) % 64

		if (bitset[bitsetItem] & (1 << bitsetOffset)) > 0 {
			return false
		}

		bitset[bitsetItem] |= (1 << bitsetOffset)
	}

	return true
}

func IsUniqueBytesWithFullScan(input []byte) bool {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}

func IsUniqueBytesWithSorting(input []byte) bool {
	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })

	for i := 1; i < len(input); i++ {
		if input[i-1] == input[i] {
			return false
		}
	}

	return true
}
