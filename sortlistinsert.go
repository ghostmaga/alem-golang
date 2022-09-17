package piscine

func push(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	if l == nil {
		l.Data = data_ref
	}
	push(l, data_ref)
	cur := l
	for cur != nil {
		next := cur.Next
		for next != nil {
			if cur.Data > next.Data {
				cur.Data, next.Data = next.Data, cur.Data
			}
			next = next.Next
		}
		cur = cur.Next
	}
	return l
}
