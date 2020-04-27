package automaton

type automatonState uint16

type Automaton interface {
	Match(input string) bool
}
