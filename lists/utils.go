package lists

func Create(n int) *Node {
	var head *Node
	for i := n; i > 0; i-- {
		node := Node{Data: i}
		node.Next = head
		head = &node
	}
	return head
}
