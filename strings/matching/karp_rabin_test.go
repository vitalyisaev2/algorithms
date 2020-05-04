package matching

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKarpRabinHash(t *testing.T) {
	input := []rune("AБВГДЕЙКА")
	hashOutput := hash(input)
	assert.Equal(t, int32(282352), hashOutput)
	rehashOutput := rehash(('A'), ('Я'), len(input), hashOutput)
	assert.Equal(t, int32(532495), rehashOutput)

	// compare rehash and hash from scratch
	assert.Equal(t, hash([]rune("БВГДЕЙКАЯ")), rehashOutput)
}
