package arrays

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
