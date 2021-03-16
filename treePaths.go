package random

type Node struct {
	Name string
	Next Nodes
}

type Nodes []*Node

func (n Nodes) String() string {
	var out string
	for _, nn := range n {
		out += nn.Name + ","
	}
	if len(n) > 0 {
		return "[" + out[:len(out)-1] + "]"
	}
	return ""
}

func (n *Node) GetNext() Nodes {
	return n.Next
}

func (n *Node) GetPaths() []Nodes {
	var res []Nodes
	next := n.GetNext()
	if len(next) == 0 {
		return append(res, Nodes{n})
	}

	for _, nn := range next {
		paths := nn.GetPaths()
		for _, p := range paths {
			res = append(res, append([]*Node{n}, p...))
		}
	}
	return res
}
