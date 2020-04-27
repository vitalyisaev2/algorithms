package automaton

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
