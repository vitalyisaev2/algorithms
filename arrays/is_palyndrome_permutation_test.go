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

	{
		s1 := utils.RandomASCIIString(128)
		s2 := s1 + utils.ReverseString(s1)
		result = append(result, isPalindromePermutationTestCase{
			input:  s2,
			output: true,
		})
	}

	{
		s1 := utils.RandomASCIIString(128)
		s2 := s1 + "x" + utils.ReverseString(s1)
		result = append(result, isPalindromePermutationTestCase{
			input:  s2,
			output: true,
		})
	}

	result = append(result, isPalindromePermutationTestCase{
		input:  utils.RandomASCIIString(256),
		output: false,
	})

	return result
}

func TestIsPalindromePermutation(t *testing.T) {
	testCases := makeIsPalindromePermutationTestCases()

	callbacks := []isPalindromePermutation{
		IsPalindromePermutationWithCounter,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.output, cb(tc.input), "input: '%v'", tc.input)
			}
		})
	}
}
