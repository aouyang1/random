package random

import (
	"testing"
)

func TestNumJewelsInStones(t *testing.T) {
	testCases := []struct {
		j        string
		s        string
		expected int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}

	for _, c := range testCases {
		if res := numJewelsInStones(c.j, c.s); res != c.expected {
			t.Fatalf("Expected %d, but got %d for j: %s, and s: %s", c.expected, res, c.j, c.s)
		}
	}
}
