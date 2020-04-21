package matching

type graph struct {
	vertexNumber  int
	edgeNumber    int
	vertexCounter int
	initial       int
	terminal      []int
	target        []int
	suffixLink    []int
	length        []int
	position      []int
	shift         []int
}

func (g *graph) vertex() int {
	g.vertexNumber++
	return g.vertexNumber
}

func newGraph(v, e int) *graph {
	return &graph{
		vertexNumber:  v,
		edgeNumber:    e,
		vertexCounter: 1,
		initial:       0,
	}
}

func newAutomaton(v, e int) *graph {
	g := newGraph(v, e)
	g.target = make([]int, 0, e)
	g.terminal = make([]int, 0, v)

	return g
}

const undefined = -1

func newSuffixAutomaton(v, e int) *graph {
	g := newAutomaton(v, e)

	// no memset in Go
	for i := range g.target {
		g.target[i] = undefined
	}

	g.suffixLink = make([]int, 0, v)
	g.length = make([]int, 0, v)
	g.position = make([]int, 0, v)
	g.shift = make([]int, 0, e)

	return g
}

func newTrie(v, e int) *graph {
	g := newAutomaton(v, e)

	for i := range g.target {
		g.target[i] = undefined
	}

	g.suffixLink = make([]int, 0, v)
	g.length = make([]int, 0, v)
	g.position = make([]int, 0, v)
	g.shift = make([]int, 0, e)

	return g
}
