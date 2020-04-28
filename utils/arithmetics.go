package utils

func PowerInt32(base, exponent int32) int32 {
	var result int32 = 1

	for exponent > 1 {
		if exponent%2 == 1 {
			result *= base
		}

		base *= base
		exponent /= 2
	}

	if exponent > 0 {
		result *= base
	}

	return result
}

func PowerInt64(base, exponent int64) int64 {
	var result int64 = 1

	for exponent > 1 {
		if exponent%2 == 1 {
			result *= base
		}

		base *= base
		exponent /= 2
	}

	if exponent > 0 {
		result *= base
	}

	return result
}
