package automaton

type automatonState uint16

type Automaton struct {
	currentState automatonState
	acceptStates []bool
	transitions  []map[rune]automatonState
}

func (a *Automaton) Match(input string) bool {
	for _, r := range input {
		possibleTransitions := a.transitions[a.currentState]

		nextState, exists := possibleTransitions[r]
		if !exists {
			return false
		}

		if a.acceptStates[nextState] {
			return true
		}

		a.currentState = nextState
	}

	return false
}

func (a *Automaton) Reset() { a.currentState = 0 }

func longestPrefixLength(line string, r rune) int {
	target := line + string(r)
	i := len(line)

	for ; i > 0; i-- {
		prefix := line[:i]
		j := len(target) - i
		suffix := target[j:]

		if prefix == suffix {
			break
		}
	}

	return i
}

func NewAutomaton(pattern string) *Automaton {
	// build main state transitions line, also collect set of unique patten runes (alphabet)
	i := automatonState(0)
	alphabet := map[rune]struct{}{}
	transitions := make([]map[rune]automatonState, len(pattern)+1)

	for _, r := range pattern {
		transitions[i] = map[rune]automatonState{r: i + 1}
		alphabet[r] = struct{}{}
		i++
	}

	// set last state as successful
	acceptStates := make([]bool, len(pattern))
	acceptStates[i] = true

	// take into account all alphabet characters and build additional edges

	return &Automaton{
		currentState: 0,
		acceptStates: acceptStates,
		transitions:  transitions,
	}
}
