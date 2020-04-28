package matching

import (
	"unicode/utf8"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

func hash(input string) int32 {
	var (
		size       = utf8.RuneCountInString(input)
		sum  int32 = 0
	)

	// TODO: byte shift for faster polynomial computation
	for i, r := range input {
		sum += int32(r) * utils.PowerInt32(2, int32(size-1-i))
	}

	return sum
}

func rehash(first, last rune, size int, sum int32) int32 {
	pow := utils.PowerInt32(2, int32(size-1))
	return 2*(sum-int32(first)*pow) + int32(last)
}

func karpRabin(pattern, text string) bool {
	patternHash := hash(pattern)
	textHash := hash(text[:len(pattern)])

	patternSize := utf8.RuneCountInString(pattern)
	textSize := utf8.RuneCountInString(text)

OUTER:
	for j := 0; j <= textSize-patternSize; j++ {
		// hashes matched, need to compare char by char
		if patternHash == textHash {
			for i := 0; i < patternSize; i++ {
				// it was a collision
				if pattern[i] != text[j+i] {
					continue OUTER
				}
			}
			return true
		}

		textHash = rehash(rune(text[j]), rune(text[j+patternSize]), patternSize, textHash)
	}

	return false
}
