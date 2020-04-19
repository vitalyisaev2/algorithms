package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type oneAwayTestCase struct {
	s1     string
	s2     string
	result bool
}

func makeOneAwayTestCase() []oneAwayTestCase {
	result := []oneAwayTestCase{
		{
			s1:     "pale",
			s2:     "ple",
			result: true,
		},
		{
			s1:     "pales",
			s2:     "pale",
			result: true,
		},
		{
			s1:     "pale",
			s2:     "bale",
			result: true,
		},
		{
			s1:     "pale",
			s2:     "bae",
			result: false,
		},
		{
			s1:     "apple",
			s2:     "aple",
			result: true,
		},
	}

	return result
}

func TestOneAway(t *testing.T) {
	testCases := makeOneAwayTestCase()

	callbacks := []oneAway{
		OneAway,
	}

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			for _, tc := range testCases {
				assert.Equal(t, tc.result, cb(tc.s1, tc.s2), "input: '%v', '%v'", tc.s1, tc.s2)
			}
		})
	}
}
