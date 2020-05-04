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
		&callbackMatcher{matcher: karpRabin},
	}
}

type matchingTestCase struct {
	pattern string
	text    string
	result  bool
}

const largeUnicodeText = `
Владимир Щировский

Вчера я умер и меня
Старухи чинно обмывали.
Потом - толпа и в душном зале
Блистали капельки огня.
И было очень тошно мне
Взирать на смертный мой декорум,
Внимать безмерно глупым спорам
О некой Божеской стране.
И становлся страшным зал
От пенья, ладана и плача...
И, если б я бы вам сказал,
Что смерть свершается иначе...
Но мчалось солнце, шла весна,
Звенели деньги, пели люди,
И отходили от окна,
Случайно вспомнив о простуде.
Сквозь запотевшее стекло
Вбегал апрель крылатой ланью.
А в это время утекло
Моё посмертное сознанье.
И друг мой надевал пальто,
И день был светел, светел, светел...
И как я перешёл в ничто -
Никто, конечно, не заметил.

1929 
`

func makeMatchingTestCases() []matchingTestCase {
	return []matchingTestCase{
		// Positive
		{
			pattern: "ABC",
			text:    "ABC",
			result:  true,
		},
		{
			pattern: "GCAGAGAG",
			text:    "GCATCGCAGAGAGTATACAGTACG",
			result:  true,
		},
		{
			pattern: "ACACACACACACACACACACAGA",
			text:    "ACACACACACACACACACACACACACACACACACACACACAGAC",
			result:  true,
		},
		{
			pattern: "ГДЕ",
			text:    "АБВГДЕ",
			result:  true,
		},
		{
			pattern: "заметил",
			text:    largeUnicodeText,
			result:  true,
		},
		// Negative
		{
			pattern: "ABC",
			text:    "ABD",
			result:  false,
		},
		{
			pattern: "ABC",
			text:    "ABDE",
			result:  false,
		},
		{
			pattern: "АБВ",
			text:    "АБГ",
			result:  false,
		},
		{
			pattern: "АБВ",
			text:    "АБГЕ",
			result:  false,
		},
		{
			pattern: "ACACACACACACACACACACAGB",
			text:    "ACACACACACACACACACACACACACACACACACACACACAGAC",
			result:  false,
		},
	}
}

func TestMatching(t *testing.T) {
	testCases := makeMatchingTestCases()

	for _, m := range makeMatchers() {
		t.Run(m.String(), func(t *testing.T) {
			for i, tc := range testCases {
				t.Run(fmt.Sprint(i), func(t *testing.T) {
					m.prepare(tc.pattern, tc.text)
					assert.Equal(t, tc.result, m.search())
				})
			}
		})
	}
}

func BenchmarkMatching(b *testing.B) {
	testCases := makeMatchingTestCases()

	for _, m := range makeMatchers() {
		b.Run(m.String(), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, tc := range testCases {

					// don't take into account preparation phase, it's too complex
					b.StopTimer()
					m.prepare(tc.pattern, tc.text)
					b.StartTimer()

					result := m.search()
					if result != tc.result {
						b.FailNow()
					}
				}
			}
		})
	}
}
