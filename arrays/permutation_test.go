package arrays

import (
	"testing"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type isPermutationTestCase struct {
	s1     string
	s2     string
	result bool
}

func makeIsPermutationTestCase() []isPermutationTestCase {
	var rcv []isPermutationTestCase

	// permutated ASCII strings
	{
		s1 := utils.RandomASCIIString(256)

		s2 := s1[1:] + string(s1[0])

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: true})
	}

	// different ASCII strings
	{
		s1 := utils.RandomASCIIString(256)
		s2 := utils.RandomASCIIString()

		rcv = append(rcv, isPermutationTestCase{s1: s1, s2: s2, result: true})
	}
}

func TestIsPermutation(t *testing.T) {

}
