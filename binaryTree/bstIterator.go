package binarytree

type BSTIterator struct {
	root    *TreeNode
	path    []*TreeNode
	pointer int
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{
		root:    root,
		path:    []*TreeNode{},
		pointer: 0,
	}

	inOrderTraverse(b.root, &b.path)

	return b
}

func (this *BSTIterator) Next() int {
	val := this.path[this.pointer].Val
	this.pointer++
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.pointer < len(this.path)
}

func inOrderTraverse(root *TreeNode, path *[]*TreeNode) {
	walk(root, path)
}

func walk(root *TreeNode, path *[]*TreeNode) {
	if root == nil {
		return
	}
	walk(root.Left, path)
	*path = append(*path, root)
	walk(root.Right, path)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
