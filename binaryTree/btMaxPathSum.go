package binarytree

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	_ = dfsMathPathSum(root, &maxSum)
	return maxSum
}

func dfsMathPathSum(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	// Recursively calculate the maximum path sum of the left and right subtrees
	leftSum := max(0, dfsMathPathSum(node.Left, maxSum))
	rightSum := max(0, dfsMathPathSum(node.Right, maxSum))

	// Update the maximum path sum overall if necessary
	*maxSum = max(*maxSum, leftSum+rightSum+node.Val)

	// Return the maximum path sum that includes this node
	return max(leftSum, rightSum) + node.Val
}
