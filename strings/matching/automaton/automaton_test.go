package automaton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPrefixLength(t *testing.T) {
	assert.Equal(t, 1, longestPrefixLength("A", 'A'))
	assert.Equal(t, 1, longestPrefixLength("ACA", 'A'))
	assert.Equal(t, 1, longestPrefixLength("ACACA", 'A'))
	assert.Equal(t, 4, longestPrefixLength("ACACA", 'C'))
	assert.Equal(t, 1, longestPrefixLength("ACACAGA", 'A'))
	assert.Equal(t, 2, longestPrefixLength("ACACAGA", 'C'))
}
