package matching

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"gitlab.com/vitalyisaev2/algorithms/strings/matching/automaton"

	"gitlab.com/vitalyisaev2/algorithms/utils"
)

type matcher interface {
	prepare(pattern, text string) // preparation phase
	search() bool                 // search phase
	fmt.Stringer
}

type callbackMatcher struct {
	pattern string
	text    string
	matcher func(pattern, text string) bool
}

func (tc *callbackMatcher) prepare(pattern, text string) {
	tc.pattern, tc.text = pattern, text
}

func (tc *callbackMatcher) search() bool {
	return tc.matcher(tc.pattern, tc.text)
}

func (tc *callbackMatcher) String() string {
	name := utils.FunctionName(tc.matcher)
	parts := strings.Split(name, ".")

	return parts[len(parts)-1]
}

type automatonMatcher struct {
	text        string
	automaton   automaton.Automaton
	constructor func(string) automaton.Automaton
}

func (tc *automatonMatcher) prepare(pattern, text string) {
	tc.text = text
	tc.automaton = tc.constructor(pattern)
}

func (tc *automatonMatcher) search() bool {
	return tc.automaton.Match(tc.text)
}

func (tc *automatonMatcher) String() string {
	byPoints := strings.Split(utils.FunctionName(tc.constructor), ".")
	return strings.Trim(byPoints[len(byPoints)-1], "New")
}

func makeMatchers() []matcher {
	return []matcher{
		&callbackMatcher{matcher: bruteForce},
		&automatonMatcher{constructor: automaton.NewAutomatonWithTransitionMap},
		&automatonMatcher{constructor: automaton.NewAutomatonWithTransitionSlice},
	}
}

type matchingTestCase struct {
	pattern string
	text    string
	result  bool
}

func makeMatchingTestCases() []matchingTestCase {
	return []matchingTestCase{
		{
			pattern: "GCAGAGAG",
			text:    "GCATCGCAGAGAGTATACAGTACG",
			result:  true,
		},
		{
			pattern: "ACACAGA",
			text:    "ACACACACACACACACACACACACAGAC",
			result:  true,
		},
	}
}

func TestMatching(t *testing.T) {
	testCases := makeMatchingTestCases()

	for _, m := range makeMatchers() {
		for _, tc := range testCases {
			t.Run(m.String(), func(t *testing.T) {
				t.Run(tc.pattern, func(t *testing.T) {
					m.prepare(tc.pattern, tc.text)
					assert.Equal(t, tc.result, m.search())
				})
			})
		}
	}
}

func BenchmarkMatching(b *testing.B) {
	testCases := makeMatchingTestCases()

	for _, m := range makeMatchers() {
		b.Run(m.String(), func(b *testing.B) {
			for _, tc := range testCases {
				b.Run(tc.pattern, func(b *testing.B) {
					m.prepare(tc.pattern, tc.text)
					for i := 0; i < b.N; i++ {
						result := m.search()
						if result != tc.result {
							b.FailNow()
						}
					}
				})
			}
		})
	}
}
