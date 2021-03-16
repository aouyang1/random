package random

import (
	"testing"
)

type testDataLRU struct {
	size     int
	ops      []lruOP
	expected []*expectedLRU
}

type lruOP struct {
	op  string
	key string
	val string
}

type expectedLRU struct {
	val string
	err error
}

func TestLRU(t *testing.T) {
	testData := []testDataLRU{
		{
			4,
			[]lruOP{
				{"put", "a", "apple"},
				{"get", "a", ""},
				{"put", "b", "bear"},
				{"put", "c", "corn"},
				{"put", "d", "deer"},
				{"get", "a", ""},
				{"get", "b", ""},
				{"get", "c", ""},
				{"get", "d", ""},
				{"put", "e", "elephant"},
				{"get", "a", ""},
				{"get", "e", ""},
				{"get", "b", ""},
				{"put", "f", "fairy"},
				{"get", "c", ""},
				{"get", "d", ""},
				{"get", "b", ""},
			},
			[]*expectedLRU{
				nil,
				&expectedLRU{"apple", nil},
				nil,
				nil,
				nil,
				&expectedLRU{"apple", nil},
				&expectedLRU{"bear", nil},
				&expectedLRU{"corn", nil},
				&expectedLRU{"deer", nil},
				nil,
				&expectedLRU{"", errNoKey},
				&expectedLRU{"elephant", nil},
				&expectedLRU{"bear", nil},
				nil,
				&expectedLRU{"", errNoKey},
				&expectedLRU{"deer", nil},
				&expectedLRU{"bear", nil},
			},
		},
		{
			2,
			[]lruOP{
				{"put", "2", "1"},
				{"put", "1", "1"},
				{"put", "2", "3"},
				{"put", "4", "1"},
				{"get", "1", ""},
				{"get", "2", ""},
			},
			[]*expectedLRU{
				nil,
				nil,
				nil,
				nil,
				&expectedLRU{"", errNoKey},
				&expectedLRU{"3", nil},
			},
		},
	}
	for _, td := range testData {
		lru := NewLRU(td.size)
		for i, ops := range td.ops {
			switch ops.op {
			case "get":
				val, err := lru.Get(ops.key)
				exp := td.expected[i]
				if exp == nil {
					t.Error("should have gotten a result")
					return
				}
				if exp.err != err {
					t.Errorf("expected error %v, but got %v", exp.err, err)
					return
				}
				if exp.val != val {
					t.Errorf("expected value, %s, but got %s", exp.val, val)
					return
				}
			case "put":
				lru.Put(ops.key, ops.val)
			}
		}
	}

}
