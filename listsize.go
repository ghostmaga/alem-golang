package piscine

func ListSize(l *List) int {
	s := 0
	if l.Head == nil {
		return 0
	} else if l.Head == l.Tail {
		return 1
	} else {
		for l.Head != nil {
			l.Head = l.Head.Next
			s++
		}
	}
	return s
}
