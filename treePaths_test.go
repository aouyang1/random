package random

import (
	"testing"
)

func TestGetPaths(t *testing.T) {
	a := &Node{Name: "a"}
	b := &Node{Name: "b"}
	c := &Node{Name: "c"}
	d := &Node{Name: "d"}
	a.Next = []*Node{b, c}
	b.Next = []*Node{d}
	c.Next = []*Node{d}

	expected := []Nodes{
		{a, b, d},
		{a, c, d},
	}
	paths := a.GetPaths()
	if len(paths) != len(expected) {
		t.Fatalf("expected %d results but got %d", len(expected), len(paths))
	}
	for i, p := range paths {
		if p.String() != expected[i].String() {
			t.Fatalf("expected %s, but got %s", expected[i].String(), p.String())
		}
	}
}
