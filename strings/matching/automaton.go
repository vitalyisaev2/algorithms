package matching

import (
	"gitlab.com/vitalyisaev2/algorithms/strings/matching/automaton"
)

//nolint:deadcode
func withAutomaton(pattern, text string) bool {
	aut := automaton.NewAutomaton(pattern)
	return aut.Match(text)
}
