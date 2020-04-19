package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	assert.Equal(t, "a", ReverseString("a"))
	assert.Equal(t, "aaaa", ReverseString("aaaa"))
	assert.Equal(t, "ba", ReverseString("ab"))
	assert.Equal(t, "barc", ReverseString("crab"))
}
