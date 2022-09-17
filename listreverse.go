package piscine

func ListReverse(l *List) {
	cur := l.Head
	var next *NodeL
	var prev *NodeL
	l.Tail = l.Head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
}
