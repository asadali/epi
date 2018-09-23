package lists

func MergeSorted(l1, l2 *Node) *Node {
	var head, smaller, prev, temp *Node
	for l1 != nil && l2 != nil {
		if l2.Data < l1.Data {
			smaller = l2
			l2 = l2.Next
		} else {
			smaller = l1
			l1 = l1.Next
		}
		// first node
		if prev == nil {
			head = smaller
			prev = smaller
		} else {
			if prev.Next == smaller {
				// no linking needed
				prev = prev.Next
			} else {
				temp = prev.Next
				prev.Next = smaller
				smaller.Next = temp
				prev = smaller
			}
		}
	}

	if l1 != nil {
		prev.Next = l1
	} else if l2 != nil {
		prev.Next = l2
	}
	return head
}
