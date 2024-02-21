package binarytree

func kthSmallest(root *TreeNode, k int) int {
	var path []int

	// in-order traversal
	inOrder(root, &path)

	return path[k]
}

func inOrder(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	inOrder(root.Left, path)
	*path = append(*path, root.Val)
	inOrder(root.Right, path)
}
