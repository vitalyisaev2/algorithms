package matching

import (
	"gitlab.com/vitalyisaev2/algorithms/strings/matching/automaton"
)

// Space complexity: O(m * sigma?)
// Time complexity (preprocessing): O(m * sigma?)
// Time complexity (search): O(n)
//nolint:deadcode
func automatonWithTransitionMap(pattern, text string) bool {
	aut := automaton.NewAutomatonWithTransitionMap(pattern)
	return aut.Match(text)
}

// Space complexity: O(m * sigma?)
// Time complexity (preprocessing): O(m * sigma?)
// Time complexity (search): O(n)
//nolint:deadcode
func automatonWithTransitionSlice(pattern, text string) bool {
	aut := automaton.NewAutomatonWithTransitionSlice(pattern)
	return aut.Match(text)
}
