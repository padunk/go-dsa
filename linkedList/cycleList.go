package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	next := head.Next
	list := map[*ListNode]bool{
		head: true,
	}

	for next != nil {
		if list[next] {
			return true
		}
		list[next] = true
		next = next.Next
	}

	return false
}

// other solutions
func hasCycleOther(head *ListNode) bool {
	visited_nodes := make(map[*ListNode]bool)
	current_node := head
	for current_node != nil {
		if visited_nodes[current_node] {
			return true
		}
		visited_nodes[current_node] = true
		current_node = current_node.Next
	}
	return false
}

func hasCycleOtherFloyd(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
