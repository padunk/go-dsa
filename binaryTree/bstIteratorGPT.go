package binarytree

type BSTIteratorGPT struct {
	stack []*TreeNode
}

func ConstructorGPT(root *TreeNode) BSTIteratorGPT {
	return BSTIteratorGPT{stack: []*TreeNode{root}}
}

/** @return the next smallest number */
func (this *BSTIteratorGPT) Next() int {
	for len(this.stack) > 0 {
		// Traverse left subtree until the leftmost leaf
		for this.stack[len(this.stack)-1].Left != nil {
			this.stack = append(this.stack, this.stack[len(this.stack)-1].Left)
		}

		// Get the leftmost leaf
		node := this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]

		// Move to the right subtree
		if node.Right != nil {
			this.stack = append(this.stack, node.Right)
		}

		return node.Val
	}

	return -1 // Shouldn't reach here if the iterator is used correctly
}

/** @return whether we have a next smallest number */
func (this *BSTIteratorGPT) HasNext() bool {
	return len(this.stack) > 0
}
