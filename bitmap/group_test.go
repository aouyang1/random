package bitmap

import (
	"encoding/json"
	"fmt"
)

func ExampleGroupSearch() {
	g := NewGroup()
	g.Insert(LabelValues{
		Label("1"): Value("a"),
		Label("2"): Value("b"),
	})
	g.Insert(LabelValues{
		Label("1"): Value("a"),
		Label("2"): Value("b"),
		Label("3"): Value("c"),
	})
	// insert duplicate
	g.Insert(LabelValues{
		Label("1"): Value("a"),
		Label("2"): Value("b"),
		Label("3"): Value("c"),
	})
	g.Insert(LabelValues{
		Label("1"): Value("a"),
		Label("2"): Value("b"),
		Label("3"): Value("c"),
		Label("4"): Value("d"),
	})
	g.Insert(LabelValues{
		Label("1"): Value("e"),
		Label("3"): Value("f"),
	})

	res := g.Search(LabelValues{
		Label("1"): Value("a"),
		Label("2"): Value("b"),
		Label("3"): Value("c"),
	})

	fmt.Println(res)
	out, _ := json.Marshal(g.Attributes)
	fmt.Printf("Attributes: %+v\n", string(out))

	out, _ = json.Marshal(g.Index)
	fmt.Printf("Index: %+v", string(out))

	// Output: [map[1:a 2:b 3:c] map[1:a 2:b 3:c 4:d]]
	// Attributes: [{"1":"a","2":"b"},{"1":"a","2":"b","3":"c"},{"1":"a","2":"b","3":"c","4":"d"},{"1":"e","3":"f"}]
	// Index: {"1":{"a":{"bitmap":[7]},"e":{"bitmap":[8]}},"2":{"b":{"bitmap":[7]}},"3":{"c":{"bitmap":[6]},"f":{"bitmap":[8]}},"4":{"d":{"bitmap":[4]}}}
}
