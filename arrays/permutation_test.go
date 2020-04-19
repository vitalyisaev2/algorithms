package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type isPermutationTestCase struct {
	s1     string
	s2     string
	result bool
}

func makeIsPermutationTestCase() []isPermutationTestCase {
	var rcv []isPermutationTestCase

	// shuffled ASCII strings
	{
		s1 := utils.RandomASCIIString(256)
		s2 := utils.ShuffleString(s1)

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: true})
	}

	// different ASCII strings
	{
		s1 := utils.RandomASCIIString(256)
		s2 := utils.RandomASCIIString(256)

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: false})
	}

	// shuffled Cyrillic strings
	{
		s1 := utils.RandomUnicodeString(256, utils.CyrillicRussian)
		s2 := utils.ShuffleString(s1)

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: true})
	}

	// different Cyrillic strings
	{
		s1 := utils.RandomUnicodeString(256, utils.CyrillicRussian)
		s2 := utils.RandomUnicodeString(256, utils.CyrillicRussian)

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: false})
	}

	return rcv
}

func TestIsPermutation(t *testing.T) {
	callbacks := []isPermutation{
		IsPermutationWithSorting,
		IsPermutationWithRuneCount,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range makeIsPermutationTestCase() {
				assert.Equal(t, tc.result, cb(tc.s1, tc.s2))
			}
		})
	}
}

func BenchmarkIsPermutation(b *testing.B) {
	testCases := makeIsPermutationTestCase()

	callbacks := []isPermutation{
		IsPermutationWithSorting,
		IsPermutationWithRuneCount,
	}

	for _, cb := range callbacks {
		b.Run(utils.FunctionName(cb), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tc := range testCases {
					output := cb(tc.s1, tc.s2)
					if output != tc.result {
						b.FailNow()
					}
				}
			}
		})
	}
}
