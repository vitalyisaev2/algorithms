package matching

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

func TestMatching(t *testing.T) {
	callbacks := []matcher{
		bruteForce,
	}

	pattern := "GCAGAGAG"
	text := "GCATCGCAGAGAGTATACAGTACG"

	for _, cb := range callbacks {
		t.Run(utils.FunctionName(cb), func(t *testing.T) {
			assert.Equal(t, true, cb(pattern, text))
		})
	}
}
