package random

import (
	"testing"
)

func TestUniqueMorseRepresentations(t *testing.T) {
	testCases := []struct {
		words    []string
		expected int
	}{
		{
			[]string{"gin", "zen", "gig", "msg"},
			2,
		},
	}

	for _, c := range testCases {
		if res := uniqueMorseRepresentations(c.words); res != c.expected {
			t.Fatalf("Expected %d representations but got %d", c.expected, res)
		}
	}
}
