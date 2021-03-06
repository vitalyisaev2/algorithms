package automaton

import (
	"unicode/utf8"
)

var _ Automaton = (*automatonWithTransitionMap)(nil)

type automatonWithTransitionMap struct {
	currentState automatonState
	acceptStates []bool
	transitions  []map[rune]automatonState
}

func (a *automatonWithTransitionMap) Match(input string) bool {
	// reset state
	a.currentState = 0

	for _, r := range input {
		possibleTransitions := a.transitions[a.currentState]

		nextState, exists := possibleTransitions[r]
		if !exists {
			a.currentState = 0
			continue
		}

		if a.acceptStates[nextState] {
			return true
		}

		a.currentState = nextState
	}

	return false
}

func NewAutomatonWithTransitionMap(pattern string) Automaton {
	// build main state transitions line, also collect set of unique patten runes (ab)
	i := automatonState(0)
	alphabet := map[rune]struct{}{}
	transitions := make([]map[rune]automatonState, utf8.RuneCountInString(pattern)+1)

	for _, r := range pattern {
		transitions[i] = map[rune]automatonState{r: i + 1}
		alphabet[r] = struct{}{}
		i++
	}

	// set last state as successful
	acceptStates := make([]bool, len(pattern)+1)
	acceptStates[i] = true

	// take into account all ab characters and build additional edges
	for i := 1; i < len(transitions)-1; i++ {
		for r := range alphabet {
			// omit existing state transitions
			if _, exists := transitions[i][r]; exists {
				continue
			}

			if length := longestPrefixLength(pattern[:i], r); length != 0 {
				transitions[i][r] = length
			}
		}
	}

	return &automatonWithTransitionMap{
		currentState: 0,
		acceptStates: acceptStates,
		transitions:  transitions,
	}
}
