package binarytree

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depthLeft := 1 + maxDepth(root.Left)
	depthRight := 1 + maxDepth(root.Right)

	if depthLeft > depthRight {
		return depthLeft
	}
	return depthRight
}
