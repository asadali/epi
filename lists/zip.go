package lists

func Zip(head *Node) {
	mid := FindMidpoint(head)
	revmid := Reverse(mid)
	Interleave(head, revmid)
}

func FindMidpoint(head *Node) *Node {
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
			slow = slow.Next
		}
	}
	return slow
}

func Reverse(head *Node) *Node {
	var prev, next *Node
	curr := head
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

// interleave two equal lists or
// with difference of length 1
// l1 is the original list
// l2 is the halflist which differ by 1
func Interleave(l1, l2 *Node) {
	for l2 != nil && l1.Next != l2 {
		l1next := l1.Next
		l2next := l2.Next

		l1.Next = l2
		l2.Next = l1next
		l2 = l2next
		l1 = l1next
	}
}
