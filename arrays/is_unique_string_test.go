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

func makeIsUniqueStringTestCases() []isUniqueStringTestCase {
	// конструирование наиболее длинной уникальной строки из возможных
	buf := strings.Builder{}

	for i := int(0); i < 256; i++ {
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
			input:  utils.RandomString(257),
			output: false,
		},
		{
			input:  longestUniqueLine,
			output: true,
		},
	}
}

func TestIsUniqueString(t *testing.T) {
	testCases := makeIsUniqueStringTestCases()

	callbacks := []isUniqueString{
		isUniqueStringWithBitArray,
		isUniqueStringWithBoolArray,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.output, cb(tc.input))
			}
		})
	}
}

func BenchmarkIsUniqueString(b *testing.B) {
	testCases := makeIsUniqueStringTestCases()

	callbacks := []isUniqueString{
		isUniqueStringWithBitArray,
		isUniqueStringWithBoolArray,
	}

	for _, cb := range callbacks {
		b.Run(utils.FunctionName(cb), func(b *testing.B) {
			b.ReportAllocs()
			b.ResetTimer()
			for _, tc := range testCases {
				_ = cb(tc.input)
			}
		})
	}
}
