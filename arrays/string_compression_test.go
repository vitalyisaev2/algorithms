package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type stringCompressionTestCase struct {
	input  string
	output string
}

func makeStringCompressionTestCases() []stringCompressionTestCase {
	return []stringCompressionTestCase{
		{
			input:  "aabcccccaaa",
			output: "a2b1c5a3",
		},
		{
			input:  "aabcccccaaad",
			output: "a2b1c5a3d1",
		},
		{
			input:  "aabcccccaaaddd",
			output: "a2b1c5a3d3",
		},
		{
			input:  "dddd",
			output: "d4",
		},
	}
}

func TestStringCompression(t *testing.T) {
	testCases := makeStringCompressionTestCases()

	callbacks := []stringCompression{
		StringCompression,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				output, err := cb(tc.input)
				assert.NoError(t, err)
				assert.Equal(t, tc.output, output)
			}
		})
	}
}
