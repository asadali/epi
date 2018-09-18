package lists

func MergeSorted(l1, l2 *Node) *Node {
	l1curr := l1
	l2curr := l2

	var head, curr *Node
	for l1curr != nil && l2curr != nil {
		if l2curr.Data > l1curr.Data {
			largernode = l2curr
			l2curr = l2curr.Next
		} else {
			largernode = l1curr
			l1curr = l1curr.Next
		}
		temp = head
		head = largernode
		largernode.next = temp
	}
}
