package arrays

import (
	"strings"
	"testing"

	"gitlab.com/vitalyisaev2/algorithms/utils"

	"github.com/stretchr/testify/assert"
)

type isUniqueStringTestCase struct {
	input  string
	output bool
}

func makeIsUniqueStringTestCase() []isUniqueStringTestCase {
	// конструирование наиболее длинной уникальной строки из возможных
	buf := strings.Builder{}

	for i := int(0); i < 128; i++ {
		_ = buf.WriteByte(byte(i))
	}

	longestUniqueLine := buf.String()

	return []isUniqueStringTestCase{
		{
			input:  "a",
			output: true,
		},
		{
			input:  "aa",
			output: false,
		},
		{
			input:  "ab",
			output: true,
		},
		{
			input:  "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ",
			output: true,
		},
		{
			input:  utils.RandomASCIIString(129),
			output: false,
		},
		{
			input:  longestUniqueLine,
			output: true,
		},
	}
}

func TestIsUniqueString(t *testing.T) {
	testCases := makeIsUniqueStringTestCase()

	callbacks := []isUniqueString{
		IsUniqueStringWithBitArray,
		IsUniqueStringWithBoolArray,
		IsUniqueStringWithFullScan,
		IsUniqueStringWithSorting,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.output, cb(tc.input), "input: '%v'", tc.input)
			}
		})
	}
}

func BenchmarkIsUniqueString(b *testing.B) {
	testCases := makeIsUniqueStringTestCase()

	callbacks := []isUniqueString{
		IsUniqueStringWithBitArray,
		IsUniqueStringWithBoolArray,
		IsUniqueStringWithFullScan,
		IsUniqueStringWithSorting,
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
