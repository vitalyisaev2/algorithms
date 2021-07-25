package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	type testCase struct {
		input       []int
		outputSlice []int
		outputValue int
	}

	testCases := []testCase{
		// {
		// 	input:       []int{1, 1, 2},
		// 	outputSlice: []int{1, 2},
		// 	outputValue: 2,
		// },
		// {
		// 	input:       []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
		// 	outputSlice: []int{0, 1, 2, 3, 4},
		// 	outputValue: 5,
		// },
		{
			input:       []int{1, 1, 1, 2},
			outputSlice: []int{1, 2},
			outputValue: 2,
		},
	}

	for _, tc := range testCases {
		outputValue := RemoveDuplicatesFromSortedArray(tc.input)
		require.Equal(t, tc.outputValue, outputValue)
		require.Equal(t, tc.outputSlice, tc.input[:tc.outputValue])
	}
}
