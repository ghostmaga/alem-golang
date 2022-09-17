package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	if l.Head == nil {
		return nil
	}
	for l.Head != nil {
		if comp(l.Head.Data, ref) {
			return &ref
		}
		l.Head = l.Head.Next
	}
	return nil
}
