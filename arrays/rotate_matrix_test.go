package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateMatrix(t *testing.T) {
	t.Run("1x1", func(t *testing.T) {
		input := Matrix{{1}}
		output := Matrix{{1}}
		RotateMatrix(input)
		assert.Equal(t, output, input)
	})

	t.Run("2x2", func(t *testing.T) {
		input := Matrix{
			{1, 2},
			{3, 4},
		}
		output := Matrix{
			{3, 1},
			{4, 2},
		}
		RotateMatrix(input)
		assert.Equal(t, output, input)
	})

	t.Run("3x3", func(t *testing.T) {
		input := Matrix{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}
		output := Matrix{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}
		RotateMatrix(input)
		assert.Equal(t, output, input)
	})

	t.Run("4x4", func(t *testing.T) {
		input := Matrix{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}
		output := Matrix{
			{13, 9, 5, 1},
			{14, 10, 6, 2},
			{15, 11, 7, 3},
			{16, 12, 8, 4},
		}
		RotateMatrix(input)
		assert.Equal(t, output, input)
	})

	t.Run("5x5", func(t *testing.T) {
		input := Matrix{
			{1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15},
			{16, 17, 18, 19, 20},
			{21, 22, 23, 24, 25},
		}
		output := Matrix{
			{21, 16, 11, 6, 1},
			{22, 17, 12, 7, 2},
			{23, 18, 13, 8, 3},
			{24, 19, 14, 9, 4},
			{25, 20, 15, 10, 5},
		}
		RotateMatrix(input)
		assert.Equal(t, output, input)
	})
}
