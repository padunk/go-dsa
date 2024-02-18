package binarytree

import "slices"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	j := 0

	for len(queue) > 0 {
		size := len(queue)
		result = append(result, []int{})

		for i := 0; i < size; i++ {
			node := queue[i]
			result[j] = append(result[j], node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
		j++
	}

	for i, v := range result {
		if i%2 != 0 {
			slices.Reverse(v)
		}
	}

	return result
}

func zigzagLevelOrderGPT(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := [][]int{}
	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		var levelNodes []int
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				levelNodes = append(levelNodes, node.Val)
			} else {
				levelNodes = append([]int{node.Val}, levelNodes...)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, levelNodes)
		level++
	}

	return result
}
