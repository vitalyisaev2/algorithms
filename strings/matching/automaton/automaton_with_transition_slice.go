package automaton

import (
	"sort"
)

var _ Automaton = (*automatonWithTransitionSlice)(nil)

// ab is a sorted slice of runes
type alphabet []rune

func (ab alphabet) runePosition(r rune) (position int, exists bool) {
	position = sort.Search(
		len(ab),
		func(i int) bool { return ab[i] >= r },
	)

	exists = position < len(ab)

	return
}

type automatonWithTransitionSlice struct {
	currentState automatonState
	acceptStates []bool
	ab           alphabet           // sorted slice of symbols
	transitions  [][]automatonState // outer - slice of states, inner - slice of transitions
}

func (a *automatonWithTransitionSlice) Match(input string) bool {
	// reset state
	a.currentState = 0

	for _, r := range input {
		runePositionInAlphabet, exists := a.ab.runePosition(r)

		if !exists {
			// no such rune in ab
			a.currentState = 0
			continue
		}

		nextState := a.transitions[a.currentState][runePositionInAlphabet]

		if a.acceptStates[nextState] {
			return true
		}

		a.currentState = nextState
	}

	return false
}

func NewAutomatonWithTransitionSlice(pattern string) Automaton {
	// first of all collect unique runes, than sort them
	uniqueRunes := map[rune]struct{}{}
	for _, r := range pattern {
		uniqueRunes[r] = struct{}{}
	}

	// sorted slice of runes
	ab := make(alphabet, 0, len(uniqueRunes))
	for r := range uniqueRunes {
		ab = append(ab, r)
	}

	sort.Slice(ab, func(i, j int) bool {
		return ab[i] < ab[j]
	})

	// construct transitions
	i := automatonState(0)
	transitions := make([][]automatonState, len(pattern)+1)

	for _, r := range pattern {
		transitions[i] = make([]automatonState, len(ab))

		runePositionInAlphabet, _ := ab.runePosition(r)

		transitions[i][runePositionInAlphabet] = i + 1
		i++
	}

	// set last state as successful
	acceptStates := make([]bool, len(pattern)+1)
	acceptStates[i] = true

	// take into account all alphabet characters and build additional edges
	for i := 1; i < len(transitions)-1; i++ {
		for r := range uniqueRunes {
			// omit existing state transitions
			runePositionInAlphabet, _ := ab.runePosition(r)

			if transitions[i][runePositionInAlphabet] != 0 {
				continue
			}

			if length := longestPrefixLength(pattern[:i], r); length != 0 {
				transitions[i][runePositionInAlphabet] = length
			}
		}
	}

	return &automatonWithTransitionSlice{
		currentState: 0,
		ab:           ab,
		acceptStates: acceptStates,
		transitions:  transitions,
	}
}
