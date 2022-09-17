package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		*l1 = *l2
		return
	}
	l1.Tail.Next = l2.Head
}
