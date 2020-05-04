package matching

import (
	"gitlab.com/vitalyisaev2/algorithms/utils"
)

func hash(input []rune) int32 {
	var sum int32 = 0

	// TODO: byte shift for faster polynomial computation
	for i, r := range input {
		sum += int32(r) * utils.PowerInt32(2, int32(len(input)-1-i))
	}

	return sum
}

func rehash(first, last rune, size int, sumPrev int32) int32 {
	pow := utils.PowerInt32(2, int32(size-1))
	sumNext := 2*(sumPrev-int32(first)*pow) + int32(last)

	return sumNext
}

func karpRabin(_pattern, _text string) bool {
	pattern := []rune(_pattern)
	text := []rune(_text)

	patternHash := hash(pattern)
	textHash := hash(text[:len(pattern)])

	// strings of equal size
	if len(text) == len(pattern) {
		return textHash == patternHash
	}

	edge := len(text) - len(pattern)

OUTER:
	for j := 0; j <= edge; j++ {
		// hashes matched, need to compare char by char
		if patternHash == textHash {
			for i := 0; i < len(pattern); i++ {
				// it was a collision
				if pattern[i] != text[j+i] {
					continue OUTER
				}
			}
			return true
		}

		if j+len(pattern) == len(text) {
			break OUTER
		}

		textHash = rehash(text[j], text[j+len(pattern)], len(pattern), textHash)
	}

	return false
}
