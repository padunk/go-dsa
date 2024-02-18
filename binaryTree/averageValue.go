package binarytree

func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{float64(root.Val)}
	preOrderTraversal(root, &result)
	return result
}

func preOrderTraversal(root *TreeNode, result *[]float64) *[]float64 {
	if root.Left == nil && root.Right == nil {
		return result
	}

	var average float64

	if root.Left == nil {
		average = float64(root.Right.Val)
	} else if root.Right == nil {
		average = float64(root.Left.Val)
	} else {
		average = (float64(root.Left.Val) + float64(root.Right.Val)) / 2
	}

	*result = append(*result, average)
	preOrderTraversal(root.Left, result)
	preOrderTraversal(root.Right, result)

	return result
}

func averageOfLevelsGPT(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelSum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		average := float64(levelSum) / float64(levelSize)
		result = append(result, average)
	}

	return result
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevelsOther(root *TreeNode) []float64 {
	stack := []*TreeNode{root}

	res := []float64{float64(root.Val)}
	level := 1
	sums := 0.0
	counts := 0.0
	for len(stack) != 0 {
		cur := stack[0]
		stack = stack[1:]

		if cur.Right != nil {
			stack = append(stack, cur.Right)
			counts++
			sums += float64(cur.Right.Val)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
			counts++
			sums += float64(cur.Left.Val)
		}

		level--
		if level == 0 {
			level = len(stack)
			res = append(res, sums/counts)
			sums = 0
			counts = 0
		}
	}
	return res[:len(res)-1]
}
