package random

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	testCases := []struct {
		str      string
		expected string
	}{
		{"A", "a"},
		{"", ""},
		{"b", "b"},
		{"AbC", "abc"},
	}

	for _, c := range testCases {
		if res := toLowerCase(c.str); res != c.expected {
			t.Fatalf("Expected %s, but got %s", c.expected, res)
		}
	}
}
