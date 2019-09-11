package random

import (
	"testing"
)

func TestLongestRepeatedPrefix(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"flow",
		},
		{
			[]string{""},
			"",
		},
		{
			[]string{},
			"",
		},
		{
			[]string{"flower", "flow", "flight", "car", "cars", "carter"},
			"flow",
		},
		{
			[]string{"flower", "flow", "flight", "pretend", "pretentious", "pretense"},
			"preten",
		},
	}

	for _, c := range testCases {
		if res := longestRepeatedPrefix(c.strs); res != c.expected {
			t.Fatalf("Expected %s, but got %s", c.expected, res)
		}
	}
}
