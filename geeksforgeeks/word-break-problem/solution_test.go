package word_break_problem_test

import (
	"testing"

	sut "github.com/minizilla/minmax/geeksforgeeks/word-break-problem"
	"github.com/minizilla/testr"
)

func TestWordBreakProblem(t *testing.T) {
	dict := []string{"pine", "apple", "pen", "pineapple", "applepen"}

	tests := map[string]struct {
		s   string
		out []string
	}{
		"empty":           {"", nil},
		"pinepen":         {"pinepen", []string{"pine pen"}},
		"pineapple":       {"pineapple", []string{"pineapple", "pine apple"}},
		"penpineapplepen": {"penpineapplepen", []string{"pen pine applepen", "pen pineapple pen", "pen pine apple pen"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert := testr.New(t)
			out := sut.Solution(tc.s, dict)
			assert.Equal(out, tc.out)
		})
	}
}
