package piscine

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = n
		l.Tail = l.Head
	} else {
		next := &NodeL{
			Data: l.Head.Data,
			Next: l.Head.Next,
		}
		l.Head = n
		n.Next = next
	}
}
