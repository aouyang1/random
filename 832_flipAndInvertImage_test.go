package random

import (
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	testCases := []struct {
		input    [][]int
		expected [][]int
	}{
		{
			[][]int{
				[]int{1, 1, 0},
				[]int{1, 0, 1},
				[]int{0, 0, 0},
			},
			[][]int{
				[]int{1, 0, 0},
				[]int{0, 1, 0},
				[]int{1, 1, 1},
			},
		},
		{
			[][]int{
				[]int{1, 1, 0, 0},
				[]int{1, 0, 0, 1},
				[]int{0, 1, 1, 1},
				[]int{1, 0, 1, 0},
			},
			[][]int{
				[]int{1, 1, 0, 0},
				[]int{0, 1, 1, 0},
				[]int{0, 0, 0, 1},
				[]int{1, 0, 1, 0},
			},
		},
	}

	for _, c := range testCases {
		out := flipAndInvertImage(c.input)
		if len(out) != len(c.expected) {
			t.Fatalf("Expected %d rows but got %d", len(c.expected), len(out))
		}
		if len(out) > 0 {
			if len(out[0]) != len(c.expected[0]) {
				t.Fatalf("Expected %d columns but got %d", len(c.expected[0]), len(out[0]))
			}
			for ri, row := range out {
				for ci, v := range row {
					if v != c.expected[ri][ci] {
						t.Fatalf("Expected %d on row: %d, col %d, but got %d", c.expected[ri][ci], ri, ci, v)
					}
				}
			}
		}
	}

}
