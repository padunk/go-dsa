package binarytree

import "math"

func getMinimumDifference(root *TreeNode) int {
	minValue := math.Inf(1)
	var path []int
	// in-order traversal
	walkTree(root, &path)

	for j := 1; j < len(path); j++ {
		diff := math.Abs(float64(path[j-1]) - float64(path[j]))
		minValue = min(diff, minValue)
	}

	return int(minValue)
}

func walkTree(root *TreeNode, path *[]int) {
	if root == nil {
		return
	}
	walkTree(root.Left, path)
	*path = append(*path, root.Val)
	walkTree(root.Right, path)
}

func getMinimumDifferenceGPT(root *TreeNode) int {
	res, pre := math.MaxInt64, -1
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val-pre < res {
			res = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return res
}
