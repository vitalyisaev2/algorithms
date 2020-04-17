package arrays

import (
	"fmt"
)

type isUniqueString func(string) bool

func isUniqueStringWithBoolArray(input string) bool {
	if len(input) > 256 {
		return false
	}

	visited := make([]bool, 256)

	for _, char := range input {
		if visited[int(char)] {
			return false
		}

		visited[int(char)] = true
	}

	return true
}

func isUniqueStringWithBitArray(input string) bool {
	if len(input) > 256 {
		return false
	}

	bitset := [4]int64{0, 0, 0, 0}

	fmt.Println(input)
	for _, char := range input {
		bitsetItem := int(char) / 64
		bitsetOffset := int(char) % 64
		fmt.Println(bitsetItem, bitsetOffset)
		if (bitset[bitsetItem] & (1 << bitsetOffset)) > 0 {
			return false
		}

		bitset[bitsetItem] |= (1 << bitsetOffset)
	}

	return true
}
