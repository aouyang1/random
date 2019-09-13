package random

// TreeNode Is a node that points to a left and right TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(v int, l *TreeNode, r *TreeNode) *TreeNode {
	return &TreeNode{v, l, r}
}

// Build return the treeNode after building out the entire tree
// from the provided list of integers
func BuildTreeNode(c []*int) *TreeNode {
	if len(c) == 0 {
		return &TreeNode{}
	}

	rootNode := NewTreeNode(*c[0], nil, nil)
	var levelNodes, nextLevelNodes []*TreeNode
	levelNodes = append(levelNodes, rootNode)

	if len(c) == 1 {
		return rootNode
	}
	c = c[1:]

	for _, n := range levelNodes {
		if len(c) > 0 {
			if c[0] != nil {
				node := NewTreeNode(*c[0], nil, nil)
				n.Left = node
				nextLevelNodes = append(nextLevelNodes, node)
			}
			if len(c) == 1 {
				return rootNode
			}
			c = c[1:]

			if c[0] != nil {
				node := NewTreeNode(*c[0], nil, nil)
				n.Right = node
				nextLevelNodes = append(nextLevelNodes, node)
			}
			if len(c) == 1 {
				return rootNode
			}
			c = c[1:]
		}
	}
	return rootNode
}
