package lists

import (
	"strconv"
)

type Node struct {
	Data int
	Next *Node
}

func (n Node) String() string {
	if n.Next == nil {
		return strconv.Itoa(n.Data)
	}
	return strconv.Itoa(n.Data) + " -> " + n.Next.String()
}

func (n Node) Len() int {
	if n.Next == nil {
		return 1
	}
	return 1 + n.Next.Len()
}
