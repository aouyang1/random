package random

import (
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	testCases := []struct {
		ip       string
		expected string
	}{
		{
			"1.1.1.1",
			"1[.]1[.]1[.]1",
		},
		{
			"255.100.50.0",
			"255[.]100[.]50[.]0",
		},
	}

	for _, c := range testCases {
		if res := defangIPaddr(c.ip); res != c.expected {
			t.Fatalf("Expected %s, but got %s", c.expected, res)
		}
	}
}
