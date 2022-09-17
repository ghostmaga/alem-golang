package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	if n1 == nil && n2 == nil {
		return nil
	} else if n1 == nil {
		return n2
	} else if n2 == nil {
		return n1
	}

	new := &NodeI{Data: n1.Data}
	n1 = n1.Next
	for n1 != nil {
		push(new, n1.Data)
		n1 = n1.Next
	}
	for n2 != nil {
		push(new, n2.Data)
		n2 = n2.Next
	}

	ListSort(new)
	return new
}
