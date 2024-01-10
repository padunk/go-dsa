package structure

type Stack struct {
	Top *Node
}

func (s *Stack) Push(v int) {
	s.Top = &Node{
		value: v,
		prev:  s.Top.prev,
	}
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	value := s.Top.value
	s.Top = s.Top.prev
	return value
}

func (s *Stack) Peek() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.Top.value
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

// type Stack struct {
// 	Nodes  []Node
// 	Length int
// 	Head   Node
// }

// func (s Stack) New() Stack {
// 	return Stack{
// 		Length: 0,
// 		Nodes:  []Node{},
// 		Head:   Node{},
// 	}
// }

// func (s Stack) Push(v int) {
// 	s.Length++
// 	if s.Length == 0 {
// 		s.Head = Node{}
// 		return
// 	}
// 	prevHead := &s.Head
// 	s.Head = Node{
// 		value: v,
// 		prev:  prevHead,
// 	}
// 	return
// }

// func (s Stack) Pop() {
// 	if s.Length == 0 {
// 		return
// 	}

// 	if s.Length-1 <= 0 {
// 		s.Length = 0
// 	} else {
// 		s.Length--
// 	}
// 	s.Head = *s.Head.prev

// 	return
// }

// func (s Stack) Peek() int {
// 	return s.Head.value
// }
