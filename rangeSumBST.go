package random

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
	}

	return rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}
