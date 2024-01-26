package binarytree

import "slices"

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	rootIndex := 0
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}

	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]

	leftPreorder := preorder[1 : len(leftInorder)+1]
	rightPreorder := preorder[len(leftInorder)+1:]

	root := &TreeNode{Val: rootVal}
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)

	return root
}

// Other Solutions
func buildTreeOther(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootIndex := slices.Index(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:rootIndex+1], inorder[:rootIndex]),
		Right: buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:]),
	}
}
