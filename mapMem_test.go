package random

import (
	"math/rand"
	"testing"
)

type randRule struct {
	uid        uint64
	metric     string
	fabric     string
	tag        string
	rangequery string
	iql        string
	mout       string
	cout       string
	atype      []string
}

var x map[uint64]*randRule

func BenchmarkEmptyMapCap(b *testing.B) {
	rand.Seed(86)
	numRules := int(280000)

	for i := 0; i < b.N; i++ {
		x = make(map[uint64]*randRule, numRules)
		for j := 0; j < numRules; j++ {
			r := &randRule{
				uid:        rand.Uint64(),
				metric:     "asdfa;sglkhwaawefawefasdfa;sglkhwaawefawe",
				fabric:     "fabric1",
				tag:        "",
				rangequery: "%{fabric1}.myservice.1",
				iql:        "",
				mout:       "ptag1",
				cout:       "ptag1",
				atype:      []string{"aggregate", "max", "min", "average"},
			}
			x[r.uid] = r
		}
	}
}
