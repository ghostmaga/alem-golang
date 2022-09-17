package piscine

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	if l.Head == nil {
		return
	}
	upd := &List{}
	for l.Head != nil {
		if cond(l.Head) {
			f(l.Head)
		}
		ListPushBack(upd, l.Head.Data)
		l.Head = l.Head.Next
	}
	*l = *upd
}
