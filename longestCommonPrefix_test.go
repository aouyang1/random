package random

import "testing"

func TestLongestcommonPrefix(t *testing.T) {
	testCases := []struct {
		strs     []string
		expected string
	}{
		{
			[]string{"flower", "flow", "flight"},
			"fl",
		},
		{
			[]string{"dog", "racecar", "car"},
			"",
		},
		{
			[]string{"dog", "cars", "car"},
			"",
		},
		{
			[]string{"dog"},
			"dog",
		},
		{
			[]string{""},
			"",
		},
		{
			[]string{},
			"",
		},
	}

	for _, c := range testCases {
		if res := longestCommonPrefix(c.strs); res != c.expected {
			t.Fatalf("Expected %s, but got %s", c.expected, res)
		}
	}
}
