package structure

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length uint
}

// func peek(l *LinkedList) *Node {
// 	if l.head.value {
// 		return l.head
// 	}
// }

func (l *LinkedList) prepend(n *Node) {
	if l.length > 0 {
		n.next = l.head
		l.head.prev = n
		l.head = n
	} else {
		l.head = n
		l.tail = n
	}
	l.length++
}

func (l *LinkedList) append(n *Node) {
	if l.length > 0 {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
	} else {
		l.head = n
		l.tail = n
	}
}
