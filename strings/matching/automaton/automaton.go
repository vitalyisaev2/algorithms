package automaton

type automatonState uint16

type Automaton struct {
	currentState automatonState
	acceptStates []bool
	transitions  []map[rune]automatonState
}

func (a *Automaton) Match(input string) bool {
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

func longestPrefixLength(line string, r rune) automatonState {
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

	return automatonState(i)
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
	acceptStates := make([]bool, len(pattern)+1)
	acceptStates[i] = true

	// take into account all alphabet characters and build additional edges
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

	return &Automaton{
		currentState: 0,
		acceptStates: acceptStates,
		transitions:  transitions,
	}
}
