package binarytree

func rightSideView(root *TreeNode) []int {
	result := []int{}
	walkRightSide(root, &result)
	return result
}

func walkRightSide(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	*path = append(*path, root.Val)
	if root.Right == nil && root.Left != nil {
		walkRightSide(root.Left, path)
	}
	walkRightSide(root.Right, path)
}

func rightSideViewGPT(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[i]

			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		queue = queue[size:]
	}

	return result
}
