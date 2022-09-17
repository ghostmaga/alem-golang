package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	cur := l
	if l == nil {
		return nil
	}
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
