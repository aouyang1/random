package random

import (
	"testing"
)

func TestRemoveOuterParentheses(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"(()())(())", "()()()"},
		{"(()())(())(()(()))", "()()()()(())"},
		{"()()", ""},
		{"", ""},
		{"()", ""},
		{"(())()", "()"},
		{"(()(()))", "()(())"},
	}

	for _, c := range testCases {
		if out := removeOuterParentheses(c.input); out != c.expected {
			t.Fatalf("Expected %s, but got %s", c.expected, out)
		}
	}
}
