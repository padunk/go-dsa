package binarytree

import "slices"

func buildTreePost(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	rootVal := postorder[len(postorder)-1]
	rootIndex := slices.Index(inorder, rootVal)
	return &TreeNode{
		Val:   rootVal,
		Left:  buildTree(inorder[:rootIndex], postorder[:rootIndex]),
		Right: buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1]),
	}
}
