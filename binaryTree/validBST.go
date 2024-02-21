package binarytree

func isValidBST(root *TreeNode) bool {
	var path []int
	var inOrderTraversal func(*TreeNode)

	inOrderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrderTraversal(node.Left)
		path = append(path, node.Val)
		inOrderTraversal(node.Right)
	}

	for i := 1; i < len(path); i++ {
		if path[i] <= path[i-1] {
			return false
		}
	}

	return true
}

func valid(r, min, max *TreeNode) bool {
	if r == nil {
		return true
	}
	if max != nil && max.Val <= r.Val {
		return false
	}
	if min != nil && min.Val >= r.Val {
		return false
	}
	return valid(r.Left, min, r) && valid(r.Right, r, max)
}

func isValidBSTOther(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return valid(root, nil, nil)
}
