package structure

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

// search
func walk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	path = append(path, curr.value)

	walk(curr.left, path)
	walk(curr.right, path)

	return path
}

func pre_order_search(head *BinaryNode) []int {
	return walk(head, []int{})
}

// implement breadth first search
func bfs(head *BinaryNode, needle int) bool {
	queue := []*BinaryNode{head}

	for i := 0; len(queue) > 0; i++ {
		curr := queue[i]
		queue = queue[i+1:]

		if curr.value == needle {
			return true
		}

		queue = append(queue, curr.left)
		queue = append(queue, curr.right)
	}

	return false
}

// compare if two binary tree is the same
func compareBinaryTree(a, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compareBinaryTree(a.left, b.left) && compareBinaryTree(a.right, b.right)
}
