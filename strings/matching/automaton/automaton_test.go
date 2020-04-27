package automaton

import (
	"strings"
	"testing"

	"gitlab.com/vitalyisaev2/algorithms/utils"

	"github.com/stretchr/testify/assert"
)

func TestLongestPrefixLength(t *testing.T) {
	t.Run("positive", func(t *testing.T) {
		assert.Equal(t, automatonState(1), longestPrefixLength("A", 'A'))
		assert.Equal(t, automatonState(1), longestPrefixLength("ACA", 'A'))
		assert.Equal(t, automatonState(1), longestPrefixLength("ACACA", 'A'))
		assert.Equal(t, automatonState(4), longestPrefixLength("ACACA", 'C'))
		assert.Equal(t, automatonState(1), longestPrefixLength("ACACAGA", 'A'))
		assert.Equal(t, automatonState(2), longestPrefixLength("ACACAGA", 'C'))
	})

	t.Run("negative", func(t *testing.T) {
		assert.Equal(t, automatonState(0), longestPrefixLength("A", 'C'))
		assert.Equal(t, automatonState(0), longestPrefixLength("ACACA", 'G'))
	})
}

func TestAutomatonWithTransitionMap(t *testing.T) {
	automaton := NewAutomatonWithTransitionMap("ACACAGA").(*automatonWithTransitionMap)

	t.Run("current state", func(t *testing.T) {
		assert.Equal(t, automatonState(0), automaton.currentState)
	})

	t.Run("accept state", func(t *testing.T) {
		acceptStates := make([]bool, 8)
		for i := 0; i < len(acceptStates)-1; i++ {
			acceptStates[i] = false
		}
		acceptStates[7] = true

		assert.Equal(t, acceptStates, automaton.acceptStates)
	})

	t.Run("transition states", func(t *testing.T) {
		assert.Len(t, automaton.transitions, 8)

		transitions := []map[rune]automatonState{
			{
				'A': 1,
			},
			{
				'A': 1,
				'C': 2,
			},
			{
				'A': 3,
			},
			{
				'A': 1,
				'C': 4,
			},
			{
				'A': 5,
			},
			{
				'A': 1,
				'C': 4,
				'G': 6,
			},
			{
				'A': 7,
			},
			nil,
		}

		assert.Equal(t, transitions, automaton.transitions)
	})
}

type constructor func(string) Automaton

func makeConstructors() []constructor {
	return []constructor{
		NewAutomatonWithTransitionMap,
		NewAutomatonWithTransitionSlice,
	}
}

func TestAutomatonMatch(t *testing.T) {
	for _, cn := range makeConstructors() {
		functionName := strings.Trim(utils.FunctionName(cn), "New")

		t.Run(functionName, func(t *testing.T) {
			testCases := []struct {
				pattern string
				text    string
				result  bool
			}{
				{
					pattern: "ACACAGA",
					text:    "ACACAGAC",
					result:  true,
				},
				{
					pattern: "ACACAGA",
					text:    "CRAB",
					result:  false,
				},
				{
					pattern: "GCAGAGAG",
					text:    "GCATCGCAGAGAGTATACAGTACG",
					result:  true,
				},
			}

			for _, tc := range testCases {
				aut := cn(tc.pattern)
				assert.Equal(t, tc.result, aut.Match(tc.text))
			}
		})
	}
}
