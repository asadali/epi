package lists

func CreateWithStartInterval(n, start, interval int) *Node {
	var head *Node
	end := start + interval*(n-1)
	for i := n; i > 0; i-- {
		node := Node{Data: end}
		node.Next = head
		head = &node
		end = end - interval
	}
	return head
}

func CreateWithStart(n, start int) *Node {
	return CreateWithStartInterval(n, start, 1)
}
func Create(n int) *Node {
	return CreateWithStartInterval(n, 1, 1)
}
