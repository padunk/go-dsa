package linkedlist

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	copy := &Node{}
	current := head
	copyMap := map[int]*Node{}
	copyMapKey := 0

	// copy the node value and next node
	for current != nil {
		copy.Val = head.Val
		copy.Random = &Node{}

		if current.Next != nil {
			copy.Next = &Node{}
			copyMap[copyMapKey] = copy
			copyMapKey++

			copy = copy.Next
		}
		current = current.Next
	}

	// copy the random node
	current = head
	for current != nil {
		if current.Random != nil {
			index := findIndex(head, current.Random)
			copy.Random = copyMap[index]
		} else {
			copy.Random = nil
		}
		copy = copy.Next
		current = current.Next
	}

	return copy
}

func findIndex(head *Node, r *Node) int {
	i := 0
	current := head

	for current != nil {
		if current == r {
			break
		}
		current = current.Next
		i++
	}

	return i
}

// GPT
func copyRandomListGPT(head *Node) *Node {
	if head == nil {
		return nil
	}

	// Create a mapping between original and copied nodes
	nodeMap := make(map[*Node]*Node)

	// First pass: create nodes without random pointers and populate the map
	current := head
	for current != nil {
		nodeMap[current] = &Node{Val: current.Val}
		current = current.Next
	}

	// Second pass: set next and random pointers based on the map
	current = head
	for current != nil {
		copyNode := nodeMap[current]
		copyNode.Next = nodeMap[current.Next]
		copyNode.Random = nodeMap[current.Random]
		current = current.Next
	}

	// Return the head of the copied list
	return nodeMap[head]
}
