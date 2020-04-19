package arrays

import (
	"strings"
	"testing"

	"gitlab.com/vitalyisaev2/algorithms/utils"

	"github.com/stretchr/testify/assert"
)

type isUniqueBytesTestCase struct {
	input  []byte
	output bool
}

func (tc isUniqueBytesTestCase) cloneInput() []byte {
	clone := make([]byte, len(tc.input))
	copy(clone, tc.input)

	return clone
}

func makeIsUniqueBytesTestCase() []isUniqueBytesTestCase {
	// конструирование наиболее длинной уникальной строки из возможных
	buf := strings.Builder{}

	for i := int(0); i < 128; i++ {
		_ = buf.WriteByte(byte(i))
	}

	longestUniqueLine := buf.String()

	return []isUniqueBytesTestCase{
		{
			input:  []byte("a"),
			output: true,
		},
		{
			input:  []byte("aa"),
			output: false,
		},
		{
			input:  []byte("ab"),
			output: true,
		},
		{
			input:  []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"),
			output: true,
		},
		{
			input:  utils.RandomBytes(129),
			output: false,
		},
		{
			input:  []byte(longestUniqueLine),
			output: true,
		},
	}
}

func TestIsUniqueBytes(t *testing.T) {
	testCases := makeIsUniqueBytesTestCase()

	callbacks := []isUniqueBytes{
		IsUniqueBytesWithBitArray,
		IsUniqueBytesWithBoolArray,
		IsUniqueBytesWithFullScan,
		IsUniqueBytesWithSorting,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.output, cb(tc.cloneInput()), "input: '%v'", tc.input)
			}
		})
	}
}

func BenchmarkIsUniqueString(b *testing.B) {
	testCases := makeIsUniqueBytesTestCase()

	callbacks := []isUniqueBytes{
		IsUniqueBytesWithBitArray,
		IsUniqueBytesWithBoolArray,
		IsUniqueBytesWithFullScan,
		IsUniqueBytesWithSorting,
	}

	for m, cb := range callbacks {
		b.Run(utils.FunctionName(cb), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tc := range testCases {
					input := tc.input

					// функция IsUniqueBytesWithSorting изменяет входящий массив
					if m == 3 {
						b.StopTimer()
						input = tc.cloneInput()
						b.StartTimer()
					}

					output := cb(input)
					if output != tc.output {
						b.FailNow()
					}
				}
			}
		})
	}
}
