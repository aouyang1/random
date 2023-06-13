package bitmap

type Label string
type Value string
type LabelValues map[Label]Value

type Group struct {
	Attributes []LabelValues                  `json:"attributes"`
	Index      map[Label]map[Value]*Container `json:"index"`
}

func NewGroup() *Group {
	return &Group{
		Index: make(map[Label]map[Value]*Container),
	}
}

// Insert stores new label values and indexes them
func (g *Group) Insert(lv LabelValues) {
	if g.Exists(lv) {
		return
	}

	g.Attributes = append(g.Attributes, lv)
	idx := len(g.Attributes) - 1
	for l, v := range lv {
		values, exists := g.Index[l]
		if !exists {
			values = make(map[Value]*Container)
			g.Index[l] = values
		}
		container, exists := values[v]
		if !exists {
			container = NewContainer()
			values[v] = container
		}
		container.Insert(idx)
	}
}

// Exists checks if the input label values exists in the group already
func (g *Group) Exists(lv LabelValues) bool {
	for l, v := range lv {
		values, exists := g.Index[l]
		if !exists {
			return false
		}
		_, exists = values[v]
		if !exists {
			return false
		}
	}
	return true
}

// Search finds any label values that match with the input search label values. Any of the
// stored label values must contain at least all of the label values provided by the search
// input to be returned.
func (g *Group) Search(lv LabelValues) []LabelValues {
	var res *Container
	for l, v := range lv {
		values, exists := g.Index[l]
		if !exists {
			return nil
		}
		container, exists := values[v]
		if !exists {
			return nil
		}
		if res == nil {
			res = container.Copy()
		}
		res.And(container)
	}

	idx := res.GetAll()
	if len(idx) == 0 {
		return nil
	}
	out := make([]LabelValues, 0, len(idx))
	for _, i := range idx {
		out = append(out, g.Attributes[i])
	}
	return out
}
