package matching

import (
	"gitlab.com/vitalyisaev2/algorithms/strings/matching/automaton"
)

func withAutomaton(pattern, text string) bool {
	aut := automaton.NewAutomaton(pattern)
	return aut.Match(text)
}
