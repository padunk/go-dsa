package structure

func depthFirstSearch(curr *BinaryNode, needle int) bool {
	if curr == nil {
		return false
	}

	if curr.value == needle {
		return true
	}

	if curr.value < needle {
		return depthFirstSearch(curr.right, needle)
	}

	return depthFirstSearch(curr.left, needle)
}

func dfs(head *BinaryNode, needle int) bool {
	return depthFirstSearch(head, needle)
}
