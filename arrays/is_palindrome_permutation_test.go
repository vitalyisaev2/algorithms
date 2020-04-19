package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type isPalindromePermutationTestCase struct {
	input  string
	output bool
}

func makeIsPalindromePermutationTestCases() []isPalindromePermutationTestCase {
	result := []isPalindromePermutationTestCase{
		{
			"a",
			true,
		},
		{
			"ab",
			false,
		},
		{
			"aba",
			true,
		},
	}

	for _, size := range []int{1, 32, 1024, 32768} {
		{
			s1 := utils.RandomASCIIString(size)
			s2 := s1 + utils.ReverseString(s1)
			result = append(result, isPalindromePermutationTestCase{
				input:  s2,
				output: true,
			})
		}

		{
			s1 := utils.RandomASCIIString(size)
			s2 := s1 + "x" + utils.ReverseString(s1)
			result = append(result, isPalindromePermutationTestCase{
				input:  s2,
				output: true,
			})
		}

		result = append(result, isPalindromePermutationTestCase{
			input:  utils.RandomASCIIString(size),
			output: size == 1,
		})
	}

	return result
}

func TestIsPalindromePermutation(t *testing.T) {
	testCases := makeIsPalindromePermutationTestCases()

	callbacks := []isPalindromePermutation{
		IsPalindromePermutationWithCounter,
		IsPalindromePermutationWithCounterOneWay,
		IsPalindromePermutationWithBitArray,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.output, cb(tc.input), "input: '%v'", tc.input)
			}
		})
	}
}

func BenchmarkIsPalindromePermutation(b *testing.B) {
	testCases := makeIsPalindromePermutationTestCases()

	callbacks := []isPalindromePermutation{
		IsPalindromePermutationWithCounter,
		IsPalindromePermutationWithCounterOneWay,
		IsPalindromePermutationWithBitArray,
	}

	for _, cb := range callbacks {
		b.Run(utils.FunctionName(cb), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tc := range testCases {
					output := cb(tc.input)
					if output != tc.output {
						b.FailNow()
					}
				}
			}
		})
	}
}
